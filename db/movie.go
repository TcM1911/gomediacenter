package db

import (
	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2"
	"labix.org/v2/mgo/bson"
	"log"
)

////////////
// Public //
////////////

// IsMovieFileAlreadyKnown looks if the file is already in the database. If it is, the id is returned
// as the second argument.
func IsMovieFileAlreadyKnown(relativePath string, parentId bson.ObjectId) (bool, bson.ObjectId) {
	session := GetDBSession()
	defer session.Close()

	var movieId bson.ObjectId

	search := bson.M{"ParentId": parentId, "Path": relativePath}
	q := session.DB(DATABASE_NAME).C(MOVIE_COLLECTION).Find(search)
	var movie *gomediacenter.Movie
	err := q.One(movie)

	if err == mgo.ErrNotFound {
		return false, movieId
	}
	if err != nil {
		// Log the error and treat like it's a new file.
		log.Println("Error when trying to find", relativePath, "in the database:", err)
		return false, movieId
	}
	return true, movie.Id
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
