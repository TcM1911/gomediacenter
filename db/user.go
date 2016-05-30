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

// GetAllUsers returns all the users. The list can be filtered using the filter map.
func (d *DB) GetAllUsers(filter map[string]interface{}) ([]*gomediacenter.User, error) {

	terms := []bson.M{}
	for k, v := range filter {
		term := bson.M{"policy." + k: v}
		terms = append(terms, term)
	}

	var search interface{}
	if len(terms) == 0 {
		search = nil
	} else {
		search = bson.M{"$and": terms}
	}

	q := d.session.DB(DATABASE_NAME).C(USER_COLLECTION).Find(search)

	var users []*gomediacenter.User
	if err := q.All(&users); err != nil {
		return nil, err
	}
	return users, nil
}
