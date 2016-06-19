package db

import (
	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
)

// SaveSessions saves the sessions to the database.
func (d *DB) SaveSessions(sessions []*gomediacenter.Session) error {
	for _, session := range sessions {
		_, err := d.session.DB(DATABASE_NAME).C(SESSIONS).UpsertId(session.Id, session)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSavedSessions gets the saved sessions from the database.
func (d *DB) GetSavedSessions() ([]*gomediacenter.Session, error) {
	var sessions []*gomediacenter.Session
	if err := d.session.DB(DATABASE_NAME).C(SESSIONS).Find(nil).All(&sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

// RemoveSession removes a session from the database.
func (d *DB) RemoveSession(id string) error {
	return d.session.DB(DATABASE_NAME).C(SESSIONS).RemoveId(bson.ObjectIdHex(id))
}