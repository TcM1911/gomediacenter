// +build integration

package routes_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/auth"
	"github.com/tcm1911/gomediacenter/mediaserver/routes"
	"github.com/tcm1911/gomediacenter/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/ory-am/dockertest.v2"
)

const authTestHeader = `MediaBrowser Client="Emby Web Client", Device="Chrome 50.0.2661.50", DeviceId="cae2cc5be4e17f1d0a486d0c8fdb4789f4f1e99c", Version="3.0.5912.0", UserId="f40b2df070cf46e686bcbdd388d8706c"`

var serverURL string
var dbStruct *mongo.DB

func TestGetPublicUserListing(t *testing.T) {
	assert := assert.New(t)

	resp, _ := http.Get(serverURL + "/Users/Public")
	var users []gomediacenter.PublicUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		assert.Fail("Error when decoding the repsonse body: " + err.Error())
	}

	adminUserReturned, normalUserReturned := false, false
	for _, user := range users {
		switch user.Name {
		case adminName:
			adminUserReturned = adminID.Hex() == user.ID.Hex()
		case userName:
			normalUserReturned = userID.Hex() == user.ID.Hex()
		}
	}
	assert.True(adminUserReturned, "Admin user should be in the user listing.")
	assert.True(normalUserReturned, "Normal user should be in the user listing.")
}

func TestAuthenticatingUserByName(t *testing.T) {
	assert := assert.New(t)

	resp, code, err := gomediacenter.AuthenticateUserByNameReqest(
		adminName,
		adminPassword,
		serverURL,
		authTestHeader)
	if err != nil {
		assert.Fail("Error when sending login request: " + err.Error())
	}

	assert.Equal(http.StatusOK, code)

	assert.NotNil(resp.Token, "Nil token.")

	adminToken = resp.Token

	if adminToken == "" {
		t.FailNow()
	}
}

func TestAuthenticatingUserByIdAndLogout(t *testing.T) {
	assert := assert.New(t)

	resp, code, err := gomediacenter.AuthenticateUserByIDReqest(
		adminID.Hex(),
		adminPassword,
		serverURL,
		authTestHeader)
	if err != nil {
		assert.Fail("Error when sending login request:" + err.Error())
	}

	assert.Equal(http.StatusOK, code)

	token := resp.Token
	log.Println("Session token:", token)
	assert.NotNil(token, "Nil token.")

	// Logging out.
	loggedOut, err := gomediacenter.LogoutUserReq(adminID.Hex(), token, serverURL)
	if err != nil {
		assert.Fail("Failed to logout: " + err.Error())
	}
	assert.True(loggedOut)
}

func TestShouldOnlyAccessUserDataIfLoggedInAsCorrectUser(t *testing.T) {
	assert := assert.New(t)
	adminUser, code, err := gomediacenter.GetUser(adminID.Hex(), adminToken, serverURL)
	if err != nil {
		assert.Fail("Error when getting admin user data: " + err.Error())
	}
	assert.Equal(adminName, adminUser.Name)
	assert.Equal(http.StatusOK, code)

	token := bson.NewObjectId()
	adminByNormalUser, failCode, err := gomediacenter.GetUser(adminID.Hex(), token.Hex(), serverURL)
	if err == nil {
		assert.Fail("This request should fail.")
	}
	assert.Nil(adminByNormalUser, "Body should be nil.")
	assert.Equal(http.StatusUnauthorized, failCode)
}

func TestAdminCanAccessUserDataIfLoggedIn(t *testing.T) {
	assert := assert.New(t)
	user, code, err := gomediacenter.GetUser(userID.Hex(), adminToken, serverURL)
	if err != nil {
		assert.Fail("Error when getting user data: " + err.Error())
	}
	assert.Equal(userName, user.Name)
	assert.Equal(http.StatusOK, code)
}

