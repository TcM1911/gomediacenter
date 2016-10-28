package main

import (
	"sync"

	"github.com/tcm1911/gomediacenter"
)

// Holds the libraries watched by this instance.
var handler map[string]gomediacenter.Library
var handlerMutex sync.RWMutex

// Adds a library to the handler.
func addToLibraryHandler(d gomediacenter.Library) {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	if handler == nil {
		handler = make(map[string]gomediacenter.Library)
	}
	handler[d.ID.String()] = d
}

// Remove a library from the handler.
func removeLibraryFromHandler(id gomediacenter.ID) {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	if _, ok := handler[id.String()]; ok {
		delete(handler, id.String())
	}
}

// Returns a copy of the library data.
func getCopyOfLibrary(id gomediacenter.ID) gomediacenter.Library {
	handlerMutex.RLock()
	defer handlerMutex.RUnlock()
	lib := handler[id.String()]
	return lib
}

func fetchLibraryDataFromDB(db gomediacenter.LibraryMaintainer, id gomediacenter.ID) error {
	newLib, err := db.GetLibraryByID(id)
	if err != nil {
		return err
	}
	addToLibraryHandler(*newLib)
	return nil
}
