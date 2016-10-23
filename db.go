package gomediacenter

import "time"

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
	FindItemByID(id ID) (MEDIATYPE, interface{}, error)
	FindItemUserData(uid, itemID ID) (*ItemUserData, error)
}

// IntroFinder can find intros for an item.
type IntroFinder interface {
	FindItemIntro(id ID) ([]*Intro, error)
}

// ItemUserSaver can find user data for an item.
type ItemUserSaver interface {
	InsertItemUserData(userData *ItemUserData) error
}

// UserManager can create, update, and remove user profiles.
type UserManager interface {
	AddNewUser(user *User) error
	UpdateUser(ID ID, user *User) error
	UpdateUserPolicy(ID ID, policy *UserPolicy) error
	UpdateUserConfiguration(ID ID, config *UserConfig) error
	GetUserByID(ID ID) (*User, error)
	GetUserByName(name string) (*User, error)
	GetAllUsers(filter map[string]interface{}) ([]*User, error)
	DeleteUser(ID ID) error
	ChangeUserPassword(uid ID, newPassword []byte) error
}

type LibraryMaintainer interface {
	NewLibrary(name string, libraryType MEDIATYPE) (*Library, error)
	GetLibraryByID(id ID) (*Library, error)
	UpdateLibraryLastScannedTime(id ID, time time.Time) error
	PruneMissingItemsFromLibrary(items map[string]struct{}, libID ID, libType MEDIATYPE) ([]string, error)
	MovieMaintainer
}

type MovieMaintainer interface {
	IsMovieFileAlreadyKnown(path string, parentID ID) bool
	InsertNewMovie(movie *Movie) (ID, error)
}
