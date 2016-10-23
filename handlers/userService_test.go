package handlers

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
	"github.com/tcm1911/gomediacenter/auth"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	assert := assert.New(t)

	db := &mockDB{}
	db.On("AddNewUser", mock.Anything).Return(nil)

	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(gomediacenter.NewUserRequest{Name: "testUser"})

	req, _ := http.NewRequest("POST", "/Users/New", b)

	// Add to context
	OpenContext(req)
	defer CloseContext(req)

	writer := httptest.NewRecorder()
	NewUser(db).ServeHTTP(writer, req)

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
	uid := gomediacenter.NewID()
	user.ID = uid

	// Request
	r, _ := http.NewRequest("GET", "/", nil)
	r = r.WithContext(gomediacenter.AddUIDToContext(r.Context(), uid))

	// Setup db mock.
	db := &mockDB{}
	db.On("GetUserByID", uid).Return(user, nil)

	recorder := httptest.NewRecorder()

	// Call http handler.
	GetUserByID(db).ServeHTTP(recorder, r)

	body, err := getBodyStringFromRecorder(recorder)
	if err != nil {
		assert.Fail("Could not read the body.")
	}

	assert.Contains(body, `"Name":"testUser"`)
	assert.Contains(body, `"id":"`+uid.String()+`"`)
}

func TestGetAllUsers(t *testing.T) {
	assert := assert.New(t)

	user1 := gomediacenter.NewUser("user1")
	user2 := gomediacenter.NewUser("user2")
	users := []*gomediacenter.User{user1, user2}

	r, _ := http.NewRequest("GET", "/", nil)
	OpenContext(r)
	defer CloseContext(r)

	db := &mockDB{}
	db.On("GetAllUsers", mock.Anything).Return(users, nil)

	queryVars := url.Values{}
	SetContextVar(r, "queryVars", queryVars)

	recorder := httptest.NewRecorder()
	GetAllUsers(db).ServeHTTP(recorder, r)

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

	user := &gomediacenter.User{ID: gomediacenter.NewID(), Name: "User", HasPasswd: false}
	r, db := passwordChangeContext(user, "", "password")
	defer CloseContext(r)

	recorder := httptest.NewRecorder()

	ChangeUserPassword(db).ServeHTTP(recorder, r)

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
		ID:        gomediacenter.NewID(),
		Name:      "User",
		HasPasswd: true,
		Password:  currentHash,
	}
	r, db := passwordChangeContext(user, currentPass, "newPassword")
	defer CloseContext(r)
	if err != nil {
		assert.Fail(err.Error())
	}

	recorder := httptest.NewRecorder()

	ChangeUserPassword(db).ServeHTTP(recorder, r)

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
		ID:        gomediacenter.NewID(),
		Name:      "User",
		HasPasswd: true,
		Password:  currentHash,
	}
	r, db := passwordChangeContext(user, "wrongPassword", "newPassword")
	defer CloseContext(r)
	if err != nil {
		assert.Fail(err.Error())
	}

	recorder := httptest.NewRecorder()

	ChangeUserPassword(db).ServeHTTP(recorder, r)

	assert.Equal(http.StatusUnauthorized, recorder.Code)
}

func passwordChangeContext(user *gomediacenter.User, currentPass, newPass string) (*http.Request, gomediacenter.UserManager) {
	// Mock setup.
	db := new(mockDB)
	db.On("ChangeUserPassword", user.ID, mock.Anything).Return(nil)
	db.On("GetUserByID", user.ID).Return(user, nil)

	// Context setup.
	r := setupPostReqTest(user.ID, gomediacenter.PasswordRequest{
		New: newPass, Current: currentPass})

	return r, db
}

