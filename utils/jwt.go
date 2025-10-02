package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret-key-secret-key")

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	jwtKeyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	}
	token, err := jwt.Parse(tokenStr, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractUserID(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, ok := claims["user_id"].(float64)
		if ok {
			return uint(uid), nil
		}
	}
	return 0, fmt.Errorf("failed to extract user_id")
}

func GenerateToken(userID uint, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"expiration": time.Now().Add(expiration).Unix(),
	})
	return token.SignedString(jwtSecret)
}
