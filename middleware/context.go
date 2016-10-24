package middleware

import "net/http"

func WithContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//controllers.OpenContext(r)
		//defer controllers.CloseContext(r)
		next(w, r)
	}
}
