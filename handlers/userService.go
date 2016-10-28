package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/middleware"
	"golang.org/x/crypto/bcrypt"
)

/////////////
// Structs //
/////////////

type client struct {
	Client, Device, DeviceID, Version string
}

////////////
// Public //
////////////

// GetAllUsers returns all the users. Route: "/Users". This action requires an
// authenticated user. The list can be filtered by: IsHidden, IsDisabled, and IsGuest.
func GetAllUsers(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving request for all users.")
		queryVars := middleware.GetQueryVarsFromContext(r.Context())
		filter := make(map[string]interface{})
		keys := []string{"IsHidden", "IsDisabled", "IsGuest"}
		for _, key := range keys {
			if val := queryVars.Get(key); val != "" {
				if parsedVal, err := strconv.ParseBool(val); err == nil {
					filter[key] = parsedVal
					log.Printf("Search filtered by %s = %t\n", key, parsedVal)
				}
			}
		}
		users, err := getFilteredUserList(filter, db)
		if err != nil {
			http.Error(w, "error when getting all users", http.StatusInternalServerError)
			log.Printf("Error when querying for all users: %s", err)
			return
		}
		writeJSONBody(w, users)
	}
}

// GetAllUsersPublic gets a list of publicly visible users for display on a login screen.
func GetAllUsersPublic(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving request for public viewing of all users.")
		filter := make(map[string]interface{})
		filter["IsHidden"] = false
		users, err := getFilteredUserList(filter, db)
		if err != nil {
			http.Error(w, "error when getting all users", http.StatusInternalServerError)
			log.Printf("Error when querying for all users: %s", err)
			return
		}

		var publicList []*gomediacenter.PublicUserResponse
		for _, user := range users {
			publicList = append(publicList,
				&gomediacenter.PublicUserResponse{Name: user.Name, ID: user.ID})
		}
		writeJSONBody(w, publicList)
	}
}

// GetUserByID gets a user by Id. Route: /Users/{uid}.
// Can only be accessed by the authenticated user or admin.
func GetUserByID(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := gomediacenter.GetUIDFromContext(r.Context())

		user, err := db.GetUserByID(uid)
		if ok := checkAndWriteHTTPErrorForIDQueries(uid.String(), err,
			"Error while retrieving the user", w); !ok {
			return
		}
		writeJSONBody(w, user)
	}
}

// GetOfflineUserByID gets an offline user record by Id.
func GetOfflineUserByID(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetUserByID(db).ServeHTTP(w, r)
	}
}

// DeleteUser deletes a user and all it's item data when a DELETE request is sent to to /Users/{uid}.
// This action requires admin rights.
func DeleteUser(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving removing user request")
		uid := gomediacenter.GetUIDFromContext(r.Context())
		err := db.DeleteUser(uid)
		if ok := checkAndWriteHTTPErrorForIDQueries(uid.String(), err,
			"Error when deleting", w); !ok {
			return
		}
		log.Printf("User %s removed.", uid.String())
		w.WriteHeader(http.StatusOK)
	}
}

// Authenticate authenticates a user when a POST is sent to /Users/{uid}/Authenticate.
// The password is past in the body in the parameter password.
func Authenticate(db gomediacenter.UserManager, auth gomediacenter.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Processing authentication request")
		uid := gomediacenter.GetUIDFromContext(r.Context())
		user, err := db.GetUserByID(uid)
		if ok := checkAndWriteHTTPErrorForIDQueries(uid.String(), err, "Error when getting user data", w); !ok {
			return
		}
		var form gomediacenter.LoginRequest
		defer gomediacenter.CloseReqBody(r)
		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			logError(w, err, "Error when processing auth request body:",
				"Error when processing the request", http.StatusBadRequest)
			return
		}
		passwd := form.Password
		authenticateUser(auth, user, passwd, w, r)
	}
}

