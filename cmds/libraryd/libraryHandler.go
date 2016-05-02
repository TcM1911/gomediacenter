package main

import (
	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

// Holds the libraries watched by this instance.
var handler map[bson.ObjectId]db.Library
var handlerMutex sync.RWMutex

// Adds a library to the handler.
func addToLibraryHandler(d db.Library) {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	if handler == nil {
		handler = make(map[bson.ObjectId]db.Library)
	}
	handler[d.Id] = d
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
func getCopyOfLibrary(id bson.ObjectId) db.Library {
	handlerMutex.RLock()
	defer handlerMutex.RUnlock()
	lib := handler[id]
	return lib
}

func fetchLibraryDataFromDB(id bson.ObjectId) error {
	newLib, err := db.GetLibraryById(id)
	if err != nil {
		return err
	}
	addToLibraryHandler(*newLib)
	return nil
}
