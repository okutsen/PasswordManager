package api

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// FIXME: move to env vars
	SigningKey = ""
)

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// TODO: authorization level
		"authorized": true,
		// User ID
		"user": "",
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenStr, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
