package library

import (
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
