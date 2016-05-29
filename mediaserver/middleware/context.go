package middleware

import (
	"github.com/tcm1911/gomediacenter/mediaserver/controllers"
	"net/http"
)

func WithContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controllers.OpenContext(r)
		defer controllers.CloseContext(r)
		next(w, r)
	}
}
