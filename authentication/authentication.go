package authentication

import (
	"fmt"
	"gkru-service/helper"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func CreateToken(username string, logger *logrus.Logger) (string, error) {
	if err := godotenv.Load(); err != nil {
        logger.Warn("Gagal memuat file .env")
    }

	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	if privateKeyPath == "" {
		return "", fmt.Errorf("PRIVATE_KEY_PATH tidak diset")
	}

	secretKey, err := os.ReadFile(privateKeyPath)
	helper.PanicIfError(err)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, logger *logrus.Logger) error {
	if err := godotenv.Load(); err != nil {
        logger.Warn("Gagal memuat file .env")
    }

	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	if privateKeyPath == "" {
		return fmt.Errorf("PRIVATE_KEY_PATH tidak diset")
	}

	secretKey, err := os.ReadFile(privateKeyPath)
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
