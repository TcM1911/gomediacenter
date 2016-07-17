package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tcm1911/gomediacenter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func TestNewUser(t *testing.T) {
	assert := assert.New(t)

	db := new(mockDB)
	db.On("AddNewUser", mock.Anything).Return(nil)

	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(gomediacenter.NewUserRequest{Name: "testUser"})

	req, _ := http.NewRequest("POST", "/Users/New", b)

	// Add to context
	OpenContext(req)
	defer CloseContext(req)
	SetContextVar(req, "db", db)

	writer := httptest.NewRecorder()
	NewUser(writer, req)

	body, err := getBodyStringFromRecorder(writer)
	if err != nil {
		assert.Fail("Could not read the body.")
	}

	assert.Contains(body, `"Name":"testUser"`)

	assert.Contains(body, `"Configuration":{`)
	assert.Contains(body, `"HasPassword":false`)
	assert.Contains(body, `"HasConfiguredPassword":false`)
	assert.Contains(body, `"HasConfiguredEasyPassword":false`)
	assert.Contains(body, `"EnableLocalPassword":false`)
	assert.Contains(body, `"IncludeTrailersInSuggestions":true,`)
	assert.Contains(body, `"RememberAudioSelections":true`)
	assert.Contains(body, `"RememberSubtitleSelections":true`)
	assert.Contains(body, `"EnableNextEpisodeAutoPlay":true`)
	assert.Contains(body, `"HidePlayedInLatest":true`)
	assert.Contains(body, `"EnableChannelView":false`)
	assert.Contains(body, `"PlayDefaultAudioTrack":true`)
	assert.Contains(body, `"DisplayMissingEpisodes":false`)
	assert.Contains(body, `"DisplayUnairedEpisodes":false`)
	assert.Contains(body, `"GroupMoviesIntoBoxSets":false`)
	assert.Contains(body, `"SubtitleMode":"Default"`)
	assert.Contains(body, `"DisplayCollectionsView":false,`)
	assert.Contains(body, `"DisplayFoldersView":false`)

	assert.Contains(body, `"Policy":{`)
	assert.Contains(body, `"IsAdministrator":false`)
	assert.Contains(body, `"IsHidden":false`)
	assert.Contains(body, `"IsDisabled":false`)
	assert.Contains(body, `"EnableUserPreferenceAccess":true`)
	assert.Contains(body, `"EnableRemoteControlOfOtherUsers":false`)
	assert.Contains(body, `"EnableSharedDeviceControl":true`)
	assert.Contains(body, `"EnableMediaPlayback":true`)
	assert.Contains(body, `"EnableAudioPlaybackTranscoding":true`)
	assert.Contains(body, `"EnableVideoPlaybackTranscoding":true`)
	assert.Contains(body, `"EnableContentDeletion":false`)
	assert.Contains(body, `"EnableContentDownloading":true`)
	assert.Contains(body, `"EnableSync":true`)
	assert.Contains(body, `"EnableSyncTranscoding":true`)
	assert.Contains(body, `"InvalidLoginAttemptCount":0`)
	assert.Contains(body, `"EnablePublicSharing":true`)
}

func TestGetUserByID(t *testing.T) {
	assert := assert.New(t)

	// Create test user.
	user := gomediacenter.NewUser("testUser")
	uid := bson.NewObjectId()
	user.ID = uid

	// Request
	r, _ := http.NewRequest("GET", "/", nil)
	OpenContext(r)
	defer CloseContext(r)

	// Setup db mock.
	db := new(mockDB)
	db.On("GetUserByID", uid.Hex()).Return(user, nil)
	SetContextVar(r, "db", db)

	// Setup pathVars
	pathVars := make(map[string]string)
	pathVars["uid"] = uid.Hex()
	SetContextVar(r, "pathVars", pathVars)

	recorder := httptest.NewRecorder()

	// Call http handler.
	GetUserByID(recorder, r)

	body, err := getBodyStringFromRecorder(recorder)
	if err != nil {
		assert.Fail("Could not read the body.")
	}

	assert.Contains(body, `"Name":"testUser"`)
	assert.Contains(body, `"id":"`+uid.Hex()+`"`)
}

