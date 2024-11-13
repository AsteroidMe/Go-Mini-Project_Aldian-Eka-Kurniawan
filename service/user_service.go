package service

import (
	"eco-journal/entities"
	"eco-journal/middleware"
	"eco-journal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(user *entities.User) (*entities.User, error)
	Login(email, password string) (string, error)
}

type userService struct {
	userRepo repository.UserRepoInterface
}

func NewUserService(userRepo repository.UserRepoInterface) *userService {
	return &userService{userRepo}
}

func (s *userService) Register(user *entities.User) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	token, err := middleware.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
