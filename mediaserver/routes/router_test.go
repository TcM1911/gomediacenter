// +build integration

package routes_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers/auth"
	"github.com/tcm1911/gomediacenter/mediaserver/routes"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/ory-am/dockertest.v2"
)

const authTestHeader = `MediaBrowser Client="Emby Web Client", Device="Chrome 50.0.2661.50", DeviceId="cae2cc5be4e17f1d0a486d0c8fdb4789f4f1e99c", Version="3.0.5912.0", UserId="f40b2df070cf46e686bcbdd388d8706c"`

var serverURL string

func init() {
	handler := routes.NewAPIRouter()
	server := httptest.NewServer(handler)
	serverURL = server.URL
}

func testAPIGetRoutes(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		url  string
		code int
		body string
	}{
		{"/Albums/testID/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Albums/testID/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Artists/testID/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/testName/InstantMix", http.StatusNotImplemented, "Not yet implemented"},
		{"/Artists/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Audio/testID/stream", http.StatusNotImplemented, "Not yet implemented"},
		{"/Audio/testID/stream.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Audio/testID/hls/segmentId/stream.mp3", http.StatusNotImplemented, "Not yet implemented"},
		{"/Audio/testID/hls1/playlistId/segmentId.ts", http.StatusNotImplemented, "Not yet implemented"},

		{"/Auth/Keys", http.StatusNotImplemented, "Not yet implemented"},

		{"/Branding/Configuration", http.StatusNotImplemented, "Not yet implemented"},
		{"/Branding/Css", http.StatusNotImplemented, "Not yet implemented"},

		{"/Channels", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/Features", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/Folder", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/testID/Features", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/testID/Items", http.StatusNotImplemented, "Not yet implemented"},
		{"/Channels/testID/Latest", http.StatusNotImplemented, "Not yet implemented"},

		{"/Collections", http.StatusNotImplemented, "Not yet implemented"},

		{"/Connect/Pending", http.StatusNotImplemented, "Not yet implemented"},
		{"/Connect/Exchange", http.StatusNotImplemented, "Not yet implemented"},

		{"/Devices", http.StatusNotImplemented, "Not yet implemented"},
		{"/Devices/CameraUploads", http.StatusNotImplemented, "Not yet implemented"},
		{"/Devices/Info", http.StatusNotImplemented, "Not yet implemented"},
		{"/Devices/Capabilities", http.StatusNotImplemented, "Not yet implemented"},

		{"/DisplayPreferences/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/Dlna/testId/description", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/contentdirectory/contentdirectory.xml", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/connectionmanager/connectionmanager.xml", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/mediareceiverregistrar/mediareceiverregistrar.xml", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/testId/icons/testfilename", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/ProfileInfos", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/Profiles/Default", http.StatusNotImplemented, "Not yet implemented"},
		{"/Dlna/Profiles/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/Environment/DirectoryContents", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/NetworkShares", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/NetworkDevices", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/Drives", http.StatusNotImplemented, "Not yet implemented"},
		{"/Environment/ParentPath", http.StatusNotImplemented, "Not yet implemented"},

		{"/Games/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Games/SystemSummaries", http.StatusNotImplemented, "Not yet implemented"},
		{"/Games/PlayerIndex", http.StatusNotImplemented, "Not yet implemented"},

		{"/GameGenres", http.StatusNotImplemented, "Not yet implemented"},
		{"/GameGenres/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/GameGenres/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/GameGenres/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Genres", http.StatusNotImplemented, "Not yet implemented"},
		{"/Genres/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/Genres/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Genres/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Images/General", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/Ratings", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/MediaInfo", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/General/testName/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/Ratings/testTheme/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/MediaInfo/testTheme/testName", http.StatusNotImplemented, "Not yet implemented"},
		{"/Images/Remote", http.StatusNotImplemented, "Not yet implemented"},

		{"/Items/Counts", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/Filters", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/Ancestors", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/CriticReviews", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/Download", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/File", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/ThemeMedia", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/ThemeSongs", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testID/ThemeVideos", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/YearIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/Images/testType/testIndex/tag/format/maxwidth/maxheight/percentplayed/unplayedcount", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/InstantMix", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/ExternalIdInfos", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/RemoteSearch/Image", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/MetadataEditor", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/PlaybackInfo", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteImages", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteImages/Providers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteSearch/Subtitles/Providers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Items/testId/RemoteSearch/Subtitles/testlang", http.StatusNotImplemented, "Not yet implemented"},

		{"/Library/FileOrganization", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/FileOrganization/SmartMatches", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/MediaFolders", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/PhysicalPaths", http.StatusNotImplemented, "Not yet implemented"},
		{"/Library/VirtualFolders", http.StatusNotImplemented, "Not yet implemented"},

		{"/LiveTv/Info", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Channels", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Channels/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings/Groups", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Timers/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Timers/Defaults", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Timers", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Programs", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Programs/Recommended", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Programs/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/SeriesTimers/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/SeriesTimers", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Recordings/Groups/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/GuideInfo", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Folder", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/ListingProviders/Lineups", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/ListingProviders/SchedulesDirect/Countries", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/Registration", http.StatusNotImplemented, "Not yet implemented"},
		{"/LiveTv/TunerHosts/Satip/IniMappings", http.StatusNotImplemented, "Not yet implemented"},

		{"/Localization/Cultures", http.StatusNotImplemented, "Not yet implemented"},
		{"/Localization/Countries", http.StatusNotImplemented, "Not yet implemented"},
		{"/Localization/ParentalRatings", http.StatusNotImplemented, "Not yet implemented"},
		{"/Localization/Options", http.StatusNotImplemented, "Not yet implemented"},

		{"/Movies/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Movies/Recommendations", http.StatusNotImplemented, "Not yet implemented"},

		{"/MusicGenres/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres/testName/InstantMix", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres", http.StatusNotImplemented, "Not yet implemented"},
		{"/MusicGenres/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/News/Product", http.StatusNotImplemented, "Not yet implemented"},

		{"/Notifications/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Notifications/testId/Summary", http.StatusNotImplemented, "Not yet implemented"},
		{"/Notifications/Types", http.StatusNotImplemented, "Not yet implemented"},
		{"/Notifications/Services", http.StatusNotImplemented, "Not yet implemented"},

		{"/Packages/Reviews/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/Persons/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Persons/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Persons", http.StatusNotImplemented, "Not yet implemented"},
		{"/Persons/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Playback/BitrateTest", http.StatusNotImplemented, "Not yet implemented"},

		{"/Playlists/testId/Items", http.StatusNotImplemented, "Not yet implemented"},
		{"/Playlists/testId/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Providers/Chapters", http.StatusNotImplemented, "Not yet implemented"},
		{"/Providers/Subtitles/Subtitles/testId", http.StatusNotImplemented, "Not yet implemented"},

		{"/ScheduledTasks/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/ScheduledTasks", http.StatusNotImplemented, "Not yet implemented"},

		{"/Search/Hints", http.StatusNotImplemented, "Not yet implemented"},

		{"/Sessions", http.StatusNotImplemented, "Not yet implemented"},

		{"/Shows/NextUp", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/Upcoming", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/testId/Episodes", http.StatusNotImplemented, "Not yet implemented"},
		{"/Shows/testId/Seasons", http.StatusNotImplemented, "Not yet implemented"},

		{"/Social/Shares/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Social/Shares/Public/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Social/Shares/Public/testId/Images", http.StatusNotImplemented, "Not yet implemented"},
		{"/Social/Shares/Public/testId/Item", http.StatusNotImplemented, "Not yet implemented"},

		{"/Songs/testId/InstantMix", http.StatusNotImplemented, "Not yet implemented"},

		{"/Startup/Complete", http.StatusNotImplemented, "Not yet implemented"},
		{"/Startup/Configuration", http.StatusNotImplemented, "Not yet implemented"},
		{"/Startup/User", http.StatusNotImplemented, "Not yet implemented"},

		{"/Studios/testName/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Studios/testName/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Studios", http.StatusNotImplemented, "Not yet implemented"},
		{"/Studios/testName", http.StatusNotImplemented, "Not yet implemented"},

		{"/Sync/Jobs/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/JobItems", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Jobs", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Targets", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Options", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/JobItems/testId/File", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/JobItems/testId/AdditionalFiles", http.StatusNotImplemented, "Not yet implemented"},
		{"/Sync/Items/Ready", http.StatusNotImplemented, "Not yet implemented"},

		{"/System/ActivityLog/Entries", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Configuration/testKey", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Configuration/MetadataOptions/Default", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Configuration/MetadataPlugins", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Info", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Info/Public", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Logs", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Logs/Log", http.StatusNotImplemented, "Not yet implemented"},
		{"/System/Endpoint", http.StatusNotImplemented, "Not yet implemented"},

		{"/Trailers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Trailers/testId/Similar", http.StatusNotImplemented, "Not yet implemented"},

		{"/Users", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/Public", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Offline", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Views", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/SpecialViewOptions", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/GroupingOptions", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testId/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/Root", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId/Intros", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId/LocalTrailers", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/testId/SpecialFeatures", http.StatusNotImplemented, "Not yet implemented"},
		{"/Users/testUid/Items/Latest", http.StatusNotImplemented, "Not yet implemented"},

		{"/Videos/testId/master.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testID/hls/playlistId/teststream.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testID/hls/playlistId/segmentId.ts", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testID/hls1/playlistId/segmentId.ts", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/master.mpd", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/dash/testrepId/testSegmentId.m4s", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/Subtitles/testIndex", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/streamFormat", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/startPos/streamFormat", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/mediaSrc/Subtitles/testIndex/subtitles.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/live.m3u8", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/stream.mkv", http.StatusNotImplemented, "Not yet implemented"},
		{"/Videos/testId/AdditionalParts", http.StatusNotImplemented, "Not yet implemented"},

		{"/Years", http.StatusNotImplemented, "Not yet implemented"},
		{"/Years/testYear", http.StatusNotImplemented, "Not yet implemented"},
		{"/Years/testYear/Images/testType", http.StatusNotImplemented, "Not yet implemented"},
		{"/Years/testYear/Images/testType/testIndex", http.StatusNotImplemented, "Not yet implemented"},
	}

	for _, test := range tests {
		url := serverURL + test.url

		fmt.Printf("Testing GET request to %s\n", url)
		resp, _ := http.Get(url)

		reader := resp.Body
		body, _ := ioutil.ReadAll(reader)
		reader.Close()

		if resp.StatusCode != test.code {
			t.Errorf("Wrong status code for %s. Got: %d, expected %d", test.url, resp.StatusCode, test.code)
		}

		if string(body) != test.body {
			t.Errorf("Wrong reesponse body for %s. Got: %s, expected %s", test.url, body, test.body)
		}
	}

}

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

func TestMain(m *testing.M) {
	log.Println("Starting docker container...")
	c, err := dockertest.ConnectToMongoDB(15, 5*time.Second, func(url string) bool {
		log.Println("Connecting to the database...")
		db.Connect(url)
		session := db.GetDBSession()
		return session.Ping() == nil
	})
	defer db.Disconnect()
	if err != nil {
		log.Fatalln("Error while starting up the container:", err)
	}
	log.Println("Connected to the database.")

	// DB setup
	setupDB()

	// Start session manager
	shutdown := auth.Run()

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

func setupDB() {
	db := db.GetDB()
	defer db.Close()

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
	d := db.GetDB()
	defer d.Close()

	user := setupUser(u, pw, admin)
	if err := d.AddNewUser(user); err != nil {
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
