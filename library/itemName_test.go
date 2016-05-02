package library

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldDetermineIfFileIsAVideoFile(t *testing.T) {
	cases := []struct {
		fileName    string
		isVideoFile bool
	}{
		{`The Movie Title (2222).mkv`, true},
		{`The.Movie.Title.2223.DVDRip-Name.avi`, true},
		{`The_Movie_Title-2222-Rip-Something.mov`, true},
		{`02-Artist.Name-Song.Title.mp3`, false},
		{`cover.jpg`, false},
		{`The Movie_title-[2000].mp4`, true},
		{`File without an extension`, false},
	}

	for i, test := range cases {
		isVideoFile := IsVideoFile(test.fileName)
		errorMsg := "Case number " + string(i) + " failed."
		assert.Equal(t, test.isVideoFile, isVideoFile, errorMsg)
	}
}

func TestExtractionOfMInformationFromTheFileName(t *testing.T) {
	cases := []struct {
		fileName string
		title    string
		year     int
	}{
		{`Movie.Name.2016.PROPER.BDRip.x264-GROUP`, "Movie Name", 2016},
		{`Movie.Name.2016.BDRip.x264-GROUP`, "Movie Name", 2016},
		{`Movie.Name.2015.1080p.BluRay.x264.REPACK-GROUP`, "Movie Name", 2015},
		{`Movie.Name.2015.1080p.BluRay.x264-GROUP`, "Movie Name", 2015},
	}

	for i, test := range cases {
		name, year := ParseMovieInfo(test.fileName)
		errorMsg := "Case number " + strconv.Itoa(i) + " failed."
		assert.Equal(t, test.title, name, errorMsg)
		assert.Equal(t, test.year, year, errorMsg)
	}
}

func TestGetRelativePath(t *testing.T) {
	assert := assert.New(t)

	tests := []struct{ abs, libroot, rel string }{
		{"/abs/lib/path/action/movie.mkv", "/abs/lib/path", "action/movie.mkv"},
		{"/abs/lib/path/action/movie.mkv", "/abs/lib/path/", "action/movie.mkv"},
	}

	for _, test := range tests {
		rel, err := GetRelativePath(test.abs, test.libroot)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(test.rel, rel)
	}
}
