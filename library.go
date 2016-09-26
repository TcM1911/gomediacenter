package gomediacenter

import "time"

// Library is a struct holding the data for a library.
type Library struct {
	ID          ID        `bson:"id"`
	Name        string    `bson:"name"`
	LastScanned time.Time `bson:"last_scanned"`
	Type        MEDIATYPE `bson:"type"`
}
