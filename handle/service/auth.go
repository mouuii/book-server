package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userId uint) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId;
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString([]byte("secret"))
}
