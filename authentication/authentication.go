package authentication

import (
	"fmt"
	"gkru-service/helper"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (string, error) {
	secretKey, err := os.ReadFile("D:\\GoLang\\private.pem")
	helper.PanicIfError(err)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) error {
	secretKey, err := os.ReadFile("D:\\GoLang\\private.pem")
	helper.PanicIfError(err)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secretKey, nil
	})
   
	if err != nil {
	   return err
	}
   
	if !token.Valid {
	   return fmt.Errorf("invalid token")
	}
   
	return nil
 }