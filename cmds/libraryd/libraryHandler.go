package main

import (
	"sync"

	"github.com/tcm1911/gomediacenter"

	"gopkg.in/mgo.v2/bson"
)

// Holds the libraries watched by this instance.
var handler map[bson.ObjectId]gomediacenter.Library
var handlerMutex sync.RWMutex

// Adds a library to the handler.
func addToLibraryHandler(d gomediacenter.Library) {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	if handler == nil {
		handler = make(map[bson.ObjectId]gomediacenter.Library)
	}
	handler[d.ID] = d
}

// Remove a library from the handler.
func removeLibraryFromHandler(id bson.ObjectId) {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	if _, ok := handler[id]; ok {
		delete(handler, id)
	}
}

// Returns a copy of the library data.
func getCopyOfLibrary(id bson.ObjectId) gomediacenter.Library {
	handlerMutex.RLock()
	defer handlerMutex.RUnlock()
	lib := handler[id]
	return lib
}

func fetchLibraryDataFromDB(db gomediacenter.LibraryMaintainer, id bson.ObjectId) error {
	newLib, err := db.GetLibraryByID(id)
	if err != nil {
		return err
	}
	addToLibraryHandler(*newLib)
	return nil
}
