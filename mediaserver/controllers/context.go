package controllers

import (
	"net/http"
	"sync"
)

var context map[*http.Request]map[string]interface{}
var contextLock sync.RWMutex

func OpenContext(r *http.Request) {
	contextLock.Lock()
	defer contextLock.Unlock()

	if context == nil {
		context = map[*http.Request]map[string]interface{}{}
	}
	context[r] = map[string]interface{}{}
}

func CloseContext(r *http.Request) {
	contextLock.Lock()
	defer contextLock.Unlock()
	delete(context, r)
}

func GetContextVar(r *http.Request, key string) interface{} {
	contextLock.RLock()
	defer contextLock.RUnlock()
	return context[r][key]
}

func SetContextVar(r *http.Request, key string, val interface{}) {
	contextLock.Lock()
	defer contextLock.Unlock()
	context[r][key] = val
}
