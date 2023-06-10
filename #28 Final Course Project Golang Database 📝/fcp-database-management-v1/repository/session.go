package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	return s.db.Create(&session).Error // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	return s.db.Where("token = ?", token).Delete(&model.Session{}).Error // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	return s.db.Table("sessions").Where("username = ? AND expiry = ?", session.Username, session.Expiry).Update("token", session.Token).Error // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	var session model.Session
	return s.db.Where("username = ?", name).First(&session).Error // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := s.db.Where("token = ?", token).First(&session).Error
	if err != nil {
		return model.Session{}, err
	}
	return session, nil // TODO: replace this
}
