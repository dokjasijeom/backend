package common

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/exception"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func GenerateToken(email string, roles []map[string]interface{}, config configuration.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpire, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTE"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"email": email,
		"roles": roles,
		"exp":   time.Now().Add(time.Minute * time.Duration(jwtExpire)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
