package repositories

import (
	"database/sql"
	"fmt"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"log"
)

type UserRepoDb struct {
	dbClient *sql.DB
}

func NewUserRepoDb(db *sql.DB) UserRepoDb {
	return UserRepoDb{dbClient: db}
}

func (u UserRepoDb) FindAll() ([]domains.User, error) {
	rows, err := u.dbClient.Query("select * from users")
	if err != nil {
		log.Println("error while getting rows", err)

		return nil, err
	}
	defer rows.Close()

	var users []domains.User
	for rows.Next() {
		var u domains.User
		err = rows.Scan(
			&u.Id,
			&u.Email,
			&u.PasswordHash,
			&u.CreatedAt,
		)
		if err != nil {
			log.Println("error while scanning data", err)

			return nil, err
		}
		users = append(users, u)
	}

	//fmt.Println(users)

	return users, nil
}

func (u UserRepoDb) AddUser(user *domains.User) (string, error) {
	query := `INSERT INTO users (id, email, password_hash, created_at) 
			  VALUES ($1, $2, $3, $4)`

	res, err := u.dbClient.Exec(query, user.Id, user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		log.Println("failed to create a user, err: ", err)
		return "", err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Println("failed to create a user, err: ", err)
		return "", err
	}

	log.Println(affectedRows)

	return user.Id, nil
}

func (u UserRepoDb) UpdateUser(userIdFromToken string, user *domains.User) error {
	var err error
	updateQuery := `UPDATE users SET email = $1, password_hash = $2 WHERE id = $3`
	updateEmailQuery := `UPDATE users SET email = $1 WHERE id = $2`
	updatePasswordQuery := `UPDATE users SET password_hash = $1 WHERE id = $2`

	var res sql.Result

	switch {
	case user.Email != "" && user.PasswordHash == "":
		// update email address
		res, err = u.dbClient.Exec(updateEmailQuery, user.Email, userIdFromToken)
	case user.Email == "" && user.PasswordHash != "":
		// update password
		res, err = u.dbClient.Exec(updatePasswordQuery, user.PasswordHash, userIdFromToken)
	default:
		// update both
		res, err = u.dbClient.Exec(updateQuery, user.Email, user.PasswordHash, userIdFromToken)
	}

	if err != nil {
		log.Println("failed to update user, err: ", err)
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Println("failed to create a user, err: ", err)
		return err
	}

	log.Println(affectedRows)

	return nil
}

func (u UserRepoDb) GetUserById(id string) (*domains.User, error) {
	var user domains.User
	getQuery := `SELECT * FROM users WHERE id = $1`
	err := u.dbClient.QueryRow(getQuery, id).
		Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		log.Printf("failed to get user with id: %s, error: %v\n", id, err)

		return nil, fmt.Errorf("invalid userId: %s, err: %v", id, err)
	} else {
		return &user, nil
	}
}