// AuthenticateByName authenticates a user when a POST is sent to /Users/{uid}/AuthenticateByName.
// Username and password is past in the body as the parameters Username and password.
func AuthenticateByName(db gomediacenter.UserManager, auth gomediacenter.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Processing authentication request by name")

		var form gomediacenter.LoginRequest
		defer gomediacenter.CloseReqBody(r)
		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			logError(w, err, "Error when processing auth request body:",
				"Error when processing the request", http.StatusBadRequest)
			return
		}
		username, passwd := form.Name, form.Password
		if username == "" {
			logError(w, errors.New("empty username"), "Username was missing in the request.", "username can't be empty", http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByName(username)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Printf("Error when trying to retrieve user profile for %s: %s", username, err)
			return
		}
		authenticateUser(auth, user, passwd, w, r)
	}
}

// LogoutUser logs out the user and removes the sessions from the session manager.
func LogoutUser(auth gomediacenter.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := gomediacenter.GetUIDFromContext(r.Context())
		key := gomediacenter.GetIDFromContext(r.Context())
		if ok := auth.RemoveSession(uid, key); ok {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}

// ChangeUserPassword handles a password change for a user.
// New password and current password are past as body parameters
// newPassword and currentPassword.
func ChangeUserPassword(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := gomediacenter.GetUIDFromContext(r.Context())
		log.Printf("Changing the password for %s", uid.String())
		var req gomediacenter.PasswordRequest
		defer gomediacenter.CloseReqBody(r)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logError(w, err, "Faild to decode the password request body:", "faild to decode request body",
				http.StatusBadRequest)
			return
		}
		newPass, currentPass := req.New, req.Current
		if newPass == "" {
			logError(w, errors.New("no password given"), "Failed to handle password change request:",
				"new password is required", http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByID(uid)
		if ok := checkAndWriteHTTPErrorForIDQueries(uid.String(), err, "Error when getting user data", w); !ok {
			return
		}
		if user.HasPasswd && (bcrypt.CompareHashAndPassword(
			user.Password,
			[]byte(currentPass)) != nil) {
			logmsg := fmt.Sprintf("Verification of the current password failed for %s", uid.String())
			logError(w, errors.New("password doesn't match"), logmsg,
				"password and uid doesn't match", http.StatusUnauthorized)
			return
		}
		passHash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
		if err != nil {
			logError(w, err, "Error when generating password hash:", "failed to generate a new hash",
				http.StatusInternalServerError)
			return
		}
		if err := db.ChangeUserPassword(uid, passHash); err != nil {
			logError(w, err, "Error when updating password for"+uid.String()+":", "Error when updating the password",
				http.StatusInternalServerError)
			return
		}
		log.Printf("Password for %s has been updated.\n", uid.String())
		w.WriteHeader(http.StatusOK)
	}
}

//[Route("/Users/{Id}/EasyPassword", "POST", Summary = "Updates a user's easy password")]
//[Authenticated]

// UpdateUser handles an user update request.
func UpdateUser(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handle user update request.")
		var newUserStruct *gomediacenter.User
		defer gomediacenter.CloseReqBody(r)
		if err := json.NewDecoder(r.Body).Decode(&newUserStruct); err != nil {
			logError(w, err,
				"Error when decoding a user update request:",
				"Failed to decode user update request", http.StatusBadRequest)
			return
		}

		if newUserStruct == nil {
			logError(w, nil, "The new user struct is nil, aborting",
				"Bad request.", http.StatusBadRequest)
			return
		}

		uid := gomediacenter.GetUIDFromContext(r.Context())
		log.Printf("User %s is being updated.\n", uid.String())
		if err := db.UpdateUser(uid, newUserStruct); err != nil {
			logError(w, err, "Error when processing the request:",
				"Error processing the request", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// UpdateUserPolicy updates a user policy. This action can only be performed by an admin.
// The new policy is in the request body.
func UpdateUserPolicy(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p *gomediacenter.UserPolicy
		uid, err := getUIDAndRequestBody(r, w, &p)
		if err != nil {
			return
		}

		if err = db.UpdateUserPolicy(uid, p); err != nil {
			logError(w, err, "Error when updating user policy.", "Error when updating the user's policy.", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// UpdateUserConfiguration updates a user's configuration. This action can be performed by the user and admin.
// The new configuration is in the request body.
func UpdateUserConfiguration(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c *gomediacenter.UserConfig
		uid, err := getUIDAndRequestBody(r, w, &c)
		if err != nil {
			return
		}
		log.Printf("Updating user %s's configuration.", uid.String())
		if err = db.UpdateUserConfiguration(uid, c); err != nil {
			logError(w, err, "Error when updating user's configuration.",
				"Error when updating user's configuration.", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}

// NewUser creates a user. The name of the user to be created is passed as the
// parameter Name in the POST body.
func NewUser(db gomediacenter.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userReq gomediacenter.NewUserRequest
		defer gomediacenter.CloseReqBody(r)
		if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
			logError(w, err, "Error when decoding the request:", "Failed to process the request",
				http.StatusBadRequest)
			return
		}
		name := userReq.Name

		if name == "" {
			logError(w, errors.New("empty username"), "Failed process adding new user request:",
				"No username given", http.StatusBadRequest)
			return
		}

		log.Println("Creating a new user named:", name)
		user := gomediacenter.NewUser(name)

		if err := db.AddNewUser(user); err != nil {
			logError(w, err, "Error when saving the new user to the database:",
				"Failing to processing the request.", http.StatusInternalServerError)
			return
		}

		// Return the user in the response body. This is how Emby does it.
		writeJSONBody(w, user)
	}
}

//[Route("/Users/ForgotPassword", "POST", Summary = "Initiates the forgot password process for a local user")]
//[ApiMember(Name = "EnteredUsername", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/ForgotPassword/Pin", "POST", Summary = "Redeems a forgot password pin")]
//[ApiMember(Name = "Pin", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]

/////////////
// Private //
/////////////

func getUIDFromPathVars(r *http.Request) (string, error) {
	vars, ok := GetContextVar(r, "pathVars").(map[string]string)
	if !ok {
		return "", errors.New("failed to type assert pathVars to a map[string]string")
	}
	return vars["uid"], nil
}

func authenticateUser(auth gomediacenter.SessionManager, user *gomediacenter.User, passwd string, w http.ResponseWriter, r *http.Request) {
	if user.HasPasswd && (bcrypt.CompareHashAndPassword(user.Password, []byte(passwd)) != nil) {
		logError(w, errors.New("incorrect password"), "Authentication failed:",
			"Username and password missmatch.", http.StatusUnauthorized)
		return
	}
	client, err := parseAuthHeader(r)
	if err != nil {
		logError(w, err, "Bad auth header:", "", http.StatusBadRequest)
		return
	}

	authToken := gomediacenter.NewRandomID()
	session := &gomediacenter.Session{
		ID:            authToken,
		UserID:        user.ID,
		UserName:      user.Name,
		Admin:         user.Policy.Admin,
		DeviceID:      client.DeviceID,
		DeviceName:    client.Device,
		Client:        client.Client,
		ClientVersion: client.Version,
	}
	auth.AddSession(session)
	log.Println(user.Name, "authenticated.")
	resp := &gomediacenter.AuthUserResponse{}
	resp.Token = authToken.String()
	resp.Session = session
	resp.User = user
	writeJSONBody(w, resp)
}

func parseAuthHeader(r *http.Request) (client, error) {
	var client client
	header := r.Header.Get("x-emby-authorization")
	if header == "" {
		return client, errors.New("missing x-emby-authorization header")
	}

	// TODO: Rewrite!
	a := strings.Split(header, ",")
	for _, v := range a {
		keyVal := strings.Split(v, "=")
		k := strings.TrimSpace(keyVal[0])
		val := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(keyVal[1]), `"`), `"`)

		switch k {
		case "Device":
			client.Device = val
		case "DeviceId":
			client.DeviceID = val
		case "Version":
			client.Version = val
		case "MediaBrowser Client":
			client.Client = val
		}
	}

	if client.Client == "" && client.Version == "" && client.DeviceID == "" && client.Device == "" {
		return client, errors.New("missing information in the header")
	}

	return client, nil
}

func getUIDAndRequestBody(r *http.Request, w http.ResponseWriter, v interface{}) (gomediacenter.ID, error) {
	defer gomediacenter.CloseReqBody(r)
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		logError(w, err, "Error when decoding request body.", "Error when processing the request.", http.StatusBadRequest)
		return gomediacenter.ID{}, err
	}

	uid := gomediacenter.GetUIDFromContext(r.Context())
	return uid, nil
}
