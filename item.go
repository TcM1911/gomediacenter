package gomediacenter

import (
	"errors"
	"log"
	"sync"
)

// StaticItemProvider is the interface that should be implemented
// by static item providers.
type StaticItemProvider interface {
	// ProvideItem returns true if the provider can provide the item.
	ProvideItem(ID) bool
	// GetItem returns the item as a ReadSeekCloser.
	GetItem(ID) (ReadSeekCloser, error)
	// GetID returns the ID for the provider.
	GetID() ID
}

var staticProviderLock sync.RWMutex
var staticProviders map[string]StaticItemProvider
var setup sync.Once

var (
	// ErrProviderNotRegistered is returned when trying to remove a providers
	// that is not registered.
	ErrProviderNotRegistered = errors.New("provider not registered")
	// ErrProviderAlreadyRegistered is returned when trying to register a provider
	// that is already registered.
	ErrProviderAlreadyRegistered = errors.New("provider already registered")
	// ErrItemMissing is returned if the item has been removed from the provider.
	ErrItemMissing = errors.New("item is missing")
	// ErrNoProviderFound is returned if no provider can serve the item.
	ErrNoProviderFound = errors.New("no provider found")
)

// RegisterStaticProvider adds the static provider to known providers.
func RegisterStaticProvider(p StaticItemProvider) error {
	id := p.GetID()
	idStr := id.String()
	staticProviderLock.Lock()
	defer staticProviderLock.Unlock()
	setup.Do(func() {
		staticProviders = make(map[string]StaticItemProvider)
	})
	_, ok := staticProviders[idStr]
	if ok {
		return ErrProviderAlreadyRegistered
	}
	staticProviders[idStr] = p
	return nil
}

// RemoveStaticProvider removes the provider from registered providers.
func RemoveStaticProvider(p StaticItemProvider) error {
	id := p.GetID()
	idStr := id.String()
	staticProviderLock.Lock()
	defer staticProviderLock.Unlock()
	// Check if provider is registered.
	_, ok := staticProviders[idStr]
	if !ok {
		return ErrProviderNotRegistered
	}
	delete(staticProviders, idStr)
	return nil
}

// GetItem gets the item from one of the registered providers.
func GetItem(id ID) (ReadSeekCloser, error) {
	var errItemMissing error
	staticProviderLock.RLock()
	defer staticProviderLock.RUnlock()
	for _, p := range staticProviders {
		if p.ProvideItem(id) {
			r, err := p.GetItem(id)
			if err == ErrItemMissing {
				// Item was in cache but removed, go to next provider.
				log.Printf("%s appears to be missing, this can be due to a cached item being removed.", id.String())
				errItemMissing = err
				continue
			}
			if err != nil {
				return nil, err
			}
			pID := p.GetID()
			log.Printf("Serving item %s from provider %s", id.String(), pID.String())
			return r, nil
		}
	}
	if errItemMissing != nil {
		return nil, errItemMissing
	}
	log.Printf("No provider found for %s", id.String())
	return nil, ErrNoProviderFound
}
