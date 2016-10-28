package middleware

import (
	"log"
	"net/http"

	"github.com/tcm1911/gomediacenter"
)

// AdminOrLoggedInUser only calls the next HandlerFunc if the session is and admin
// or a user session with the same UID as the path.
func AdminOrLoggedInUser(auth gomediacenter.SessionManager, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := gomediacenter.GetUIDFromContext(r.Context())
		token, err := gomediacenter.IDFromString(r.Header.Get(gomediacenter.SessionKeyHeader))
		if err != nil {
			log.Printf("Not a valid session key: %s", err.Error())
			http.Error(w, "Not a valid session key", http.StatusBadRequest)
			return
		}
		session := auth.GetSession(token)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !session.UserID.Equal(uid) && !session.Admin {
			w.WriteHeader(http.StatusUnauthorized)
		}
		r = r.WithContext(gomediacenter.AddIDToContext(r.Context(), session.ID))
		next(w, r)
	}
}

// Admin only calls the next function if the session is an admin session.
func Admin(auth gomediacenter.SessionManager, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := gomediacenter.IDFromString(r.Header.Get(gomediacenter.SessionKeyHeader))
		if err != nil {
			log.Printf("Not a valid session key: %s", err.Error())
			http.Error(w, "Not a valid session key", http.StatusBadRequest)
			return
		}
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
