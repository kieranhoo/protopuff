package utils

import (
	"errors"
	"fmt"
	"protopuff/internal/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	tokenString, errToken := token.SignedString([]byte(config.JwtSecret))
	if errToken != nil {
		return "", errors.New("failed to create token")
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	tokenString = removeBearerPrefix(tokenString)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.JwtSecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"]
		// _ = claims["exp"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", errors.New("token has expired")
		}
		return email.(string), nil
	}
	return "", errors.New("token invalid")
}
