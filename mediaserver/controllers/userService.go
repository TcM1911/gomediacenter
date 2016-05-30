package controllers

import (
	"log"
	"net/http"

	"strconv"

	"net/url"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetAllUsers returns all the users. Route: "/Users". This action requires an
// authenticated user. The list can be filtered by: IsHidden, IsDisabled, and IsGuest.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving request for all users.")
	queryVars := GetContextVar(r, "queryVars").(url.Values)
	filter := make(map[string]interface{})
	keys := []string{"IsHidden", "IsDisabled", "IsGuest"}
	for _, key := range keys {
		if val := queryVars.Get(key); val != "" {
			if parsedVal, err := strconv.ParseBool(val); err == nil {
				filter[key] = parsedVal
				log.Println("Search filtered by", key, "=", parsedVal)
			}
		}
	}

	db := GetContextVar(r, "db").(db.UserManager)
	users, err := db.GetAllUsers(filter)
	if err != nil {
		http.Error(w, "error when getting all users", http.StatusInternalServerError)
		log.Println("Error when querying for all users:", err)
		return
	}
	log.Println("Number of users returned:", len(users))
	if len(users) == 0 {
		// Make sure an empty array is returned instead of nil.
		users = []*gomediacenter.User{}
	}
	writeJsonBody(w, users)
}

//[Route("/Users/Public", "GET", Summary = "Gets a list of publicly visible users for display on a login screen.")]

// GetUserById gets a user by Id. Route: /Users/{uid}.
// Can only be accessed by the authenticated user or admin.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	uid := pathVars["uid"]

	if !bson.IsObjectIdHex(uid) {
		http.Error(w, "not a valid id", http.StatusBadRequest)
		return
	}

	db := GetContextVar(r, "db").(db.UserManager)

	user, err := db.GetUserById(uid)
	if err == mgo.ErrNotFound {
		http.Error(w, "no user with that id", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "error serving the request", http.StatusInternalServerError)
		log.Println("Error while retrieving the user", uid, ":", err)
		return
	}
	writeJsonBody(w, user)
}

//[Route("/Users/{Id}/Offline", "GET", Summary = "Gets an offline user record by Id")]
//[Authenticated]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Users/{Id}", "DELETE", Summary = "Deletes a user")]
//[Authenticated(Roles = "Admin")]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Users/{Id}/Authenticate", "POST", Summary = "Authenticates a user")]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Password", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/AuthenticateByName", "POST", Summary = "Authenticates a user")]
//[ApiMember(Name = "Username", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
//[ApiMember(Name = "Password", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
//[ApiMember(Name = "PasswordMd5", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/{Id}/Password", "POST", Summary = "Updates a user's password")]
//[Authenticated]

//[Route("/Users/{Id}/EasyPassword", "POST", Summary = "Updates a user's easy password")]
//[Authenticated]

//[Route("/Users/{Id}", "POST", Summary = "Updates a user")]
//[Authenticated]

//[Route("/Users/{Id}/Policy", "POST", Summary = "Updates a user policy")]
//[Authenticated(Roles = "admin")]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Users/{Id}/Configuration", "POST", Summary = "Updates a user configuration")]
//[Authenticated]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

// NewUser creates a user. The name of the user to be created is passed as the
// parameter Name in the POST body.
func NewUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	if name == "" {
		http.Error(w, "No username given", http.StatusBadRequest)
		return
	}

	log.Println("Creating a new user named:", name)
	user := gomediacenter.NewUser(name)

	db := GetContextVar(r, "db").(db.UserManager)
	if err := db.AddNewUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error when saving the new user to the database:", err)
		return
	}

	// Return the user in the response body. This is how Emby does it.
	writeJsonBody(w, user)
}

//[Route("/Users/ForgotPassword", "POST", Summary = "Initiates the forgot password process for a local user")]
//[ApiMember(Name = "EnteredUsername", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/ForgotPassword/Pin", "POST", Summary = "Redeems a forgot password pin")]
//[ApiMember(Name = "Pin", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]
