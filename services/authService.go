package services

import (
	"github.com/AnkitDhawale/TodoListApp/auth"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/AnkitDhawale/TodoListApp/helpers"
	"github.com/AnkitDhawale/TodoListApp/repositories"
)

type AuthService interface {
	Login(user *dto.User) (*dto.LoginResponse, error)
}

type DefaultAuthService struct {
	repo repositories.AuthRepo
}

func NewDefaultAuthService(repo repositories.AuthRepo) *DefaultAuthService {
	return &DefaultAuthService{repo}
}

func (service DefaultAuthService) Login(user *dto.User) (*dto.LoginResponse, error) {
	u, err := service.repo.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	_, err = helpers.IsPasswordCorrect(u.PasswordHash, user.Password)
	if err != nil {
		return nil, err
	}

	claims := u.ClaimsForUser()
	token, err := auth.NewAccessToken(claims)
	if err != nil {
		return nil, err
	}

	return token, nil
}
