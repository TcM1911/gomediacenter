package mongo

import (
	"log"

	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
)

// AddNewUser adds a new user to the Users collection.
func (d *DB) AddNewUser(user *gomediacenter.User) error {
	return d.session.DB(databaseName).C(userCollection).Insert(user)
}

// GetUserByID finds a user by it's id in the Users collection.
func (d *DB) GetUserByID(id gomediacenter.ID) (*gomediacenter.User, error) {
	q := d.session.DB(databaseName).C(userCollection).Find(bson.M{"id": id})
	var user gomediacenter.User
	if err := q.One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByName finds a user by the username in the Users collection.
func (d *DB) GetUserByName(name string) (*gomediacenter.User, error) {
	q := d.session.DB(databaseName).C(userCollection).Find(bson.M{"name": name})
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

	q := d.session.DB(databaseName).C(userCollection).Find(search)

	var users []*gomediacenter.User
	if err := q.All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteUser removes a user and user data from the database.
func (d *DB) DeleteUser(id gomediacenter.ID) error {
	// Remove all user item data.
	info, err := d.session.DB(databaseName).C(itemUserDataCollection).RemoveAll(
		bson.M{"uid": id})
	if err != nil {
		log.Println("Error when deleting user data:", err)
	}
	log.Println("Number of user data entries removed:", info.Removed)
	log.Println("Removing user", id.String())
	return d.session.DB(databaseName).C(userCollection).Remove(bson.M{"id": id})
}

// ChangeUserPassword changes the user's password.
func (d *DB) ChangeUserPassword(uid gomediacenter.ID, newPassword []byte) error {
	change := bson.M{"$set": bson.M{"password": newPassword, "haspassword": true}}
	return d.session.DB(databaseName).C(userCollection).Update(
		bson.M{"id": uid},
		change)
}

// UpdateUser updates the user profile stored in the database.
func (d *DB) UpdateUser(uid gomediacenter.ID, user *gomediacenter.User) error {
	var dbData gomediacenter.User
	c := d.session.DB(databaseName).C(userCollection)
	if err := c.Find(bson.M{"id": uid}).One(&dbData); err != nil {
		return err
	}

	// Ensure policy and configuration isn't changed.
	user.Policy = dbData.Policy
	user.Configuration = dbData.Configuration

	return c.Update(bson.M{"id": uid}, user)
}

// UpdateUserPolicy updates the user policy in the database.
func (d *DB) UpdateUserPolicy(id gomediacenter.ID, policy *gomediacenter.UserPolicy) error {
	ch := bson.M{"$set": bson.M{"policy": policy}}
	return d.session.DB(databaseName).C(userCollection).Update(bson.M{"id": id}, ch)
}

// UpdateUserConfiguration updates the user's configuration in the database.
func (d *DB) UpdateUserConfiguration(id gomediacenter.ID, cfg *gomediacenter.UserConfig) error {
	ch := bson.M{"$set": bson.M{"configuration": cfg}}
	return d.session.DB(databaseName).C(userCollection).Update(bson.M{"id": id}, ch)
}
