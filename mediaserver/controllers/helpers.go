package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func writeJSONBody(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error when writing the response to the message body", err)
	}
}

func getFilteredUserList(filter map[string]interface{}, db gomediacenter.UserManager) ([]*gomediacenter.User, error) {
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

func getIDFromPathVarAndCheckForErr(key string, r *http.Request, w http.ResponseWriter) (string, bool) {
	pathVars, ok := GetContextVar(r, "pathVars").(map[string]string)
	if !ok {
		log.Println("Failed to type set the pathVars")
		return "", false
	}
	id, ok := pathVars[key]
	if !ok || !bson.IsObjectIdHex(id) {
		logError(w, nil, "In valid id", "not a valid id", http.StatusBadRequest)
		return "", false
	}
	return id, true
}

func checkAndWriteHTTPErrorForIDQueries(id string, err error, logMsg string, w http.ResponseWriter) bool {
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

// logError writes logMsg to the log and returns the errMsg and the code to
// the api caller.
func logError(w http.ResponseWriter, e error, logMsg, errMsg string, code int) {
	http.Error(w, errMsg, code)
	log.Println(logMsg, e)
}
