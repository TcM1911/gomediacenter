package db

import "github.com/tcm1911/gomediacenter"

type Database interface {
	Connect(host string)
	Close()
	ItemFinder
	ItemUserData
}

type ItemFinder interface {
	FindItemById(id string) (gomediacenter.MEDIATYPE, interface{}, error)
	FindItemUserData(uid, itemId string) (*gomediacenter.ItemUserData, error)
}

type ItemUserData interface {
	InsertItemUserData(uid, itemId string) error
	DeleteItemUserData(uid, itemId string) error
}
