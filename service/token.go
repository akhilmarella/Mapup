package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"github.com/twinj/uuid"
)

var secretKey = []byte("nothingelse")

func CreateToken() (string, error) {
	id := uuid.NewV4().String()

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["expiry"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Error().Err(err).Any("key", secretKey).Any("action", "utils_token.go_CreateToken").
			Msg("error in creating token with secret")
		return "", err
	}

	return tokenString, nil
}
