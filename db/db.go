package db

import (
	"errors"
	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2"
	"labix.org/v2/mgo/bson"
	"log"
)

///////////////
// Constants //
///////////////
const (
	database_name             = "gomedia"
	mediatypes_collection     = "mediatypes"
	movie_collection          = "movies"
	item_user_data_collection = "itemuserdata"
)

/////////////
// Structs //
/////////////

type DB struct {
	session *mgo.Session
}

var db *DB

type mediaType struct {
	id    bson.ObjectId `bson:"_id"`
	media gomediacenter.MEDIATYPE
}

////////////
// Public //
////////////

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

// FindItemById finds an item from the database. It also returns the media type.
func (d *DB) FindItemById(id string) (gomediacenter.MEDIATYPE, interface{}, error) {
	q := d.session.DB(database_name).C(mediatypes_collection).FindId(bson.ObjectIdHex(id))

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

// FindItemUserData gets the user data for an item.
func (d *DB) FindItemUserData(uid, itemId string) (*gomediacenter.ItemUserData, error) {
	q := d.session.DB(database_name).C(item_user_data_collection).Find(bson.M{"uid": uid, "id": itemId})

	var itemUserData *gomediacenter.ItemUserData
	err := q.One(&itemUserData)
	if err == mgo.ErrNotFound {
		return gomediacenter.NewItemUserData(itemId, uid), nil
	}
	if err != nil {
		return itemUserData, err
	}
	return itemUserData, nil
}

// FindItemIntro returns intros for an item.
func (d *DB) FindItemIntro(id string) (*[]gomediacenter.Intro, error) {
	// TODO: Add support for intros
	var intros []gomediacenter.Intro
	return &intros, mgo.ErrNotFound
}

/////////////
// Private //
/////////////

func findMovieById(d *DB, id string) (*gomediacenter.Movie, error) {
	q := d.session.DB(database_name).C(movie_collection).FindId(bson.ObjectIdHex(id))
	var movie *gomediacenter.Movie
	if err := q.One(&movie); err != nil {
		return movie, err
	}
	return movie, nil
}
