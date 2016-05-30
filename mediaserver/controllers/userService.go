package controllers

import (
	"log"
	"net/http"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//[Route("/Users", "GET", Summary = "Gets a list of users")]
//[Authenticated]
//[ApiMember(Name = "IsHidden", Description = "Optional filter by IsHidden=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "IsDisabled", Description = "Optional filter by IsDisabled=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "IsGuest", Description = "Optional filter by IsGuest=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]

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
