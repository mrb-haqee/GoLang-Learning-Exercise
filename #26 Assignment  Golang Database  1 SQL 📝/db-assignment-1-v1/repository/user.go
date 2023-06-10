package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	add := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%s', '%s')",user.Username, user.Password)
	_, err := u.db.Exec(add)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	check := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", user.Username, user.Password)
	row := u.db.QueryRow(check)

	var users model.User
	err := row.Scan(&users.ID, &users.Username, &users.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}
	
	return nil // TODO: replace this
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
