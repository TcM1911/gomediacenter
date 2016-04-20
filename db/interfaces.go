package db

import "github.com/tcm1911/gomediacenter"

type Database interface {
	Connect(host string)
	Close()
	ItemFinder
	ItemUserData
	IntroFinder
}

type ItemFinder interface {
	FindItemById(id string) (gomediacenter.MEDIATYPE, interface{}, error)
	FindItemUserData(uid, itemId string) (*gomediacenter.ItemUserData, error)
}

type IntroFinder interface {
	FindItemIntro(id string) (*[]gomediacenter.Intro, error)
}

type ItemUserData interface {
	InsertItemUserData(uid, itemId string) error
	DeleteItemUserData(uid, itemId string) error
}
