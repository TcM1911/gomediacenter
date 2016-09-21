package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers"
	"gopkg.in/mgo.v2/bson"
)

func WithPathVars(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controllers.SetContextVar(r, "pathVars", mux.Vars(r))
		next(w, r)
	}
}

func WithQueryVars(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w,
				"error parsing query vars: "+err.Error(),
				http.StatusBadRequest)
		}
		controllers.SetContextVar(r, "queryVars", r.Form)
		next(w, r)
	}
}

func VerifyIds(ids []string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := controllers.GetContextVar(r, "pathVars").(map[string]string)

		for _, key := range ids {
			id := pathVars[key]
			if !bson.IsObjectIdHex(id) {
				http.Error(w, "not a valid id", http.StatusBadRequest)
				return
			}
		}

		next(w, r)
	}
}
