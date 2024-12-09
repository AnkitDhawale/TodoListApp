package domains

import (
	"github.com/AnkitDhawale/TodoListApp/auth"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	Id           string    `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

type UserRepo interface {
	FindAll() ([]User, error)
	AddUser(user *User) (string, error)
	UpdateUser(userIdFromToken string, user *User) error
	GetUserById(id string) (*User, error)
}

func (u User) ToDto() dto.User {
	return dto.User{
		Email:    u.Email,
		Password: u.PasswordHash,
	}
}

func (u User) ClaimsForUser() *auth.AccessTokenClaims {
	return &auth.AccessTokenClaims{
		UserId: u.Id,
		Email:  u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "http://localhost:8888/todoapp",
			Subject:   u.Id,
			Audience:  []string{"my-todo-list-app"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(auth.ACCESS_TOKEN_DURATION)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
