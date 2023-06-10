package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"fmt"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
}

type studentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	return []model.Student{}, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	row := s.db.QueryRow("SELECT id, name, address, class FROM students WHERE id = $1", id)

	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	return fmt.Errorf("Replace this error!") // TODO: replace this
}