func TestAuthenticateWithCorrectPassword(t *testing.T) {
	r, db, err := authenticateContextSetup("testpassword", "testpassword", "", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	manager := &auth.SimpleSessionManager{}
	Authenticate(db, manager).ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthenticateWithIncorrectPassword(t *testing.T) {
	r, db, err := authenticateContextSetup("testpassword", "incorrect", "", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	manager := &auth.SimpleSessionManager{}
	Authenticate(db, manager).ServeHTTP(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateWithoutCorrectHeader(t *testing.T) {
	r, db, err := authenticateContextSetup("testpassword", "testpassword", "", false)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	manager := &auth.SimpleSessionManager{}
	Authenticate(db, manager).ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthenticateByNameWithCorrectPassword(t *testing.T) {
	r, db, err := authenticateContextSetup("testpassword", "testpassword", "Username", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	manager := &auth.SimpleSessionManager{}
	AuthenticateByName(db, manager).ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthenticateByNameWithIncorrectPassword(t *testing.T) {
	r, db, err := authenticateContextSetup("testpassword", "incorrect", "Username", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	manager := &auth.SimpleSessionManager{}
	AuthenticateByName(db, manager).ServeHTTP(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateByNameWithoutCorrectHeader(t *testing.T) {
	r, db, err := authenticateContextSetup("testpassword", "testpassword", "Username", false)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	w := httptest.NewRecorder()
	manager := &auth.SimpleSessionManager{}
	AuthenticateByName(db, manager).ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func authenticateContextSetup(userPass, loginPass, bodyUserName string, withHeader bool) (*http.Request, gomediacenter.UserManager, error) {
	uid := gomediacenter.NewID()
	user := &gomediacenter.User{ID: uid}

	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(gomediacenter.LoginRequest{
		Name: bodyUserName, Password: loginPass})

	// Context setup.
	r, err := http.NewRequest("POST", "", b)
	if err != nil {
		return nil, nil, err
	}
	if bodyUserName == "" {
		r = r.WithContext(gomediacenter.AddUIDToContext(r.Context(), uid))
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(userPass), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	user.Name = "Username"
	user.Password = passHash
	user.HasPasswd = true
	user.Policy = &gomediacenter.UserPolicy{Admin: false}

	// Mock setup.
	db := new(mockDB)
	db.On("GetUserByID", user.ID).Return(user, nil)
	db.On("GetUserByName", bodyUserName).Return(user, nil)

	if withHeader {
		r.Header.Add(gomediacenter.SessionAuthHeader, authTestHeader)
	}

	return r, db, nil
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	user := gomediacenter.NewUser("User Name")
	uid := gomediacenter.NewID()
	user.ID = uid
	db := new(mockDB)
	db.On("UpdateUser", uid, user).Return(nil)

	r := setupPostReqTest(user.ID, user)
	defer CloseContext(r)

	recorder := httptest.NewRecorder()
	UpdateUser(db).ServeHTTP(recorder, r)

	assert.Equal(http.StatusOK, recorder.Code, "Incorrect status code.")
}

func TestUpdateUserPolicy(t *testing.T) {
	assert := assert.New(t)
	uid := gomediacenter.NewID()
	policy := &gomediacenter.UserPolicy{Admin: true}
	db := new(mockDB)
	db.On("UpdateUserPolicy", uid, policy).Return(nil)

	r := setupPostReqTest(uid, policy)
	defer CloseContext(r)

	recorder := httptest.NewRecorder()
	UpdateUserPolicy(db).ServeHTTP(recorder, r)

	assert.Equal(http.StatusOK, recorder.Code, "Incorrect status code")
}

func TestUpdateUserConfig(t *testing.T) {
	assert := assert.New(t)
	uid := gomediacenter.NewID()
	cfg := &gomediacenter.UserConfig{SubtitleMode: "test config"}
	db := new(mockDB)
	db.On("UpdateUserConfiguration", uid, cfg).Return(nil)

	r := setupPostReqTest(uid, cfg)
	defer CloseContext(r)

	w := httptest.NewRecorder()
	UpdateUserConfiguration(db).ServeHTTP(w, r)
	assert.Equal(http.StatusOK, w.Code, "Incorrect status code")
}

func setupPostReqTest(uid gomediacenter.ID, v interface{}) *http.Request {
	r, _ := gomediacenter.CreateRequestWithBody(http.MethodPost, "/", v)

	return r.WithContext(gomediacenter.AddUIDToContext(r.Context(), uid))
}
