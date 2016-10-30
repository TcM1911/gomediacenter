package mongo

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
	databaseName           = "gomedia"
	mediatypeCollection    = "mediatypes"
	movieCollection        = "movies"
	userCollection         = "users"
	itemUserDataCollection = "itemuserdata"
	libraryCollection      = "library"
	sessionsCollection     = "sessions"
)

/////////////
// Structs //
/////////////

// DB is the MongoDB implementation of the gomediacenter.DB interface.gomediacenter.
type DB struct {
	session *mgo.Session
}

type mediaType struct {
	ID    gomediacenter.ID        `bson:"id"`
	Media gomediacenter.MEDIATYPE `bson:"media"`
}

////////////
// Public //
////////////

// Connect connects to a database and returns a session.
func (d *DB) Connect(host string) {
	if d.session == nil {
		session, err := mgo.Dial(host)
		if err != nil {
			log.Fatalln(err)
		}
		d.session = session
	}
}

// Close closes the database connection.
func (d *DB) Close() {
	d.session.Close()
}

// FindItemByID finds an item from the database. It also returns the media type.
func (d *DB) FindItemByID(id gomediacenter.ID) (gomediacenter.MEDIATYPE, interface{}, error) {
	s := d.getDBSessionCopy()
	defer s.Close()
	q := getCollection(s, mediatypeCollection).Find(bson.M{"id": id})

	var mediatype mediaType
	if err := q.One(&mediatype); err != nil {
		return mediatype.Media, nil, err
	}

	if mediatype.Media == gomediacenter.MOVIE {
		movie, err := findMovieByID(d, id)
		if err != nil {
			return mediatype.Media, nil, err
		}
		return mediatype.Media, movie, nil
	}
	return mediatype.Media, nil, errors.New("no match")
}

// FindItemUserData gets the user data for an item.
func (d *DB) FindItemUserData(uid, itemID gomediacenter.ID) (*gomediacenter.ItemUserData, error) {
	s := d.getDBSessionCopy()
	defer s.Close()
	q := getCollection(s, itemUserDataCollection).Find(bson.M{"uid": uid, "id": itemID})

	var itemUserData *gomediacenter.ItemUserData
	err := q.One(&itemUserData)
	if err == mgo.ErrNotFound {
		// Return a new struct.
		data := gomediacenter.NewItemUserData(itemID, uid)
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
	err := d.session.DB(databaseName).C(itemUserDataCollection).Insert(userData)
	if err != nil {
		return err
	}
	return nil
}

// FindItemIntro returns intros for an item.
func (d *DB) FindItemIntro(id gomediacenter.ID) ([]*gomediacenter.Intro, error) {
	// TODO: Add support for intros
	var intros []*gomediacenter.Intro
	return intros, mgo.ErrNotFound
}

// InsertItemType inserts an item into the media type collection.
func (d *DB) InsertItemType(id gomediacenter.ID, gomediaType gomediacenter.MEDIATYPE) error {
	s := d.getDBSessionCopy()
	defer s.Close()

	dbMediaType := &mediaType{ID: id, Media: gomediaType}
	err := getCollection(s, mediatypeCollection).Insert(dbMediaType)
	if err != nil {
		return err
	}
	return nil
}

// RemoveItemType removes an item in the media type collection.
func (d *DB) RemoveItemType(id gomediacenter.ID) error {
	session := d.getDBSessionCopy()
	defer session.Close()

	err := getCollection(session, mediatypeCollection).Remove(bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

// getDBSession returns a copy of the database session.
func (d *DB) getDBSessionCopy() *mgo.Session {
	if d.session == nil {
		log.Fatalln("Not connected to a database.")
	}
	return d.session.Copy()
}

func getCollection(s *mgo.Session, c string) *mgo.Collection {
	return s.DB(databaseName).C(c)
}
