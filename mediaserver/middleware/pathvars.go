package middleware

import (
	"github.com/gorilla/mux"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers"
	"net/http"
)

func WithPathVars(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controllers.SetContextVar(r, "pathVars", mux.Vars(r))
		next(w, r)
	}
}
