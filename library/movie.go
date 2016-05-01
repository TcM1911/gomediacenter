package library

// Videos x264
// http://scenenotice.org/details.php?id=2187
// Movie.Name.YEAR.PROPER/READ.NFO/REPACK.BDRip/DVDRip.x264-GROUP
// TV.Show.SxxExx.PROPER/READ.NFO/REPACK.BDRip/DVDRip.x264-GROUP
// FESTIVAL, STV, LIMITED or TV
// WS/FS (rules above), PROPER, REPACK, RERIP, REAL, RETAIL, EXTENDED, REMASTERED,
// RATED, UNRATED, CHRONO, THEATRICAL, DC, SE, UNCUT, INTERNAL, DUBBED, SUBBED, FINAL, COLORIZED

// Videos HD x264
// https://scenerules.org/t.html?id=2011_X264.2.nfo
// Movie.Name.YEAR.<720p/1080p>.BluRay.x264.<PROPER/READ.NFO/REPACK>-GROUP
// TV.Show.SxxExx.<720p/1080p>.BluRay.x264.<PROPER/READ.NFO/REPACK>-GROUP

// Videos Xvid
// https://scenerules.org/t.html?id=2002_XViD.nfo
// Movie.Name.CD1.Year.Source.Codec-Group.avi
// Movie.Name.Year.Source.Codec-Group.avi
// https://scenerules.org/t.html?id=2001_XViD.nfo
// Movie.Title.Year.DVDrip.DivX-GROUP
// Movie.Title.Year.ANIME.DivX-GROUP
// Movie.Title.DVDrip.DivX-GROUP
// Movie.Title.ANIME.DivX-GROUP

import (
	"regexp"
	"strconv"
	"strings"
)

////////////
// Fields //
////////////

// Example: Movie.Name.YEAR.PROPER/READ.NFO/REPACK.BDRip/DVDRip.x264-GROUP
var sdMovieRegex = regexp.MustCompile(`([\w*\.*]*)\.(\d{4})\.([\w*\.*]*)\.*(DVDRip|BDRip)\.(Xvid|x264)-\w+`)

// Example: Movie.Name.YEAR.<720p/1080p>.BluRay.x264.<PROPER/READ.NFO/REPACK>-GROUP
var hdMovieRegex = regexp.MustCompile(`([\w*\.*]*)\.(\d{4})\.(720p|1080p)\.BluRay\.x264\.*([\w*\.*]*)-\w+`)

var acceptedTags []string = []string{"REAL", "PROPER", "REPACK", "RERIP",
	"WS", "FS", "RETAIL", "EXTENDED", "REMASTERED", "RATED", "UNRATED",
	"CHRONO", "THEATRICAL", "DC", "SE", "UNCUT", "INTERNAL", "DUBBED",
	"SUBBED", "FINAL", "COLORIZED", "FESTIVAL", "STV", "LIMITED", "TV", "READ.NFO"}

/////////////
// Structs //
/////////////

// The struct returned when SD Movie regex is matched.
type SDMovieResult struct {
	Name   string
	Year   int
	Source string
	Codec  string
	Tags   []string
}

// The struct returned when HD Movie regex is matched.
type HDMovieResult struct {
	Name       string
	Year       int
	Resolution string
	Tags       []string
}

////////////
// Public //
////////////

func MatchToSDMovieRegex(s string) (SDMovieResult, bool) {
	var result SDMovieResult
	if match := sdMovieRegex.FindStringSubmatch(s); match != nil {
		result.Name = strings.Join(strings.Split(match[1], "."), " ")
		if year, err := strconv.Atoi(match[2]); err == nil {
			result.Year = year
		}
		result.Codec = match[5]
		result.Source = match[4]
		result.Tags = extracTags(match[3])
		return result, true
	}
	return result, false
}

func MatchToHDMovieRegex(s string) (HDMovieResult, bool) {
	var result HDMovieResult
	if match := hdMovieRegex.FindStringSubmatch(s); match != nil {
		result.Name = strings.Join(strings.Split(match[1], "."), " ")
		if year, err := strconv.Atoi(match[2]); err == nil {
			result.Year = year
		}
		result.Resolution = match[3]
		result.Tags = extracTags(match[4])
		return result, true
	}
	return result, false
}

/////////////
// Private //
/////////////

func extracTags(s string) []string {
	var extractedTags []string

	//TODO: Add support for REAL.REAL tags.

	for _, tag := range acceptedTags {
		if strings.Contains(s, tag) {
			extractedTags = append(extractedTags, tag)
		}
	}
	return extractedTags
}
