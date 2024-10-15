package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var key = []byte("hmmm")

type Claims struct {
	User string `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(user string) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return key, nil })
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
