package impl

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dokjasijeom/backend/common"
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
	result := userRepository.DB.WithContext(ctx).Where("email = ?", email).Find(&userResult)
	err := result.Error

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
