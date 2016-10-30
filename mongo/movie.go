package mongo

import (
	"log"

	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
)

////////////
// Public //
////////////

// InsertNewMovie adds a new movie to the database. The function returns true if the insertion
// was successful.
func (d *DB) InsertNewMovie(movie *gomediacenter.Movie) (gomediacenter.ID, error) {
	id := gomediacenter.NewID()
	movie.ID = id
	session := d.getDBSessionCopy()
	defer session.Close()

	err := getCollection(session, movieCollection).Insert(movie)
	if err != nil {
		return id, err
	}

	if err := d.InsertItemType(id, gomediacenter.MOVIE); err != nil {
		return id, err
	}

	return id, nil
}

// IsMovieFileAlreadyKnown looks if the file is already in the database. If it is, the id is returned
// as the second argument.
func (d *DB) IsMovieFileAlreadyKnown(relativePath string, parentID gomediacenter.ID) bool {
	session := d.getDBSessionCopy()
	defer session.Close()

	search := bson.M{"path": relativePath, "parentId": parentID}
	n, err := getCollection(session, movieCollection).Find(search).Count()
	if err != nil {
		// Log the error and treat like it's a new file.
		log.Println("Error when trying to find", relativePath, "in the database:", err)
		return false
	}
	if n > 0 {
		return true
	}
	return false
}

/////////////
// Private //
/////////////

func findMovieByID(d *DB, id gomediacenter.ID) (*gomediacenter.Movie, error) {
	q := d.session.DB(databaseName).C(movieCollection).Find(bson.M{"id": id})
	var movie *gomediacenter.Movie
	if err := q.One(&movie); err != nil {
		return movie, err
	}
	return movie, nil
}
