package gomediacenter

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Database is the interface for the database.
type Database interface {
	Connect(host string)
	Close()
	//	ItemFinder
	//	ItemUserData
	//	IntroFinder
}

// ItemFinder can find an item in the database.
type ItemFinder interface {
	FindItemByID(id string) (MEDIATYPE, interface{}, error)
	FindItemUserData(uid, itemID string) (*ItemUserData, error)
}

// IntroFinder can find intros for an item.
type IntroFinder interface {
	FindItemIntro(id string) ([]*Intro, error)
}

// ItemUserSaver can find user data for an item.
type ItemUserSaver interface {
	InsertItemUserData(userData *ItemUserData) error
}

// UserManager can create, update, and remove user profiles.
type UserManager interface {
	AddNewUser(user *User) error
	UpdateUser(ID string, user *User) error
	UpdateUserPolicy(ID string, policy *UserPolicy) error
	UpdateUserConfiguration(ID string, config *UserConfig) error
	GetUserByID(hexString string) (*User, error)
	GetUserByName(name string) (*User, error)
	GetAllUsers(filter map[string]interface{}) ([]*User, error)
	DeleteUser(hexString string) error
	ChangeUserPassword(uid string, newPassword []byte) error
}

// SessionSaver can save and get stored sessions from the database.
type SessionSaver interface {
	GetSavedSessions() ([]*Session, error)
	SaveSessions([]*Session) error
}

type LibraryMaintainer interface {
	NewLibrary(name string, libraryType MEDIATYPE) (*Library, error)
	GetLibraryByID(id bson.ObjectId) (*Library, error)
	UpdateLibraryLastScannedTime(id bson.ObjectId, time time.Time) error
	PruneMissingItemsFromLibrary(items map[string]struct{}, libID string, libType MEDIATYPE) ([]string, error)
	MovieMaintainer
}

type MovieMaintainer interface {
	IsMovieFileAlreadyKnown(path string, parentID bson.ObjectId) bool
	InsertNewMovie(movie *Movie) (bson.ObjectId, error)
}
