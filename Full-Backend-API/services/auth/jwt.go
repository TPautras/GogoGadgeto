package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitub.com/TPautras/ecom/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix()})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
