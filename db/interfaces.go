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
	FindItemById(id string) (gomediacenter.MEDIATYPE, interface{}, error)
	FindItemUserData(uid, itemId string) (*gomediacenter.ItemUserData, error)
}

// IntroFinder can find intros for an item.
type IntroFinder interface {
	FindItemIntro(id string) (*[]gomediacenter.Intro, error)
}

// ItemUserData can find user data for an item.
type ItemUserData interface {
	InsertItemUserData(uid, itemId string) error
	DeleteItemUserData(uid, itemId string) error
}
