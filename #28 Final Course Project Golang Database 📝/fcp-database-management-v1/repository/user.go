package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	return u.db.Create(&user).Error // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	if user==(model.User{}){
		return errors.New("invalid User")
	}
	var users model.User
	return u.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&users).Error // TODO: replace this
}
