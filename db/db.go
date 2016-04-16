package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

type DB struct {
	session *mgo.Session
}

var db *DB

// Connect connects to a database and returns a session.
func Connect(host string) {
	if db == nil {
		session, err := mgo.Dial(host)
		if err != nil {
			log.Fatalln(err)
		}
		db = &DB{session: session}
	}
}

// Close closes the database connection.
func (d *DB) Close() {
	db.session.Close()
}

// GetDBSession returns a copy of the database session.
func (d *DB) GetDBSession() *DB {
	if d == nil {
		log.Fatalln("Not connected to a database.")
	}
	copy := d.session.Copy()
	return &DB{session: copy}
}
