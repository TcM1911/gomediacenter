package db

import "github.com/tcm1911/gomediacenter"

// Database is the interface for the database.
type Database interface {
	Connect(host string)
	Close()
	ItemFinder
	ItemUserData
	IntroFinder
}

// ItemFinder can find an item in the database.
type ItemFinder interface {
	FindItemByID(id string) (gomediacenter.MEDIATYPE, interface{}, error)
	FindItemUserData(uid, itemId string) (*gomediacenter.ItemUserData, error)
}

// IntroFinder can find intros for an item.
type IntroFinder interface {
	FindItemIntro(id string) ([]*gomediacenter.Intro, error)
}

// ItemUserData can find user data for an item.
type ItemUserData interface {
	InsertItemUserData(userData *gomediacenter.ItemUserData) error
}

// UserManager can create, update, and remove user profiles.
type UserManager interface {
	AddNewUser(user *gomediacenter.User) error
	UpdateUser(ID string, user *gomediacenter.User) error
	GetUserByID(hexString string) (*gomediacenter.User, error)
	GetUserByName(name string) (*gomediacenter.User, error)
	GetAllUsers(filter map[string]interface{}) ([]*gomediacenter.User, error)
	DeleteUser(hexString string) error
	ChangeUserPassword(uid string, newPassword []byte) error
}

// SessionManager can save and get stored sessions from the database.
type SessionManager interface {
	GetSavedSessions() ([]*gomediacenter.Session, error)
	SaveSessions([]*gomediacenter.Session) error
}
