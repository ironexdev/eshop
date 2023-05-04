package utils

import (
	"errors"

	"github.com/form3tech-oss/jwt-go"
)

type JWTTokenParser struct {
	secret string
}

func NewJWTTokenParser(secret string) *JWTTokenParser {
	return &JWTTokenParser{secret: secret}
}

func (p *JWTTokenParser) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(p.secret), nil
	})

	return token, err
}