func TestGetAllUsers(t *testing.T) {
	assert := assert.New(t)

	user1 := gomediacenter.NewUser("user1")
	user2 := gomediacenter.NewUser("user2")
	users := []*gomediacenter.User{user1, user2}

	r, _ := http.NewRequest("GET", "/", nil)
	OpenContext(r)
	defer CloseContext(r)

	db := new(mockDB)
	db.On("GetAllUsers", mock.Anything).Return(users, nil)
	SetContextVar(r, "db", db)

	queryVars := url.Values{}
	SetContextVar(r, "queryVars", queryVars)

	recorder := httptest.NewRecorder()
	GetAllUsers(recorder, r)

	body, err := getBodyStringFromRecorder(recorder)
	if err != nil {
		assert.Fail("Error when parsing the body", err)
	}

	assert.Contains(body, `[{"Name":"user1"`)
	assert.Contains(body, `},{"Name":"user2"`)
}

func TestParseAuthHeader(t *testing.T) {
	expected := client{
		Version:  "3.0.5912.0",
		Device:   "Chrome 50.0.2661.50",
		Client:   "Emby Web Client",
		DeviceID: "cae2cc5be4e17f1d0a486d0c8fdb4789f4f1e99c",
	}
	r, err := http.NewRequest("POST", "", nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	r.Header.Add("x-emby-authorization", `MediaBrowser Client="Emby Web Client", Device="Chrome 50.0.2661.50", DeviceId="cae2cc5be4e17f1d0a486d0c8fdb4789f4f1e99c", Version="3.0.5912.0", UserId="f40b2df070cf46e686bcbdd388d8706c"`)

	actual, err := parseAuthHeader(r)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, expected, actual)
}

func TestSettingAnUserPassword(t *testing.T) {
	assert := assert.New(t)

	user := &gomediacenter.User{ID: bson.NewObjectId(), Name: "User", HasPasswd: false}
	r, err := passwordChangeContext(user, "", "password")
	defer CloseContext(r)
	if err != nil {
		assert.Fail(err.Error())
	}

	recorder := httptest.NewRecorder()

	ChangeUserPassword(recorder, r)

	assert.Equal(http.StatusOK, recorder.Code)
}

func TestChangeUserPassword(t *testing.T) {
	assert := assert.New(t)

	currentPass := "currentPass"
	currentHash, err := bcrypt.GenerateFromPassword([]byte(currentPass), bcrypt.DefaultCost)
	if err != nil {
		assert.Fail(err.Error())
	}

	user := &gomediacenter.User{
		ID:        bson.NewObjectId(),
		Name:      "User",
		HasPasswd: true,
		Password:  currentHash,
	}
	r, err := passwordChangeContext(user, currentPass, "newPassword")
	defer CloseContext(r)
	if err != nil {
		assert.Fail(err.Error())
	}

	recorder := httptest.NewRecorder()

	ChangeUserPassword(recorder, r)

	assert.Equal(http.StatusOK, recorder.Code)
}

// Password should not be changed if the password in the POST doesn't match what's in the db.
func TestChangeUserPasswordWhenPasswordDoesNotMatch(t *testing.T) {
	assert := assert.New(t)

	currentHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		assert.Fail(err.Error())
	}

	user := &gomediacenter.User{
		ID:        bson.NewObjectId(),
		Name:      "User",
		HasPasswd: true,
		Password:  currentHash,
	}
	r, err := passwordChangeContext(user, "wrongPassword", "newPassword")
	defer CloseContext(r)
	if err != nil {
		assert.Fail(err.Error())
	}

	recorder := httptest.NewRecorder()

	ChangeUserPassword(recorder, r)

	assert.Equal(http.StatusBadRequest, recorder.Code)
}

func passwordChangeContext(user *gomediacenter.User, currentPass, newPass string) (*http.Request, error) {
	// Mock setup.
	db := new(mockDB)
	db.On("ChangeUserPassword", user.ID.Hex(), mock.Anything).Return(nil)
	db.On("GetUserByID", user.ID.Hex()).Return(user, nil)

	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(gomediacenter.PasswordRequest{
		New: newPass, Current: currentPass})

	// Context setup.
	r, err := http.NewRequest("POST", "", b)
	if err != nil {
		return nil, err
	}
	OpenContext(r)

	SetContextVar(r, "db", db)
	pathVars := make(map[string]string)
	pathVars["uid"] = user.ID.Hex()
	SetContextVar(r, "pathVars", pathVars)

	return r, nil
}

