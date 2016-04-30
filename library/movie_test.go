package library

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestMatchToSDMovieRegex(t *testing.T) {
	tests := []struct {
		Folder string
		Title string
		Year int
		Source string
		Codec string
		Tags []string
	}{
		{"Movie.Name.2016.PROPER.BDRip.x264-GROUP",
			"Movie Name", 2016, "BDRip", "x264", []string{"PROPER"}},
		{"The.Movie.Title.2016.REAL.PROPER.DVDRip.Xvid-GROUP",
			"The Movie Title", 2016, "DVDRip", "Xvid", []string{"REAL", "PROPER"}},
		{"The.Movie.Title.2016.DVDRip.Xvid-GROUP",
			"The Movie Title", 2016, "DVDRip", "Xvid", nil},
	}

	for i, test := range tests {
		result, _ := MatchToSDMovieRegex(test.Folder)
		num := strconv.Itoa(i)
		errorMsg := "Case number " + num + " failed."
		assert.Equal(t, test.Title, result.Name, errorMsg)
		assert.Equal(t, test.Year, result.Year, errorMsg)
		assert.Equal(t, test.Source, result.Source, errorMsg)
		assert.Equal(t, test.Codec, result.Codec, errorMsg)
		assert.Equal(t, test.Tags, result.Tags, errorMsg)
	}
}

func TestMatchToHDMovieRegex(t *testing.T) {
	tests := []struct {
		Folder string
		Title string
		Year int
		Resolution string
		Tags []string
	}{
		{"Movie.Name.2016.720p.BluRay.x264.READ.NFO-GROUP",
			"Movie Name", 2016, "720p", []string{"READ.NFO"}},
		{"The.Movie.Name.2016.1080p.BluRay.x264.REAL.PROPER-GROUP",
			"The Movie Name", 2016, "1080p", []string{"REAL", "PROPER"}},
		{"The.Movie.Name.2016.1080p.BluRay.x264-GROUP",
			"The Movie Name", 2016, "1080p", nil},
	}

	for i, test := range tests {
		result, _ := MatchToHDMovieRegex(test.Folder)
		num := strconv.Itoa(i)
		errorMsg := "Case number " + num + " failed."
		assert.Equal(t, test.Title, result.Name, errorMsg)
		assert.Equal(t, test.Year, result.Year, errorMsg)
		assert.Equal(t, test.Resolution, result.Resolution, errorMsg)
		assert.Equal(t, test.Tags, result.Tags, errorMsg)
	}
}
