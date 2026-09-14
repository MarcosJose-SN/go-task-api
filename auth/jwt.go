package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET não configurado")
	}

	secretKey := []byte(secret)

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET não configurado")
	}

	secretKey := []byte(secret)

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido")
		}

		return secretKey, nil
	})
}
