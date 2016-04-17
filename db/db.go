package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"github.com/tcm1911/gomediacenter"
	"labix.org/v2/mgo/bson"
	"errors"
)

type DB struct {
	session *mgo.Session
}

var db *DB

type mediaType struct {
	id    bson.ObjectId `bson:"_id"`
	media gomediacenter.MEDIATYPE
}

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

func (d *DB) FindItemById(id string) (gomediacenter.MEDIATYPE, interface{}, error) {
	q := d.session.DB("gomedia").C("mediatypes").FindId(bson.ObjectIdHex(id))

	var mediatype mediaType
	if err := q.One(&mediatype); err != nil {
		return mediatype.media, nil, err
	}

	if mediatype.media == gomediacenter.MOVIE {
		movie, err := findMovieById(d, id)
		if err != nil {
			return mediatype.media, nil, err
		}
		return mediatype.media, movie, nil
	}
	return mediatype.media, nil, errors.New("no match")
}

func findMovieById(d *DB, id string) (*gomediacenter.Movie, error) {
	q := d.session.DB("gomedia").C("movies").FindId(bson.ObjectIdHex(id))
	var movie *gomediacenter.Movie
	if err := q.One(&movie); err != nil {
		return movie, err
	}
	return movie, nil
}