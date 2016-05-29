package db

import "github.com/tcm1911/gomediacenter"

// AddNewUser adds a new user to the Users collection.
func (d *DB) AddNewUser(user *gomediacenter.User) error {
	return d.session.DB(DATABASE_NAME).C(USER_COLLECTION).Insert(user)
}
