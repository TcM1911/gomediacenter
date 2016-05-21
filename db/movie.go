package db

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
func InsertNewMovie(movie *gomediacenter.Movie) (bson.ObjectId, error) {
	id := bson.NewObjectId()
	movie.Id = id
	session := GetDBSession()
	defer session.Close()

	err := session.DB(DATABASE_NAME).C(MOVIE_COLLECTION).Insert(movie)
	if err != nil {
		return id, err
	}
	return id, nil
}

// IsMovieFileAlreadyKnown looks if the file is already in the database. If it is, the id is returned
// as the second argument.
func IsMovieFileAlreadyKnown(relativePath string, parentId bson.ObjectId) bool {
	session := GetDBSession()
	defer session.Close()

	search := bson.M{"path": relativePath, "parentId": parentId.Hex()}
	n, err := session.DB(DATABASE_NAME).C(MOVIE_COLLECTION).Find(search).Count()
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

func findMovieById(d *DB, id string) (*gomediacenter.Movie, error) {
	q := d.session.DB(DATABASE_NAME).C(MOVIE_COLLECTION).FindId(bson.ObjectIdHex(id))
	var movie *gomediacenter.Movie
	if err := q.One(&movie); err != nil {
		return movie, err
	}
	return movie, nil
}
