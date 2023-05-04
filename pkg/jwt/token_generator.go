package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/your_project/internal/models"
)

type TokenGenerator interface {
	GenerateToken(user *models.User) (string, error)
}

type tokenGenerator struct {
	secret string
}

func NewTokenGenerator(secret string) TokenGenerator {
	return &tokenGenerator{secret: secret}
}

func (t *tokenGenerator) GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.secret))
}
