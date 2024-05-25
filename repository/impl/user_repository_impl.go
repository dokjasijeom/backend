package impl

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dokjasijeom/backend/common"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
	"log"
	"strings"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) GetAllUsers() error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).Where("email = ?", email).Preload("LikeSeries").Preload("LikeSeries.Genres").Preload("LikeSeries.Publishers").Preload("LikeSeries.PublishDays").Preload("LikeSeries.SeriesAuthors.Person").Preload("LikeSeries.Episodes").Preload("RecordSeries.Series").Find(&userResult)
	err := result.Error

	config := configuration.New()

	for i := range userResult.LikeSeries {
		userResult.LikeSeries[i].Id = 0

		userResult.LikeSeries[i].Id = 0

		if userResult.LikeSeries[i].SeriesType == "webnovel" {
			userResult.LikeSeries[i].DisplayTags = "#웹소설 "
		} else {
			userResult.LikeSeries[i].DisplayTags = "#웹툰 "
		}

		for genreI := range userResult.LikeSeries[i].Genres {
			userResult.LikeSeries[i].DisplayTags += "#" + userResult.LikeSeries[i].Genres[genreI].Name + " "
		}
		userResult.LikeSeries[i].TotalEpisode = uint(len(userResult.LikeSeries[i].Episodes))
		userResult.LikeSeries[i].DisplayTags = userResult.LikeSeries[i].DisplayTags[:len(userResult.LikeSeries[i].DisplayTags)-1]

		// 작가 유형 반영해서 Authors 필드에 반영
		userResult.LikeSeries[i].Authors = make([]entity.Person, 0)
		for _, sa := range userResult.LikeSeries[i].SeriesAuthors {
			sa.Person.PersonType = sa.PersonType
			userResult.LikeSeries[i].Authors = append(userResult.LikeSeries[i].Authors, sa.Person)
		}
		// 제공자 정보를 Providers 필드에 반영
		userResult.LikeSeries[i].Providers = make([]entity.Provider, 0)
		for _, sp := range userResult.LikeSeries[i].SeriesProvider {
			sp.Provider.Link = sp.Link
			userResult.LikeSeries[i].Providers = append(userResult.LikeSeries[i].Providers, sp.Provider)
		}

		// publishDays remove id, Displayorder
		for j, _ := range userResult.LikeSeries[i].PublishDays {
			userResult.LikeSeries[i].PublishDays[j].Id = 0
			userResult.LikeSeries[i].PublishDays[j].DisplayOrder = 0
		}
		// publishers remove field id, description, homepageurl, series
		for j, _ := range userResult.LikeSeries[i].Publishers {
			userResult.LikeSeries[i].Publishers[j].Id = 0
			userResult.LikeSeries[i].Publishers[j].Description = ""
			userResult.LikeSeries[i].Publishers[j].HomepageUrl = ""
			userResult.LikeSeries[i].Publishers[j].Series = nil
		}

		userResult.LikeSeries[i].Thumbnail = config.Get("CLOUDINARY_URL") + userResult.LikeSeries[i].Thumbnail

	}
	userResult.LikeSeriesCount = uint(len(userResult.LikeSeries))

	return userResult, err
}

func (userRepository *userRepositoryImpl) GetUserByEmailAndPassword(email, password string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) CreateUser(email, password string) (entity.User, error) {
	var userResult entity.User
	userResult.Email = email

	encodedPassword, err := userRepository.GenerateFromPassword(password)
	if err != nil {
		exception.PanicLogging(err)
		return userResult, err
	}

	userResult.Password = encodedPassword
	result := userRepository.DB.Create(&userResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		userResult.Id = 0
		return userResult, result.Error
	}
	return userResult, nil
}

func (userRepository *userRepositoryImpl) UpdateUserHashId(ctx context.Context, email string, hashId string) error {
	var userResult entity.User

	result := userRepository.DB.WithContext(ctx).Where("email = ?", email).Find(&userResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("user not found")
	}
	userResult.HashId = hashId
	result = userRepository.DB.WithContext(ctx).Save(&userResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}
	return nil
}

func (userRepository *userRepositoryImpl) UpdateUserByEmail(email string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) DeleteUserByEmail(email string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) Authenticate(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).Where("email = ?", email).Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	err := userRepository.DB.Model(&userResult).Association("Roles").Find(&userResult.Roles)
	if err != nil {
		log.Println(err)
	}
	log.Println(userResult)
	return userResult, nil
}

func (userRepository *userRepositoryImpl) GenerateFromPassword(password string) (string, error) {
	p := &common.HashParams{
		Memory:      8 * 1024,
		Iterations:  1,
		Parallelism: 1,
		SaltLength:  8,
		KeyLength:   32,
	}

	salt, err := userRepository.GenerateRandomSalt(p.SaltLength)
	if err != nil {
		exception.PanicLogging(err)
		return "", err
	}

	passwordHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(passwordHash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func (userRepository *userRepositoryImpl) GenerateRandomSalt(saltLength uint32) ([]byte, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		exception.PanicLogging(err)
		return nil, err
	}

	return b, nil
}

func (userRepository *userRepositoryImpl) CompareHashAndPassword(password string, encodedHash string) (bool, error) {
	p, salt, hash, err := userRepository.DecodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (userRepository *userRepositoryImpl) DecodeHash(encodedHash string) (p *common.HashParams, salt []byte, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("invalid hash format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version of argon2")
	}

	p = &common.HashParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
