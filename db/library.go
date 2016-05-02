package db

import (
	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Library struct {
	Id          bson.ObjectId           `bson:"_id"`
	Name        string                  `bson:"name"`
	LastScanned time.Time               `bson:"last_scanned"`
	Type        gomediacenter.MEDIATYPE `bson:"type"`
}

func NewLibrary(name string, libraryType gomediacenter.MEDIATYPE) (*Library, error) {
	id := bson.NewObjectId()
	library := &Library{Name: name, Id: id, Type: libraryType}

	session := GetDBSession()
	defer session.Close()
	err := session.DB(DATABASE_NAME).C(LIBRARY_COLLECTION).Insert(library)
	if err != nil {
		return library, err
	}
	return library, nil
}

func GetLibraryById(id bson.ObjectId) (*Library, error) {
	session := GetDBSession()
	defer session.Close()
	q := session.DB(DATABASE_NAME).C(LIBRARY_COLLECTION).FindId(id)
	var lib Library
	if err := q.One(&lib); err != nil {
		return &lib, err
	}
	return &lib, nil
}

func UpdateLibraryLastScannedTime(id bson.ObjectId, time time.Time) error {
	session := GetDBSession()
	defer session.Close()

	if err := session.DB(DATABASE_NAME).C(LIBRARY_COLLECTION).UpdateId(id, bson.M{"$set": bson.M{"last_scanned": time}}); err != nil {
		return err
	}
	return nil
}
