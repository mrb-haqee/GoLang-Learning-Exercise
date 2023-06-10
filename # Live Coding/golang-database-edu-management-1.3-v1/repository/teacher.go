package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"fmt"
)

type TeacherRepository interface {
	FetchByID(id int) (*model.Teacher, error)
	Store(g *model.Teacher) error
	Delete(id int) error
}

type teacherRepoImpl struct {
	db *sql.DB
}

func NewTeacherRepo(db *sql.DB) *teacherRepoImpl {
	return &teacherRepoImpl{db}
}

func (g *teacherRepoImpl) FetchByID(id int) (*model.Teacher, error) {
	row := g.db.QueryRow("SELECT id, name, address, subject FROM teachers WHERE id = $1", id)

	var Teacher model.Teacher
	err := row.Scan(&Teacher.ID, &Teacher.Name, &Teacher.Address, &Teacher.Subject)
	if err != nil {
		return nil, err
	}

	return &Teacher, nil
}

func (g *teacherRepoImpl) Store(teacher *model.Teacher) error {
	add := fmt.Sprintf("INSERT INTO teachers (name, address, subject) VALUES ('%s','%s','%s')", teacher.Name, teacher.Address, teacher.Subject)
	_, err := g.db.Exec(add)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (g *teacherRepoImpl) Delete(id int) error {
	delete := fmt.Sprintf("DELETE FROM teachers WHERE id=%d;", id)
	_, err := g.db.Exec(delete)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
