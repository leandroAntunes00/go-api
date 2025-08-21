package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	SecretKey = "secret"
)

func GenerateToken(email string, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"id":    userID,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString([]byte(SecretKey))
}
