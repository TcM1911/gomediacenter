package middleware

import (
	"net/http"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers/auth"
)

// AdminOrLoggedInUser only calls the next HandlerFunc if the session is and admin
// or a user session with the same UID as the path.
func AdminOrLoggedInUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := controllers.GetContextVar(r, "pathVars").(map[string]string)
		uid := pathVars["uid"]
		token := r.Header.Get(gomediacenter.SessionKeyHeader)
		session := auth.GetSession(token)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if session.UserID == uid || session.Admin {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

// Admin only calls the next function if the session is an admin session.
func Admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(gomediacenter.SessionKeyHeader)
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
