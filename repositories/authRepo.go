package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/jackc/pgx"
)

type AuthRepo interface {
	FindUserBy(user *dto.User) (*domains.User, error)
	FindByEmail(email string) (*domains.User, error)
}

type AuthRepoDb struct {
	dbClient *pgx.Conn
}

func NewAuthRepoDb(db *pgx.Conn) AuthRepoDb {
	return AuthRepoDb{dbClient: db}
}

func (auth AuthRepoDb) FindUserBy(user *dto.User) (*domains.User, error) {
	query := `SELECT * FROM users WHERE email = $1 AND password_hash = $2`

	var usr domains.User
	err := auth.dbClient.QueryRow(query, user.Email, user.Password).
		Scan(&usr.Id, &usr.Email, &usr.PasswordHash, &usr.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid credentials")
		} else {
			return nil, fmt.Errorf("unexpected database error: %v", err)
		}
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
			return nil, errors.New("invalid email")
		} else {
			return nil, fmt.Errorf("unexpected database error: %v", err)
		}
	}

	return &usr, nil
}
