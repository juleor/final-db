package repository

import (
	"a21hc3NpZ25tZW50/model"

	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	u.db.First(&session)

	addsess := u.db.Save(&session)
	if addsess != nil {
		return addsess.Error
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	del := u.db.Delete(&tokenTarget)
	if del.Error != nil {
		return del.Error
	}

	return nil // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	upd := u.db.Update(session.Username, &SessionsRepository{})
	if upd.Error != nil {
		return upd.Error
	}
	return nil // TODO: replace this

}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	// u.db.First(&token)
	// if len(token) == 0 {
	// 	return model.Session{}, u.db.Error
	// }
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}
	if u.TokenExpired(session) {
		err := u.DeleteSessions(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, u.db.Error
	}

	return model.Session{}, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	u.db.First(&name)
	if len(u.db.Name()) == 0 {
		return model.Session{}, u.db.Error
	}
	return model.Session{}, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	u.db.First(&token)
	if len(token) == 0 {
		return model.Session{}, u.db.Error

	}

	// if token != u.TokenExpired(token)

	return model.Session{}, nil // TODO: replace this
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
