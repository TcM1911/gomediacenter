package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/tcm1911/gomediacenter"
)

func WithPathVars(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		for k, v := range vars {
			id, err := gomediacenter.IDFromString(v)
			if err != nil {
				msgerr := fmt.Sprintf("%v is not a valid id.", v)
				log.Printf("Aborting request. Could not parse %v to an id: %v", v, err.Error())
				http.Error(w, msgerr, http.StatusBadRequest)
				return
			}
			switch k {
			case "id":
				r = r.WithContext(gomediacenter.AddIDToContext(r.Context(), id))
			case "uid":
				r = r.WithContext(gomediacenter.AddUIDToContext(r.Context(), id))
			}
		}
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
		r = r.WithContext(AddQueryVarsToContext(r.Context(), r.Form))
		next(w, r)
	}
}

func VerifyIds(ids []string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*r = r.WithContext(context.WithValue(r.Context(), "pathVars", mux.Vars(r)))

		for _, key := range ids {
			id := pathVars[key]
			if !bson.IsObjectIdHex(id) {
				http.Error(w, "not a valid id", http.StatusBadRequest)
				return
			}
		}*/

		next(w, r)
	}
}

// AddQueryVarsToContext adds the query vars to the context.
func AddQueryVarsToContext(ctx context.Context, vars url.Values) context.Context {
	return context.WithValue(ctx, PathVarsCtxKey, vars)
}

// GetQueryVarsFromContext gets the query vars from the context.
func GetQueryVarsFromContext(ctx context.Context) url.Values {
	pv, ok := ctx.Value(PathVarsCtxKey).(url.Values)
	if !ok {
		panic("Failed to type assert the path vars.")
	}
	return pv
}

const PathVarsCtxKey = "PathVarsCtxKey"
