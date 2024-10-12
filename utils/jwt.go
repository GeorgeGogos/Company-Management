package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("supersecretkey")

// GenerateJWT generates a JWT for a given user
func GenerateJWT(username string) (string, error) {
	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(), // Token expiration time (72 hours)
	})

	// Sign the token using the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// IsValidJWT checks the validity of a JWT
func IsValidJWT(tokenString string) bool {
	// Parse the token and validate the signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}
