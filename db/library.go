package db

import (
	"errors"
	"time"

	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Library struct {
	Id          bson.ObjectId           `bson:"_id"`
	Name        string                  `bson:"name"`
	LastScanned time.Time               `bson:"last_scanned"`
	Type        gomediacenter.MEDIATYPE `bson:"type"`
}

// Creates a new library and returns a struct of it.
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

// Returns the library by id.
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

// Updates the time the library was lasted scanned.
func UpdateLibraryLastScannedTime(id bson.ObjectId, time time.Time) error {
	session := GetDBSession()
	defer session.Close()

	if err := session.DB(DATABASE_NAME).C(LIBRARY_COLLECTION).UpdateId(id, bson.M{"$set": bson.M{"last_scanned": time}}); err != nil {
		return err
	}
	return nil
}

// Removes items not listed in the items map. The items map contains the relative path of the
// files to keep. The function returns a slice of the items removed.
func PruneMissingItemsFromLibrary(items map[string]struct{}, libId string, libType gomediacenter.MEDIATYPE) ([]string, error) {
	session := GetDBSession()
	defer session.Close()

	var removedItems []string

	db := session.DB(DATABASE_NAME)
	var C *mgo.Collection
	switch libType {
	case gomediacenter.MOVIE:
		C = db.C(MOVIE_COLLECTION)
		var result gomediacenter.Movie

		q := bson.M{"parentId": libId}
		iterator := C.Find(q).Iter()

		for iterator.Next(&result) {
			path := result.Path
			if _, exist := items[path]; !exist {
				if err := C.RemoveId(result.ID); err != nil {
					return removedItems, err
				}
				if err := RemoveItemType(result.ID); err != nil {
					return removedItems, err
				}
				removedItems = append(removedItems, path)
			}
		}
	default:
		return removedItems, errors.New("incorrect collection type")
	}
	return removedItems, nil
}
