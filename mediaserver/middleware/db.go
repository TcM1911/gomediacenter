package middleware

import (
	"github.com/tcm1911/gomediacenter/db"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers"
	"net/http"
)

func WithDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		thisDB := db.GetDB()
		controllers.SetContextVar(r, "db", thisDB)
		next(w, r)
	}
}
