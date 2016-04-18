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
}

type ItemUserData interface {
	FindItemUserData(uid, itemId string) (gomediacenter.UserItemData, error)
	InsertItemUserData(uid, itemId string) error
	DeleteItemUserData(uid, itemId string) error
}