package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2/bson"
)

func TestNewUser(t *testing.T) {
	assert := assert.New(t)

	db := new(mockDB)
	db.On("AddNewUser", mock.Anything).Return(nil)

	postBody := url.Values{}
	postBody.Add("Name", "testUser")

	req, _ := http.NewRequest("POST", "/Users/New", nil)
	req.PostForm = postBody

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

func TestGetUserById(t *testing.T) {
	assert := assert.New(t)

	// Create test user.
	user := gomediacenter.NewUser("testUser")
	uid := bson.NewObjectId()
	user.Id = uid

	// Request
	r, _ := http.NewRequest("GET", "/", nil)
	OpenContext(r)

	// Setup db mock.
	db := new(mockDB)
	db.On("GetUserById", uid.Hex()).Return(user, nil)
	SetContextVar(r, "db", db)

	// Setup pathVars
	pathVars := make(map[string]string)
	pathVars["uid"] = uid.Hex()
	SetContextVar(r, "pathVars", pathVars)

	recorder := httptest.NewRecorder()

	// Call http handler.
	GetUserById(recorder, r)

	body, err := getBodyStringFromRecorder(recorder)
	if err != nil {
		assert.Fail("Could not read the body.")
	}

	assert.Contains(body, `"Name":"testUser"`)
	assert.Contains(body, `"id":"`+uid.Hex()+`"`)
}
