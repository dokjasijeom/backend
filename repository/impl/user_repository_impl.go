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
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
	"log"
	"slices"
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
	config := configuration.New()
	result := userRepository.DB.WithContext(ctx).Where("email = ?", email).Preload("Profile").Preload("SubscribeProvider").Find(&userResult)
	err := result.Error

	userResult.Profile.Avatar = config.Get("CLOUDINARY_URL") + userResult.Profile.Avatar
	for i := range userResult.SubscribeProvider {
		userResult.SubscribeProvider[i].Id = 0
	}

	return userResult, err
}

func (userRepository *userRepositoryImpl) GetUserByEmailAndSeries(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	config := configuration.New()
	releaseMode := config.Get("RELEASE_MODE")

	tablePrefix := func(releaseMode string) string {
		if releaseMode == "development" {
			return "dev_"
		} else {
			return ""
		}
	}(releaseMode)
	result := userRepository.DB.WithContext(ctx).Where("email = ?", email).Preload("Profile").Preload("SubscribeProvider").Preload("LikeSeries").Preload("LikeSeries.Genres").Preload("LikeSeries.Publishers").Preload("LikeSeries.PublishDays").Preload("LikeSeries.SeriesAuthors.Person").Preload("LikeSeries.SeriesProvider.Provider").Preload("LikeSeries.Episodes").Preload("RecordSeries", func(db *gorm.DB) *gorm.DB {
		return db.Order(tablePrefix + "user_record_series.id desc")
	}).Preload("RecordSeries.Series").Preload("RecordSeries.Series.Genres").Preload("RecordSeries.Series.SeriesAuthors.Person").Preload("RecordSeries.Series.SeriesProvider.Provider").Preload("RecordSeries.RecordEpisodes", func(db *gorm.DB) *gorm.DB {
		return db.Order(tablePrefix + "user_record_series_episodes.episode_number asc")
	}).Preload("CompleteRecordSeries", func(db *gorm.DB) *gorm.DB {
		return db.Where(tablePrefix+"user_record_series.read_completed = ?", true).Order(tablePrefix + "user_record_series.id desc")
	}).Preload("CompleteRecordSeries.Series").Preload("CompleteRecordSeries.Series.Genres").Preload("CompleteRecordSeries.Series.SeriesAuthors.Person").Preload("CompleteRecordSeries.Series.SeriesProvider.Provider").Preload("CompleteRecordSeries.RecordEpisodes", func(db *gorm.DB) *gorm.DB {
		return db.Order(tablePrefix + "user_record_series_episodes.episode_number asc")
	}).Find(&userResult)

	err := result.Error

	if userResult.Profile.Avatar != "" {
		userResult.Profile.Avatar = config.Get("CLOUDINARY_URL") + userResult.Profile.Avatar
	}

	for i := range userResult.SubscribeProvider {
		userResult.SubscribeProvider[i].Id = 0
	}

	for i := range userResult.LikeSeries {
		userResult.LikeSeries[i].Id = 0

		if userResult.LikeSeries[i].SeriesType == "webnovel" {
			userResult.LikeSeries[i].DisplayTags = "#웹소설 "
		} else {
			userResult.LikeSeries[i].DisplayTags = "#웹툰 "
		}

		for genreI := range userResult.LikeSeries[i].Genres {
			userResult.LikeSeries[i].Genres[genreI].Id = 0
			userResult.LikeSeries[i].DisplayTags += "#" + userResult.LikeSeries[i].Genres[genreI].Name + " "
		}
		userResult.LikeSeries[i].TotalEpisode = uint(len(userResult.LikeSeries[i].Episodes))
		userResult.LikeSeries[i].DisplayTags = userResult.LikeSeries[i].DisplayTags[:len(userResult.LikeSeries[i].DisplayTags)-1]

		// 작가 유형 반영해서 Authors 필드에 반영
		userResult.LikeSeries[i].Authors = make([]entity.Person, 0)
		for _, sa := range userResult.LikeSeries[i].SeriesAuthors {
			sa.Person.Id = 0
			sa.Person.PersonType = sa.PersonType
			userResult.LikeSeries[i].Authors = append(userResult.LikeSeries[i].Authors, sa.Person)
		}
		// 제공자 정보를 Providers 필드에 반영
		userResult.LikeSeries[i].Providers = make([]entity.Provider, 0)
		for _, sp := range userResult.LikeSeries[i].SeriesProvider {
			sp.Provider.Link = sp.Link
			sp.Provider.Id = 0
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

	for i := range userResult.RecordSeries {
		var recordProviders []string
		for _, rEpisode := range userResult.RecordSeries[i].RecordEpisodes {
			if !slices.Contains(recordProviders, rEpisode.ProviderName) {
				recordProviders = append(recordProviders, rEpisode.ProviderName)
			}
		}
		userResult.RecordSeries[i].RecordProviders = recordProviders

		if userResult.RecordSeries[i].SeriesId != 0 {
			userResult.RecordSeries[i].Series.Id = 0

			if userResult.RecordSeries[i].Series.SeriesType == "webnovel" {
				userResult.RecordSeries[i].Series.DisplayTags = "#웹소설 "
			} else {
				userResult.RecordSeries[i].Series.DisplayTags = "#웹툰 "
			}

			for genreI := range userResult.RecordSeries[i].Series.Genres {
				userResult.RecordSeries[i].Series.Genres[genreI].Id = 0
				userResult.RecordSeries[i].Series.DisplayTags += "#" + userResult.RecordSeries[i].Series.Genres[genreI].Name + " "
			}
			userResult.RecordSeries[i].Series.TotalEpisode = uint(len(userResult.RecordSeries[i].Series.Episodes))
			userResult.RecordSeries[i].Series.DisplayTags = userResult.RecordSeries[i].Series.DisplayTags[:len(userResult.RecordSeries[i].Series.DisplayTags)-1]

			// 작가 유형 반영해서 Authors 필드에 반영
			userResult.RecordSeries[i].Series.Authors = make([]entity.Person, 0)
			for _, sa := range userResult.RecordSeries[i].Series.SeriesAuthors {
				sa.Person.Id = 0
				sa.Person.PersonType = sa.PersonType
				userResult.RecordSeries[i].Series.Authors = append(userResult.RecordSeries[i].Series.Authors, sa.Person)
			}
			// 제공자 정보를 Providers 필드에 반영
			userResult.RecordSeries[i].Series.Providers = make([]entity.Provider, 0)
			for _, sp := range userResult.RecordSeries[i].Series.SeriesProvider {
				sp.Provider.Link = sp.Link
				sp.Provider.Id = 0
				userResult.RecordSeries[i].Series.Providers = append(userResult.RecordSeries[i].Series.Providers, sp.Provider)
			}

			// publishDays remove id, Displayorder
			for j, _ := range userResult.RecordSeries[i].Series.PublishDays {
				userResult.RecordSeries[i].Series.PublishDays[j].Id = 0
				userResult.RecordSeries[i].Series.PublishDays[j].DisplayOrder = 0
			}
			// publishers remove field id, description, homepageurl, series
			for j, _ := range userResult.RecordSeries[i].Series.Publishers {
				userResult.RecordSeries[i].Series.Publishers[j].Id = 0
				userResult.RecordSeries[i].Series.Publishers[j].Description = ""
				userResult.RecordSeries[i].Series.Publishers[j].HomepageUrl = ""
				userResult.RecordSeries[i].Series.Publishers[j].Series = nil
			}

			userResult.RecordSeries[i].Series.Thumbnail = config.Get("CLOUDINARY_URL") + userResult.RecordSeries[i].Series.Thumbnail
		} else {
			userResult.RecordSeries[i].Series = nil
		}
	}
	userResult.RecordSeriesCount = uint(len(userResult.RecordSeries))

	for i := range userResult.CompleteRecordSeries {
		var recordProviders []string
		for _, rEpisode := range userResult.CompleteRecordSeries[i].RecordEpisodes {
			if !slices.Contains(recordProviders, rEpisode.ProviderName) {
				recordProviders = append(recordProviders, rEpisode.ProviderName)
			}
		}
		userResult.CompleteRecordSeries[i].RecordProviders = recordProviders

		if userResult.CompleteRecordSeries[i].SeriesId != 0 {
			userResult.CompleteRecordSeries[i].Series.Id = 0

			if userResult.CompleteRecordSeries[i].Series.SeriesType == "webnovel" {
				userResult.CompleteRecordSeries[i].Series.DisplayTags = "#웹소설 "
			} else {
				userResult.CompleteRecordSeries[i].Series.DisplayTags = "#웹툰 "
			}

			for genreI := range userResult.CompleteRecordSeries[i].Series.Genres {
				userResult.CompleteRecordSeries[i].Series.Genres[genreI].Id = 0
				userResult.CompleteRecordSeries[i].Series.DisplayTags += "#" + userResult.CompleteRecordSeries[i].Series.Genres[genreI].Name + " "
			}
			userResult.CompleteRecordSeries[i].Series.TotalEpisode = uint(len(userResult.CompleteRecordSeries[i].Series.Episodes))
			userResult.CompleteRecordSeries[i].Series.DisplayTags = userResult.CompleteRecordSeries[i].Series.DisplayTags[:len(userResult.CompleteRecordSeries[i].Series.DisplayTags)-1]

			// 작가 유형 반영해서 Authors 필드에 반영
			userResult.CompleteRecordSeries[i].Series.Authors = make([]entity.Person, 0)
			for _, sa := range userResult.CompleteRecordSeries[i].Series.SeriesAuthors {
				sa.Person.Id = 0
				sa.Person.PersonType = sa.PersonType
				userResult.CompleteRecordSeries[i].Series.Authors = append(userResult.CompleteRecordSeries[i].Series.Authors, sa.Person)
			}
			// 제공자 정보를 Providers 필드에 반영
			userResult.CompleteRecordSeries[i].Series.Providers = make([]entity.Provider, 0)
			for _, sp := range userResult.CompleteRecordSeries[i].Series.SeriesProvider {
				sp.Provider.Link = sp.Link
				sp.Provider.Id = 0
				userResult.CompleteRecordSeries[i].Series.Providers = append(userResult.CompleteRecordSeries[i].Series.Providers, sp.Provider)
			}

			// publishDays remove id, Displayorder
			for j, _ := range userResult.CompleteRecordSeries[i].Series.PublishDays {
				userResult.CompleteRecordSeries[i].Series.PublishDays[j].Id = 0
				userResult.CompleteRecordSeries[i].Series.PublishDays[j].DisplayOrder = 0
			}
			// publishers remove field id, description, homepageurl, series
			for j, _ := range userResult.CompleteRecordSeries[i].Series.Publishers {
				userResult.CompleteRecordSeries[i].Series.Publishers[j].Id = 0
				userResult.CompleteRecordSeries[i].Series.Publishers[j].Description = ""
				userResult.CompleteRecordSeries[i].Series.Publishers[j].HomepageUrl = ""
				userResult.CompleteRecordSeries[i].Series.Publishers[j].Series = nil
			}

			userResult.CompleteRecordSeries[i].Series.Thumbnail = config.Get("CLOUDINARY_URL") + userResult.CompleteRecordSeries[i].Series.Thumbnail
		} else {
			userResult.CompleteRecordSeries[i].Series = nil
		}
	}
	userResult.CompleteRecordSeriesCount = uint(len(userResult.CompleteRecordSeries))

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

func (userRepository *userRepositoryImpl) DeleteUser(ctx context.Context, id uint) error {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).Where("id = ?", id).Find(&userResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("user not found")
	}
	// exist email change to uuid
	userResult.Email = strings.Replace(uuid.New().String(), "-", "", -1)

	userRepository.DB.WithContext(ctx).Save(&userResult)

	err := userRepository.DB.WithContext(ctx).Model(&userResult).Association("Roles").Clear()
	if err != nil {
		exception.PanicLogging(err)
		return err
	}
	result = userRepository.DB.WithContext(ctx).Delete(&userResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}
	return nil
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

func (userRepository *userRepositoryImpl) UpdateUserPassword(ctx context.Context, id uint, password string) error {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).Where("id = ?", id).Find(&userResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("user not found")
	}
	encodedPassword, err := userRepository.GenerateFromPassword(password)
	if err != nil {
		exception.PanicLogging(err)
		return err
	}
	userResult.Password = encodedPassword
	result = userRepository.DB.WithContext(ctx).Save(&userResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}
	return nil
}

func (userRepository *userRepositoryImpl) UpdateUserProviders(ctx context.Context, id uint, providerIds []uint) error {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).Where("id = ?", id).Find(&userResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("user not found")
	}
	userRepository.DB.WithContext(ctx).Model(&userResult).Association("SubscribeProvider").Clear()
	for _, providerId := range providerIds {
		userRepository.DB.WithContext(ctx).Model(&userResult).Association("SubscribeProvider").Append(&entity.Provider{Id: providerId})
	}
	return nil
}
