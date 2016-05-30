package db

import (
	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
)

// AddNewUser adds a new user to the Users collection.
func (d *DB) AddNewUser(user *gomediacenter.User) error {
	return d.session.DB(DATABASE_NAME).C(USER_COLLECTION).Insert(user)
}

// GetUserById finds a user by it's id in the Users collection.
func (d *DB) GetUserById(hexString string) (*gomediacenter.User, error) {
	q := d.session.DB(DATABASE_NAME).C(USER_COLLECTION).FindId(bson.ObjectIdHex(hexString))
	var user gomediacenter.User
	if err := q.One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
