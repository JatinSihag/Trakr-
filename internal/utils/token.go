package utils

import (
	"fmt"
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
func ValidateToken(tokenString string)(int64,error){
	token,err:=jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method: %v",token.Header["alg"])
		}
		return secretKey,nil
	})
	if err !=nil{
		return 0,err
	}
	if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		userId :=int64(claims["user_id"].(float64))
		return userId,nil
	}
	return 0,fmt.Errorf("invalid token")
}