func TestCreateNewUserAndChangeThePassword(t *testing.T) {
	assert := assert.New(t)
	un := "newUserName"
	user, err := gomediacenter.CreateUser(un, adminToken, serverURL)
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(un, user.Name)

	setPass := "setPassword"
	code, err := gomediacenter.ChangePassword(
		"",
		setPass,
		adminToken,
		user.ID.Hex(),
		serverURL)
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(http.StatusOK, code)

	resp, code, err := gomediacenter.AuthenticateUserByIDReqest(
		user.ID.Hex(),
		setPass,
		serverURL,
		authTestHeader)
	if err != nil {
		assert.Fail("Authentication failed: " + err.Error())
		return
	}
	assert.Equal(http.StatusOK, code)
	assert.Equal(un, resp.User.Name)

	// User changes the password.
	userToken := resp.Token
	changedPass := "changedPassword"
	code, err = gomediacenter.ChangePassword(
		setPass,
		changedPass,
		userToken,
		resp.User.ID.Hex(),
		serverURL)
	if err != nil {
		assert.Fail("Failed to change the user password.")
		return
	}
	assert.Equal(http.StatusOK, code)

	// User logs out.
	ok, err := gomediacenter.LogoutUserReq(resp.User.ID.Hex(), userToken, serverURL)
	if err != nil {
		assert.Fail("Failed to logout.")
		return
	}
	assert.True(ok, "Did user logout.")

	// Admin removes user account.
	code, err = gomediacenter.DeleteUser(resp.User.ID.Hex(), adminToken, serverURL)
	if err != nil {
		assert.Fail("Failed to remove the user")
		return
	}
	assert.Equal(http.StatusOK, code)
}

func TestAdminCanRequestAllUsersData(t *testing.T) {
	assert := assert.New(t)
	users, err := gomediacenter.GetAllUsers(adminToken, serverURL)
	if err != nil {
		assert.Fail("Request failed: " + err.Error())
		return
	}
	assert.Equal(2, len(users), "Should return 2 users")
}

func TestUpdateUserProfile(t *testing.T) {
	assert := assert.New(t)
	resp, code, err := gomediacenter.AuthenticateUserByIDReqest(userID.Hex(), userPassword, serverURL, authTestHeader)
	if code != http.StatusOK || err != nil {
		assert.Fail("Authentcation failed.")
		return
	}
	user := resp.User
	user.Name = "Changed name"
	code, err = gomediacenter.UpdateUser(user, user.ID.Hex(), resp.Token, serverURL)

	assert.Equal(http.StatusOK, code, "Wrong status code returned.")

	// Verify data in the db.
	dbData, code, err := gomediacenter.GetUser(userID.Hex(), resp.Token, serverURL)
	if code != http.StatusOK || err != nil {
		assert.Fail("Failed to retrieve data from db.")
		return
	}

	assert.Equal("Changed name", dbData.Name, "Wrong user name.")
	assert.Equal(user, dbData, "Profile doesn't match.")

	// Logout
	ok, err := gomediacenter.LogoutUserReq(userID.Hex(), resp.Token, serverURL)
	assert.True(ok, "Logout failed.")
}

func TestUpdateUserPolicy(t *testing.T) {
	assert := assert.New(t)

	uid := createNewUserInDB("User Policy", "password", false)
	rsp := loginUser(uid, "password", serverURL, authTestHeader)
	if rsp == nil {
		assert.FailNow("Authentication failed")
	}
	user := rsp.User
	token := rsp.Token
	p := user.Policy
	assert.False(p.Admin, "User should not be an admin.")

	// Test that user can't change the policy.
	p.Admin = true
	c, err := gomediacenter.UpdateUserPolicy(p, uid, token, serverURL)
	if err != nil {
		assert.FailNow("User's policy request failed: " + err.Error())
	}
	assert.Equal(http.StatusUnauthorized, c, "User should not be allowed to change the policy")

	// Admin can change the user's policy.
	cp, err := gomediacenter.UpdateUserPolicy(p, uid, adminToken, serverURL)
	if err != nil {
		assert.FailNow("Admin's policy request failed: " + err.Error())
	}
	r, c, err := gomediacenter.GetUser(uid, adminToken, serverURL)
	if err != nil {
		assert.FailNow("Getting user profile failed: " + err.Error())
	}
	assert.Equal(http.StatusOK, cp, "Wrong status code returned.")
	assert.Equal(http.StatusOK, c, "Wrong status code returned.")
	assert.True(r.Policy.Admin, "User's should now be an admin.")
}

