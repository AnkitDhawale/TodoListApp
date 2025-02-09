package repositories

import (
	"database/sql"
	"errors"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"log"
)

// ErrInvalidCredentials defines a global error for consistent comparison
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrUnexpected         = errors.New("unexpected database error")
)

type AuthRepo interface {
	FindUserBy(user *dto.User) (*domains.User, error)
	FindByEmail(email string) (*domains.User, error)
}

type AuthRepoDb struct {
	dbClient *sql.DB
}

func NewAuthRepoDb(db *sql.DB) AuthRepoDb {
	return AuthRepoDb{dbClient: db}
}

func (auth AuthRepoDb) FindUserBy(user *dto.User) (*domains.User, error) {
	query := `SELECT * FROM users WHERE email = $1 AND password_hash = $2`

	var usr domains.User
	err := auth.dbClient.QueryRow(query, user.Email, user.Password).
		Scan(&usr.Id, &usr.Email, &usr.PasswordHash, &usr.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}

		log.Println("unexpected database error: %w", err)
		return nil, ErrUnexpected
	}

	return &usr, nil
}

func (auth AuthRepoDb) FindByEmail(email string) (*domains.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var usr domains.User
	err := auth.dbClient.QueryRow(query, email).
		Scan(&usr.Id, &usr.Email, &usr.PasswordHash, &usr.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidEmail
		} else {
			log.Println("unexpected database error: %w", err)
			return nil, ErrUnexpected
		}
	}

	return &usr, nil
}
