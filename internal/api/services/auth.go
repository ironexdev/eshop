package services

import (
	"errors"

	"github.com/your_project/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type TokenGenerator interface {
	GenerateToken(user *models.User) (string, error)
}

type AuthService struct {
	userRepo       UserRepository
	tokenGenerator TokenGenerator
}

func NewAuthService(userRepo UserRepository, tokenGenerator TokenGenerator) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		tokenGenerator: tokenGenerator,
	}
}

func (a *AuthService) Register(user *models.User) error {
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}
	return a.userRepo.Create(user)
}

func (a *AuthService) Login(email, password string) (*models.User, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, errors.New("Invalid email or password")
	}

	return user, nil
}

func (a *AuthService) GenerateToken(user *models.User) (string, error) {
	return a.tokenGenerator.GenerateToken(user)
}