func TestUpdateUserConfig(t *testing.T) {
	assert := assert.New(t)
	uid := createNewUserInDB("Update Config", "password", false)
	rsp := loginUser(uid, "password", serverURL, authTestHeader)
	if rsp == nil {
		assert.FailNow("Authentication failed")
	}
	token := rsp.Token
	cfg := rsp.User.Configuration
	cfg.SubtitleMode = "SubtitleMode has been updated"

	c, err := gomediacenter.UpdateUserCfg(cfg, uid, token, serverURL)
	if err != nil {
		assert.Fail("Error when sending the change request: " + err.Error())
	}
	if c != http.StatusOK {
		assert.Fail("Wrong status code")
	}

	usr, code, err := gomediacenter.GetUser(uid, token, serverURL)

	assert.Equal(http.StatusOK, code, "Getting updated user failed.")
	assert.NoError(err, "Error returned")
	assert.Equal(cfg.SubtitleMode, usr.Configuration.SubtitleMode, "SubtitleMode not updated")

	gomediacenter.LogoutUserReq(uid, token, serverURL)
}

func TestMain(m *testing.M) {
	db := &mongo.DB{}
	log.Println("Starting docker container...")
	c, err := dockertest.ConnectToMongoDB(15, 5*time.Second, func(url string) bool {
		log.Println("Connecting to the database...")
		db.Connect(url)

		log.Println("Starting API server")
		handler := routes.NewAPIRouter(db, db)
		server := httptest.NewServer(handler)
		serverURL = server.URL

		return true
	})
	defer db.Close()
	if err != nil {
		log.Fatalln("Error while starting up the container:", err)
	}
	log.Println("Connected to the database.")

	dbStruct = db

	// DB setup
	setupDB(db)

	// Start session manager
	shutdown := auth.Run(db)

	// Run tests
	resutls := m.Run()

	// Teardown
	log.Println("Closing down and removing the container.")
	shutdown <- struct{}{}
	<-shutdown
	if err := c.KillRemove(); err != nil {
		log.Println("Error when tearing down the container:", err)
	}

	// Exit
	os.Exit(resutls)
}

const (
	adminName     = "admin"
	adminPassword = "adminpassword"
	userName      = "normaluser"
	userPassword  = "userpassword"
)

var adminID bson.ObjectId
var adminToken string
var userID bson.ObjectId

func setupDB(db *mongo.DB) {

	// Admin user
	adminUser := setupUser(adminName, adminPassword, true)
	adminID = adminUser.ID
	if err := db.AddNewUser(adminUser); err != nil {
		log.Fatalln("Error while adding admin user to the database:", err)
	}

	// Normal user
	user := setupUser(userName, userPassword, false)
	userID = user.ID
	if err := db.AddNewUser(user); err != nil {
		log.Fatalln("Error while adding normal user to the database:", err)
	}
}

func setupUser(u, pw string, admin bool) *gomediacenter.User {
	user := gomediacenter.NewUser(u)
	user.Policy.Admin = admin
	password, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Error when generating password for", u)
	}
	user.Password = password
	user.HasPasswd = admin
	return user
}

func createNewUserInDB(u, pw string, admin bool) string {
	user := setupUser(u, pw, admin)
	if err := dbStruct.AddNewUser(user); err != nil {
		log.Fatalln("Error while adding normal user to the database:", err)
	}
	return user.ID.Hex()
}

func loginUser(uid, pw, url, header string) *gomediacenter.AuthUserResponse {
	resp, code, err := gomediacenter.AuthenticateUserByIDReqest(uid, pw, url, header)
	if code != http.StatusOK || err != nil {
		return nil
	}
	return resp
}
