package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tcm1911/gomediacenter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

///////////
// Setup //
///////////

func init() {
	// Create movie data.
	vStream := new(gomediacenter.VideoStream)
	vStream.Codec = "testVideoCodec"
	vStream.AverageFrameRate = 23.921
	vStream.Cabac = true

	aStream := new(gomediacenter.AudioStream)
	aStream.Codec = "testAudioCodec"
	aStream.Channels = 2

	videoSource := new(gomediacenter.VideoSource)
	videoSource.Container = "mkv"
	videoSource.MediaStreams = []interface{}{vStream, aStream}
	videoSource.Chapters = []gomediacenter.Chapter{{Name: "Chapter 1", StartPos: 0}, {Name: "Chapter 2", StartPos: 2000}}

	id := bson.NewObjectId()
	video := gomediacenter.Video{}
	video.MediaSources = []interface{}{videoSource}

	actor := gomediacenter.Person{Id: "id", Name: "Actor name", Role: "Actor"}

	movie = new(gomediacenter.Movie)
	movie.Path = "/path/to/file.mkv"
	movie.Name = "Test Movie Title"
	movie.Id = id
	movie.Video = video
	movie.ImdbId = "imdbID"
	movie.People = []gomediacenter.Person{actor}

	// Create item user data
	userdata = gomediacenter.NewItemUserData("12345", "userId")
	userdata.Played = true
	userdata.PlayCount = 1
	userdata.LastPlayedDate = time.Now().UTC()
}

var movie *gomediacenter.Movie
var userdata *gomediacenter.ItemUserData

///////////
// Tests //
///////////

func TestUserItemHandler(t *testing.T) {
	assert := assert.New(t)

	database := new(mockDB)
	database.On("FindItemByID", "12345").Return(gomediacenter.MOVIE, movie, nil)
	database.On("FindItemUserData", "userid", "12345").Return(userdata, nil)

	// PathVars
	vars := make(map[string]string)
	vars["uid"] = "userid"
	vars["id"] = "12345"

	request, _ := http.NewRequest("GET", "/test", nil)

	// Add to context
	OpenContext(request)
	defer CloseContext(request)
	SetContextVar(request, "db", database)
	SetContextVar(request, "pathVars", vars)

	recorder := httptest.NewRecorder()
	UserItemHandler(recorder, request)

	assert.Equal(http.StatusOK, recorder.Code)

	p, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		assert.Fail("Error reading the body")
	}
	body := string(p)

	assert.Contains(body, `"Name":"Test Movie Title",`)
	assert.Contains(body, `"Container":"mkv",`)
	assert.Contains(body, `"Path":"/path/to/file.mkv",`)
	assert.Contains(body, `"Codec":"testVideoCodec",`)
	assert.Contains(body, `"AverageFrameRate":23.921,`)
	assert.Contains(body, `"Codec":"testAudioCodec",`)
	assert.Contains(body, `"Channels":2,`)
	assert.Contains(body, `"ImdbId":"imdbID",`)
	assert.Contains(body, `"People":[{"Name":"Actor name","Id":"id","Role":"Actor"`)
	assert.Contains(body, `"Chapters":[{"StartPositionTicks":0,"Name":"Chapter 1"},{"StartPositionTicks":2000,"Name":"Chapter 2"}]`)

	assert.Contains(body, `"UserData":{"PlayedPercentage":0,"PlaybackPositionTicks":0`)
	assert.Contains(body, `"Played":true,`)
	assert.Contains(body, `"PlayCount":1,`)
	assert.NotContains(body, "userId")
}

// Test that an empty struct is returned if no intros are returned from the database.
func TestUserItemIntrosHandlerNoEntry(t *testing.T) {
	assert := assert.New(t)

	database := new(mockDB)
	var array []*gomediacenter.Intro
	database.On("FindItemIntro", "12345").Return(array, mgo.ErrNotFound)

	// PathVars
	vars := make(map[string]string)
	vars["uid"] = "userid"
	vars["id"] = "12345"

	request, _ := http.NewRequest("GET", "/test", nil)

	// Add to context
	OpenContext(request)
	defer CloseContext(request)
	SetContextVar(request, "db", database)
	SetContextVar(request, "pathVars", vars)

	recorder := httptest.NewRecorder()
	UserItemIntrosHandler(recorder, request)

	assert.Equal(http.StatusOK, recorder.Code)

	p, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		assert.Fail("Error reading the body")
	}
	body := string(p)

	assert.Contains(body, `{"TotalRecordCount":0}`)
}

//////////////////////
// Helper functions //
//////////////////////
