package db

import (
	"errors"
	"log"

	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

///////////////
// Constants //
///////////////
const (
	DATABASE_NAME             = "gomedia"
	MEDIATYPE_COLLECTION      = "mediatypes"
	MOVIE_COLLECTION          = "movies"
	USER_COLLECTION           = "users"
	ITEM_USER_DATA_COLLECTION = "itemuserdata"
	LIBRARY_COLLECTION        = "library"
)

/////////////
// Structs //
/////////////

type DB struct {
	session *mgo.Session
}

var db *DB

type mediaType struct {
	Id    bson.ObjectId           `bson:"_id"`
	Media gomediacenter.MEDIATYPE `bson:"media"`
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

func Disconnect() {
	if db != nil {
		db.Close()
		db = nil
	}
}

// Close closes the database connection.
func (d *DB) Close() {
	d.session.Close()
}

func GetDB() *DB {
	newSession := GetDBSession()
	return &DB{session: newSession}
}

// GetDBSession returns a copy of the database session.
func GetDBSession() *mgo.Session {
	if db == nil {
		log.Fatalln("Not connected to a database.")
	}
	return db.session.Copy()
}

// FindItemById finds an item from the database. It also returns the media type.
func (d *DB) FindItemById(id string) (gomediacenter.MEDIATYPE, interface{}, error) {
	q := d.session.DB(DATABASE_NAME).C(MEDIATYPE_COLLECTION).FindId(bson.ObjectIdHex(id))

	var mediatype mediaType
	if err := q.One(&mediatype); err != nil {
		return mediatype.Media, nil, err
	}

	if mediatype.Media == gomediacenter.MOVIE {
		movie, err := findMovieById(d, id)
		if err != nil {
			return mediatype.Media, nil, err
		}
		return mediatype.Media, movie, nil
	}
	return mediatype.Media, nil, errors.New("no match")
}

// FindItemUserData gets the user data for an item.
func (d *DB) FindItemUserData(uid, itemId string) (*gomediacenter.ItemUserData, error) {
	q := d.session.DB(DATABASE_NAME).C(ITEM_USER_DATA_COLLECTION).Find(bson.M{"uid": uid, "id": itemId})

	var itemUserData *gomediacenter.ItemUserData
	err := q.One(&itemUserData)
	if err == mgo.ErrNotFound {
		// Return a new struct.
		data := gomediacenter.NewItemUserData(itemId, uid)
		if err := d.InsertItemUserData(data); err != nil {
			return data, err
		}
		return data, nil
	}
	if err != nil {
		return nil, err
	}
	return itemUserData, nil
}

// InsertItemUserData inserts a new record of user item data.
func (d *DB) InsertItemUserData(userData *gomediacenter.ItemUserData) error {
	err := d.session.DB(DATABASE_NAME).C(ITEM_USER_DATA_COLLECTION).Insert(userData)
	if err != nil {
		return err
	}
	return nil
}

// FindItemIntro returns intros for an item.
func (d *DB) FindItemIntro(id string) ([]*gomediacenter.Intro, error) {
	// TODO: Add support for intros
	var intros []*gomediacenter.Intro
	return intros, mgo.ErrNotFound
}

// Inserts an item into the media type collection.
func InsertItemType(id bson.ObjectId, gomediaType gomediacenter.MEDIATYPE) error {
	session := GetDBSession()
	defer session.Close()

	dbMediaType := &mediaType{Id: id, Media: gomediaType}
	err := session.DB(DATABASE_NAME).C(MEDIATYPE_COLLECTION).Insert(dbMediaType)
	if err != nil {
		return err
	}
	return nil
}

// Removes an item in the media type collection.
func RemoveItemType(id bson.ObjectId) error {
	session := GetDBSession()
	defer session.Close()

	err := session.DB(DATABASE_NAME).C(MEDIATYPE_COLLECTION).RemoveId(id)
	if err != nil {
		return err
	}
	return nil
}
