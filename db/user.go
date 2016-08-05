package db

import (
	"log"

	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
)

// AddNewUser adds a new user to the Users collection.
func (d *DB) AddNewUser(user *gomediacenter.User) error {
	return d.session.DB(DATABASE_NAME).C(USER_COLLECTION).Insert(user)
}

// GetUserById finds a user by it's id in the Users collection.
func (d *DB) GetUserByID(hexString string) (*gomediacenter.User, error) {
	q := d.session.DB(DATABASE_NAME).C(USER_COLLECTION).FindId(bson.ObjectIdHex(hexString))
	var user gomediacenter.User
	if err := q.One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByName finds a user by the username in the Users collection.
func (d *DB) GetUserByName(name string) (*gomediacenter.User, error) {
	q := d.session.DB(DATABASE_NAME).C(USER_COLLECTION).Find(bson.M{"name": name})
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

// DeleteUser removes a user and user data from the database.
func (d *DB) DeleteUser(hexString string) error {
	// Remove all user item data.
	info, err := d.session.DB(DATABASE_NAME).C(ITEM_USER_DATA_COLLECTION).RemoveAll(
		bson.M{"uid": hexString})
	if err != nil {
		log.Println("Error when deleting user data:", err)
	}
	log.Println("Number of user data entries removed:", info.Removed)
	log.Println("Removing user", hexString)
	return d.session.DB(DATABASE_NAME).C(USER_COLLECTION).RemoveId(bson.ObjectIdHex(hexString))
}

// ChangeUserPassword changes the user's password.
func (d *DB) ChangeUserPassword(uid string, newPassword []byte) error {
	change := bson.M{"$set": bson.M{"password": newPassword, "haspassword": true}}
	return d.session.DB(DATABASE_NAME).C(USER_COLLECTION).UpdateId(
		bson.ObjectIdHex(uid),
		change)
}

// UpdateUser updates the user profile stored in the database.
func (d *DB) UpdateUser(uid string, user *gomediacenter.User) error {
	var dbData gomediacenter.User
	c := d.session.DB(DATABASE_NAME).C(USER_COLLECTION)
	id := bson.ObjectIdHex(uid)
	if err := c.FindId(id).One(&dbData); err != nil {
		return err
	}

	// Ensure policy and configuration isn't changed.
	user.Policy = dbData.Policy
	user.Configuration = dbData.Configuration

	return c.UpdateId(id, user)
}

// UpdateUserPolicy updates the user policy in the database.
func (d *DB) UpdateUserPolicy(ID string, policy *gomediacenter.UserPolicy) error {
	ch := bson.M{"$set": bson.M{"policy": policy}}
	return d.session.DB(DATABASE_NAME).C(USER_COLLECTION).UpdateId(bson.ObjectIdHex(ID), ch)
}

// UpdateUserConfiguration updates the user's configuration in the database.
func (d *DB) UpdateUserConfiguration(ID string, cfg *gomediacenter.UserConfig) error {
	ch := bson.M{"$set": bson.M{"configuration": cfg}}
	return d.session.DB(DATABASE_NAME).C(USER_COLLECTION).UpdateId(bson.ObjectIdHex(ID), ch)
}
