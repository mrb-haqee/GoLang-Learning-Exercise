package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var listStudent []model.Student
	err := s.db.Table("students").Scan(&listStudent).Error
	if err != nil {
		return nil, err
	}
	return listStudent, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	return s.db.Create(student).Error // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	return s.db.Table("students").Where("id = ?", id).Updates(model.Student{Name: student.Name, Address: student.Address, ClassId: student.ClassId}).Error // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	return s.db.Delete(&model.Student{}, id).Error // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student *model.Student
	err := s.db.Table("students").Where("id = ?", id).Scan(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var student model.Student
	s.db.Table("students").Select("name, address, class_id").Scan(&student)
	if student == (model.Student{}) {
		return &[]model.StudentClass{}, nil
	}
	var studentClass *[]model.StudentClass
	err := s.db.Table("students").Select("students.name AS name, students.address AS address, classes.name AS class_name, classes.professor AS professor, classes.room_number AS room_number").
		Joins("INNER JOIN classes ON classes.id = students.class_id").Scan(&studentClass).Error
	if err != nil {
		return nil, err
	}

	return studentClass, nil // TODO: replace this
}
