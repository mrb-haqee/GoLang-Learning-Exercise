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
}

type studentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var listStudent []model.Student

	rows, err := s.db.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var student model.Student
		err = rows.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
		if err != nil {
			return nil, err
		}

		listStudent = append(listStudent, student)
	}
	return listStudent, nil // TODO: replace this
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
	add := fmt.Sprintf("INSERT INTO students (name, address, class) VALUES ('%s','%s','%s')", student.Name, student.Address, student.Class)
	_, err := s.db.Exec(add)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	update := fmt.Sprintf("UPDATE students SET name='%s', address='%s', class='%s' WHERE id=%d;", student.Name, student.Address, student.Class, id)
	_, err := s.db.Exec(update)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
