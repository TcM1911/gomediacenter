package handlers

import (
	"net/http"
	"sync"
)

var context map[*http.Request]map[string]interface{}
var contextLock sync.RWMutex

// OpenContext creates a new context object for a request.
func OpenContext(r *http.Request) {
	contextLock.Lock()
	defer contextLock.Unlock()

	if context == nil {
		context = map[*http.Request]map[string]interface{}{}
	}
	context[r] = map[string]interface{}{}
}

// CloseContext destroys the context object for a request.
func CloseContext(r *http.Request) {
	contextLock.Lock()
	defer contextLock.Unlock()
	delete(context, r)
}

// GetContextVar return a variable from the context object.
func GetContextVar(r *http.Request, key string) interface{} {
	contextLock.RLock()
	defer contextLock.RUnlock()
	return context[r][key]
}

// SetContextVar sets a variable.
func SetContextVar(r *http.Request, key string, val interface{}) {
	contextLock.Lock()
	defer contextLock.Unlock()
	context[r][key] = val
}
