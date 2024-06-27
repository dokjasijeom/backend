package impl

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
	"os"
)

func NewUserPasswordResetRepositoryImpl(DB *gorm.DB) repository.UserPasswordResetRepository {
	return &userPasswordResetRepositoryImpl{DB: DB}
}

type userPasswordResetRepositoryImpl struct {
	DB *gorm.DB
}

func (uprr *userPasswordResetRepositoryImpl) CreateUserPasswordReset(ctx context.Context, email string) (string, error) {
	// Create user password reset
	token := tokenGenerator()
	result := uprr.DB.WithContext(ctx).Create(&entity.UserPasswordReset{Email: email, Token: token})

	return token, result.Error
}

func (uprr *userPasswordResetRepositoryImpl) GetUserPasswordResetToEmail(ctx context.Context, token string) (string, error) {
	// Get user password reset
	var email string
	releaseMode := os.Getenv("RELEASE_MODE")
	tablePrefix := func(releaseMode string) string {
		if releaseMode == "development" {
			return "dev_"
		} else {
			return ""
		}
	}(releaseMode)
	// expired 3 hours
	expired := uprr.DB.WithContext(ctx).Exec("DELETE FROM " + tablePrefix + "user_password_resets WHERE created_at < NOW() - INTERVAL 3 HOUR")
	if expired.Error != nil {
		return "", expired.Error
	}

	result := uprr.DB.WithContext(ctx).Find(&entity.UserPasswordReset{}, "token = ?", token).Scan(&email)

	return email, result.Error
}

func (uprr *userPasswordResetRepositoryImpl) DeleteUserPasswordReset(ctx context.Context, token string) error {
	// Delete user password reset
	result := uprr.DB.WithContext(ctx).Model(&entity.UserPasswordReset{}).Delete(&entity.UserPasswordReset{}, "token = ?", token)

	return result.Error
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
