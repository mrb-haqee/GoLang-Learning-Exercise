package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)

	FetchByID(id int) (*model.Session, error)
}

type sessionsRepoImpl struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (u *sessionsRepoImpl) AddSessions(session model.Session) error {
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) DeleteSession(token string) error {
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailName(name string) error {
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	return model.Session{}, nil // TODO: replace this
}

func (u *sessionsRepoImpl) FetchByID(id int) (*model.Session, error) {
	row := u.db.QueryRow("SELECT id, token, username, expiry FROM sessions WHERE id = $1", id)

	var session model.Session
	err := row.Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
