package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepo struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) ClassRepo {
	return ClassRepo{db}
}

func (c ClassRepo) Init(data []model.Class) error {
	c.db.Create(&data)
	return nil // TODO: replace this
}
