package services

import (
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/AnkitDhawale/TodoListApp/repositories"
	"testing"
)

func setup(t *testing.T) {
	//mockService := NewDefaultAuthService()
}

func TestDefaultAuthService_Login(t *testing.T) {
	tests := []struct {
		name         string
		inputUser    *dto.User
		expectedResp *dto.LoginResponse
		expectedErr  error
	}{
		{
			name: "failure: invalid email",
			inputUser: &dto.User{
				Email:    "test@test.com",
				Password: "",
			},
			expectedResp: nil,
			expectedErr:  repositories.ErrInvalidEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
