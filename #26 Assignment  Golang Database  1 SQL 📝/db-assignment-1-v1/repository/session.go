package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"fmt"
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
	expiryStr := session.Expiry.Format("2006-01-02 15:04:05.999999999")
	add := fmt.Sprintf("INSERT INTO sessions (token, username, expiry) VALUES('%s', '%s', '%s')", session.Token, session.Username, expiryStr)
	_, err := u.db.Exec(add)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) DeleteSession(token string) error {
	delete := fmt.Sprintf("DELETE FROM sessions WHERE token='%s'", token)
	_, err := u.db.Exec(delete)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	update := fmt.Sprintf("UPDATE sessions SET token='%s' WHERE username='%s';", session.Token, session.Username)
	_, err := u.db.Exec(update)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailName(name string) error {
	names := fmt.Sprintf("SELECT * FROM sessions WHERE username='%s'", name)
	rows := u.db.QueryRow(names)
	var session model.Session
	err := rows.Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	tokens := fmt.Sprintf("SELECT * FROM sessions WHERE token='%s'", token)
	rows := u.db.QueryRow(tokens)
	var session model.Session
	err := rows.Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return model.Session{}, err
	}
	return session, nil // TODO: replace this
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
