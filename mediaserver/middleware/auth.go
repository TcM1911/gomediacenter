package middleware

import (
	"net/http"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers/auth"
)

func AdminOrLoggedInUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := controllers.GetContextVar(r, "pathVars").(map[string]string)
		uid := pathVars["uid"]
		token := r.Header.Get(gomediacenter.SESSION_KEY_HEADER)
		session := auth.GetSession(token)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if session.UserId == uid || session.Admin {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func Admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(gomediacenter.SESSION_KEY_HEADER)
		session := auth.GetSession(token)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if session.Admin {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
