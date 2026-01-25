package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(userId int) (string,error) {
	claims:=jwt.MapClaims{
		"user_id":userId,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString(secretKey)
}