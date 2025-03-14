package helpers

import (
	"database/sql"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"github.com/pkg/errors"
)

type Resolver interface {
	Resolve(id, email string) (*domains.User, error)
}

type UserResolver struct {
	DbClient *sql.DB
}

func NewUserResolver(db *sql.DB) *UserResolver {
	return &UserResolver{DbClient: db}
}

func (ur UserResolver) Resolve(id, email string) (*domains.User, error) {
	var user domains.User
	query := `SELECT * FROM users WHERE id = $1 AND email = $2`
	err := ur.DbClient.QueryRow(query, id, email).Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, errors.New("user not found in db")
	}

	return &user, err
}