func TestAuthenticateWithCorrectPassword(t *testing.T) {
	r, err := authenticateContextSetup("testpassword", "testpassword", "", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	Authenticate(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthenticateWithIncorrectPassword(t *testing.T) {
	r, err := authenticateContextSetup("testpassword", "incorrect", "", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	Authenticate(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateWithoutCorrectHeader(t *testing.T) {
	r, err := authenticateContextSetup("testpassword", "testpassword", "", false)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	Authenticate(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthenticateByNameWithCorrectPassword(t *testing.T) {
	r, err := authenticateContextSetup("testpassword", "testpassword", "Username", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	AuthenticateByName(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthenticateByNameWithIncorrectPassword(t *testing.T) {
	r, err := authenticateContextSetup("testpassword", "incorrect", "Username", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	AuthenticateByName(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateByNameWithoutCorrectHeader(t *testing.T) {
	r, err := authenticateContextSetup("testpassword", "testpassword", "Username", false)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	AuthenticateByName(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func authenticateContextSetup(userPass, loginPass, bodyUserName string, withHeader bool) (*http.Request, error) {

	user := &gomediacenter.User{ID: bson.NewObjectId()}

	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(gomediacenter.LoginRequest{
		Name: bodyUserName, Password: loginPass})

	// Context setup.
	r, err := http.NewRequest("POST", "", b)
	if err != nil {
		return nil, err
	}
	OpenContext(r)

	if bodyUserName == "" {
		pathVars := make(map[string]string)
		pathVars["uid"] = user.ID.Hex()
		SetContextVar(r, "pathVars", pathVars)
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(userPass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Name = "Username"
	user.Password = passHash
	user.HasPasswd = true
	user.Policy = &gomediacenter.UserPolicy{Admin: false}

	// Mock setup.
	db := new(mockDB)
	db.On("GetUserByID", user.ID.Hex()).Return(user, nil)
	db.On("GetUserByName", bodyUserName).Return(user, nil)
	SetContextVar(r, "db", db)

	if withHeader {
		r.Header.Add(gomediacenter.SessionAuthHeader, authTestHeader)
	}

	return r, nil
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	user := gomediacenter.NewUser("User Name")
	uid := bson.NewObjectId()
	user.ID = uid

	body := &bytes.Buffer{}
	_ = json.NewEncoder(body).Encode(user)

	r, _ := http.NewRequest(http.MethodPost, "/", body)
	OpenContext(r)
	defer CloseContext(r)

	pathVars := make(map[string]string)
	pathVars["uid"] = user.ID.Hex()
	SetContextVar(r, "pathVars", pathVars)

	db := new(mockDB)
	db.On("UpdateUser", uid.Hex(), user).Return(nil)
	SetContextVar(r, "db", db)

	recorder := httptest.NewRecorder()
	UpdateUser(recorder, r)

	assert.Equal(http.StatusOK, recorder.Code, "Incorrect status code.")
}

func TestUpdateUserPolicy(t *testing.T) {
	assert := assert.New(t)
	uid := "uid"
	policy := &gomediacenter.UserPolicy{Admin: true}
	body := &bytes.Buffer{}
	_ = json.NewEncoder(body).Encode(policy)

	r, _ := http.NewRequest(http.MethodPost, "/", body)
	OpenContext(r)
	defer CloseContext(r)

	pathVars := make(map[string]string)
	pathVars["uid"] = uid
	SetContextVar(r, "pathVars", pathVars)

	db := new(mockDB)
	db.On("UpdateUserPolicy", uid, policy).Return(nil)
	SetContextVar(r, "db", db)

	recorder := httptest.NewRecorder()
	UpdateUserPolicy(recorder, r)

	assert.Equal(http.StatusOK, recorder.Code, "Incorrect status code")
}
