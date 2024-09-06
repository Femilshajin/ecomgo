package auth

import (
	"strconv"
	"time"

	"github.com/femilshajin/ecomgo/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	currentTime := time.Now()
	expiration := time.Second * time.Duration(config.Envs.JWTEpiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"iss":       currentTime.Unix(),
		"expiredAt": currentTime.Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
