package api

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// FIXME: move to env vars
	ExpirationTime = time.Minute * 15
)

var SigningKey = []byte("DApAJQgpjRDHa9Ad")

func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// TODO: authorization level
		"authorized": true,
		// User ID
		"user": userID,
		"exp":  time.Now().Add(ExpirationTime).Unix(),
	})
	tokenStr, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
