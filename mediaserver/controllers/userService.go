package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers/auth"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

/////////////
// Structs //
/////////////

type client struct {
	Client, Device, DeviceId, Version string
}

////////////
// Public //
////////////

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
	users, err := getFilteredUserList(filter, r)
	if err != nil {
		http.Error(w, "error when getting all users", http.StatusInternalServerError)
		log.Println("Error when querying for all users:", err)
		return
	}
	writeJsonBody(w, users)
}

// GetAllUsersPublic gets a list of publicly visible users for display on a login screen.
func GetAllUsersPublic(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving request for public viewing of all users.")
	filter := make(map[string]interface{})
	filter["IsHidden"] = false
	users, err := getFilteredUserList(filter, r)
	if err != nil {
		http.Error(w, "error when getting all users", http.StatusInternalServerError)
		log.Println("Error when querying for all users:", err)
		return
	}

	var publicList []*gomediacenter.PublicUserResponse
	for _, user := range users {
		publicList = append(publicList,
			&gomediacenter.PublicUserResponse{Name: user.Name, Id: user.Id})
	}
	writeJsonBody(w, publicList)
}

// GetUserById gets a user by Id. Route: /Users/{uid}.
// Can only be accessed by the authenticated user or admin.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	uid, ok := getIdFromPathVarAndCheckForErr("uid", pathVars, w)
	if !ok {
		return
	}

	db := GetContextVar(r, "db").(db.UserManager)

	user, err := db.GetUserById(uid)
	if ok := checkAndWriteHTTPErrorForIdQueries(uid, err,
		"Error while retrieving the user", w); !ok {
		return
	}
	writeJsonBody(w, user)
}

// GetOfflineUserById gets an offline user record by Id.
func GetOfflineUserById(w http.ResponseWriter, r *http.Request) {
	GetUserById(w, r)
}

// A DELETE to /Users/{uid} deletes a user and all it's item data.
// This action requires admin rights.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving removing user request")
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	uid, ok := getIdFromPathVarAndCheckForErr("uid", pathVars, w)
	if !ok {
		return
	}

	db := GetContextVar(r, "db").(db.UserManager)
	err := db.DeleteUser(uid)
	if ok := checkAndWriteHTTPErrorForIdQueries(uid, err,
		"Error when deleting", w); !ok {
		return
	}
	log.Println("User", uid, "removed.")
	w.WriteHeader(http.StatusOK)
}

// A POST to /Users/{uid}/Authenticate authenticates a user.
// The password is past in the body in the parameter password.
func Authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing authentication request")
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	uid := pathVars["uid"]

	db := GetContextVar(r, "db").(db.UserManager)
	user, err := db.GetUserById(uid)
	if ok := checkAndWriteHTTPErrorForIdQueries(uid, err, "Error when getting user data", w); !ok {
		return
	}
	var form gomediacenter.LoginRequest
	json.NewDecoder(r.Body).Decode(&form)
	r.Body.Close()
	passwd := form.Password
	authenticateUser(user, passwd, w, r)
}

// A POST to /Users/AuthenticateByName authenticates a user.
// Username and password is past in the body as the parameters Username and password.
func AuthenticateByName(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing authentication request by name")

	var form gomediacenter.LoginRequest
	json.NewDecoder(r.Body).Decode(&form)
	r.Body.Close()
	username, passwd := form.Name, form.Password

	if username == "" {
		log.Println("Username was missing in the request.")
		http.Error(w, "username can't be empty", http.StatusBadRequest)
		return
	}

	db := GetContextVar(r, "db").(db.UserManager)
	user, err := db.GetUserByName(username)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println("Error when trying to retrieve user profile for", username, err)
		return
	}
	authenticateUser(user, passwd, w, r)
}

// LogoutUser logs out the user and removes the sessions from the session manager.
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	uid := pathVars["uid"]
	key := r.Header.Get(gomediacenter.SESSION_KEY_HEADER)
	if ok := auth.RemoveSession(uid, key); ok {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// A POST to /Users/{Id}/Password updates a user's password.
// New password and current password are past as body parameters
// newPassword and currentPassword.
func ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	uid := pathVars["uid"]
	log.Println("Changing the password for", uid)

	var req gomediacenter.PasswordRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to decode the request", http.StatusBadRequest)
		return
	}
	newPass, currentPass := req.New, req.Current
	if newPass == "" {
		http.Error(w, "new password is required", http.StatusBadRequest)
		return
	}

	db := GetContextVar(r, "db").(db.UserManager)
	user, err := db.GetUserById(uid)
	if ok := checkAndWriteHTTPErrorForIdQueries(uid, err, "Error when getting user data", w); !ok {
		return
	}

	if user.HasPasswd && (bcrypt.CompareHashAndPassword(
		user.Password,
		[]byte(currentPass)) != nil) {
		http.Error(w, "incorrect current password", http.StatusBadRequest)
		log.Printf("Verification of the current password for %s failed.\n", uid)
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "failed to generate a new hash", http.StatusInternalServerError)
		log.Println("Error when generating password hash:", err)
		return
	}

	if err := db.ChangeUserPassword(uid, passHash); err != nil {
		http.Error(w, "Error when updating the password", http.StatusInternalServerError)
		log.Printf("Error when updating password for %s, %s\n", uid, err)
		return
	}
	log.Printf("Password for %s has been updated.\n", uid)
	w.WriteHeader(http.StatusOK)
}

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
	var userReq gomediacenter.NewUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Failed to decode the request", http.StatusBadRequest)
		log.Println("Error when decoding the request", err)
		return
	}
	defer r.Body.Close()
	name := userReq.Name

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

/////////////
// Private //
/////////////

func authenticateUser(user *gomediacenter.User, passwd string, w http.ResponseWriter, r *http.Request) {
	if user.HasPasswd && (bcrypt.CompareHashAndPassword(user.Password, []byte(passwd)) != nil) {
		http.Error(w, "", http.StatusUnauthorized)
		log.Println("Invalid login by", user.Name)
		return
	}
	client, err := parseAuthHeader(r)
	if err != nil {
		log.Println("Bad auth header.")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authToken := bson.NewObjectId()
	session := &gomediacenter.Session{
		Id:            authToken,
		UserId:        user.Id.Hex(),
		UserName:      user.Name,
		Admin:         user.Policy.Admin,
		DeviceId:      client.DeviceId,
		DeviceName:    client.Device,
		Client:        client.Client,
		ClientVersion: client.Version,
	}
	go auth.AddSession(session)
	log.Println(user.Name, "authenticated.")
	resp := &gomediacenter.AuthUserResponse{}
	resp.Token = authToken.Hex()
	resp.Session = session
	resp.User = user
	writeJsonBody(w, resp)
}

func parseAuthHeader(r *http.Request) (client, error) {
	var client client
	header := r.Header.Get("x-emby-authorization")
	if header == "" {
		return client, errors.New("missing x-emby-authorization header")
	}

	a := strings.Split(header, ",")
	for _, v := range a {
		keyVal := strings.Split(v, "=")
		k := strings.TrimSpace(keyVal[0])
		val := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(keyVal[1]), `"`), `"`)

		switch k {
		case "Device":
			client.Device = val
		case "DeviceId":
			client.DeviceId = val
		case "Version":
			client.Version = val
		case "MediaBrowser Client":
			client.Client = val
		}
	}

	if client.Client == "" && client.Version == "" && client.DeviceId == "" && client.Device == "" {
		return client, errors.New("missing information in the header")
	}

	return client, nil
}