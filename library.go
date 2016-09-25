package gomediacenter

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Library is a struct holding the data for a library.
type Library struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	LastScanned time.Time     `bson:"last_scanned"`
	Type        MEDIATYPE     `bson:"type"`
}
