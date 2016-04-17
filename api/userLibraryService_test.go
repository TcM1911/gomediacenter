package api

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/tcm1911/gomediacenter"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
)

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
	videoSource.Path = "/path/to/file.mkv"
	videoSource.Container = "mkv"
	videoSource.MediaStreams = []interface{}{vStream, aStream, }

	id := bson.NewObjectId()
	video := new(gomediacenter.Video)
	video.Name = "Test Movie Title"
	video.Id = id
	video.MediaSources = []interface{}{videoSource}

	actor := gomediacenter.Person{Id: "id", Name: "Actor name", Role: "Actor"}

	movie = new(gomediacenter.Movie)
	movie.Video = video
	movie.ImdbId = "imdbID"
	movie.People = []gomediacenter.Person{actor}
}

var movie *gomediacenter.Movie

// Setup db mock
type mockDB struct {
	mock.Mock
}

func (m *mockDB) FindItemById(s string) (gomediacenter.MEDIATYPE, interface{}, error) {
	args := m.Called(s)
	return args.Get(0).(gomediacenter.MEDIATYPE), args.Get(1).(interface{}), args.Error(2)
}

func TestUserItemHandler(t *testing.T) {
	assert := assert.New(t)

	database := new(mockDB)
	database.On("FindItemById", "12345").Return(gomediacenter.MOVIE, movie, nil)

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
}
