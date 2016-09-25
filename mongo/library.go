package mongo

import (
	"errors"
	"time"

	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
)

// NewLibrary creates a new library and returns a struct of it.
func (d *DB) NewLibrary(name string, libraryType gomediacenter.MEDIATYPE) (*gomediacenter.Library, error) {
	id := bson.NewObjectId()
	library := &gomediacenter.Library{Name: name, ID: id, Type: libraryType}

	session := d.getDBSessionCopy()
	defer session.Close()
	err := getCollection(session, libraryCollection).Insert(library)
	if err != nil {
		return library, err
	}
	return library, nil
}

// GetLibraryByID returns the library by id.
func (d *DB) GetLibraryByID(id bson.ObjectId) (*gomediacenter.Library, error) {
	session := d.getDBSessionCopy()
	defer session.Close()
	q := getCollection(session, libraryCollection).FindId(id)
	var lib gomediacenter.Library
	if err := q.One(&lib); err != nil {
		return &lib, err
	}
	return &lib, nil
}

// UpdateLibraryLastScannedTime updates the time the library was lasted scanned.
func (d *DB) UpdateLibraryLastScannedTime(id bson.ObjectId, time time.Time) error {
	session := d.getDBSessionCopy()
	defer session.Close()

	if err := getCollection(session, libraryCollection).UpdateId(id, bson.M{"$set": bson.M{"last_scanned": time}}); err != nil {
		return err
	}
	return nil
}

// PruneMissingItemsFromLibrary removes items not listed in the items map. The items map contains the relative path of the
// files to keep. The function returns a slice of the items removed.
func (d *DB) PruneMissingItemsFromLibrary(items map[string]struct{}, libID string, libType gomediacenter.MEDIATYPE) ([]string, error) {
	session := d.getDBSessionCopy()
	defer session.Close()

	var removedItems []string

	switch libType {
	case gomediacenter.MOVIE:
		c := getCollection(session, movieCollection)
		var result gomediacenter.Movie

		q := bson.M{"parentId": libID}
		iterator := c.Find(q).Iter()

		for iterator.Next(&result) {
			path := result.Path
			if _, exist := items[path]; !exist {
				if err := c.RemoveId(result.ID); err != nil {
					return removedItems, err
				}
				if err := d.RemoveItemType(result.ID); err != nil {
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
