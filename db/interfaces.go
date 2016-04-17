package db

import "github.com/tcm1911/gomediacenter"

type Database interface {
	ItemFinder
}

type ItemFinder interface {
	FindItemById(id string) (gomediacenter.MEDIATYPE, interface{}, error)
}
