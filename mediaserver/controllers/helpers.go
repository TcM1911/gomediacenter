package controllers

import (
	"encoding/json"
	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func writeJsonBody(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func getFilteredUserList(filter map[string]interface{}, r *http.Request) ([]*gomediacenter.User, error) {
	db := GetContextVar(r, "db").(db.UserManager)
	users, err := db.GetAllUsers(filter)
	if err != nil {
		return users, err
	}
	log.Println("Number of users returned:", len(users))
	if len(users) == 0 {
		// Make sure an empty array is returned instead of nil.
		users = []*gomediacenter.User{}
	}
	return users, nil
}

func getIdFromPathVarAndCheckForErr(key string, pathVar map[string]string, w http.ResponseWriter) (string, bool) {
	id := pathVar[key]

	if !bson.IsObjectIdHex(id) {
		http.Error(w, "not a valid id", http.StatusBadRequest)
		return "", false
	}
	return id, true
}

func checkAndWriteHTTPErrorForIdQueries(id string, err error, logMsg string, w http.ResponseWriter) bool {
	if err == mgo.ErrNotFound {
		http.Error(w, "no user with that id", http.StatusBadRequest)
		return false
	} else if err != nil {
		http.Error(w, "error serving the request", http.StatusInternalServerError)
		log.Println(logMsg, id, ":", err)
		return false
	}
	return true
}
