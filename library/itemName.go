package library

import (
	"strings"
)

// IsVideoFile determines if a file is a video file based on the file extension.
// The function will only return true if the file extension matches predetermined
// file extensions.
func IsVideoFile(fileName string) bool {
	fileNameSlice := strings.Split(fileName, ".")
	ext := fileNameSlice[len(fileNameSlice) - 1]

	_, ok := videoFileExtensions[ext]
	return ok
}

// ParseMovieInfo gets the movie name and year from the file name or folder name.
func ParseMovieInfo(s string) (string, int) {
	// First try scene regex.
	if name, year := matchSceneRegex(s); name != "" && year != 0 {
		return name, year
	}
	return "", 0
}

func matchSceneRegex(s string) (string, int) {
	// Try SD regex first.
	if match, ok := MatchToSDMovieRegex(s); ok {
		return match.Name, match.Year
	}
	// Try HD regex.
	if match, ok := MatchToHDMovieRegex(s); ok {
		return match.Name, match.Year
	}
	return "", 0
}

// TV x264
//http://scenenotice.org/details.php?id=2328
// Tags:
// ALTERNATIVE.CUT, CONVERT, COLORIZED, DC, DIRFIX, DUBBED, EXTENDED, FINAL, INTERNAL, NFOFIX,
// OAR, OM, PPV, PROPER, REAL, REMASTERED, READNFO, REPACK, RERIP, SAMPLEFIX, SOURCE.SAMPLE,
// SUBBED, UNCENSORED, UNRATED, UNCUT, WEST.FEED, and WS
//
// Format:
// AHDTV, HDTV, APDTV, PDTV, ADSR, DSR
//
// Names
// Single.Episode.Special.YYYY.TAGS.[LANGUAGE].FORMAT-GROUP
// Weekly.TV.Show.SXXEXX[Episode.Part].[Episode.Title].TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Weekly.TV.Show.Special.SXXE00.Special.Title.TAGS.[LANGUAGE].FORMAT-GROUP
// Multiple.Episode.TV.Show.SXXEXX-EXX[Episode.Part].[Episode.Title].TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Miniseries.Show.Name.Part.X.[Episode.Title]. TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Daily.TV.Show.YYYY.MM.DD.[Guest.Name].TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Daily.Sport.League.YYYY.MM.DD.Event.TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Monthly.Competition.YYYY.MM.Event.TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Yearly.Competition.YYYY.Event.TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Sports.Match.YYYY[-YY].Round.XX.Event.[Team.vs.Team].TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Sport.Tournament.YYYY.Event.[Team/Person.vs.Team/Person]. TAGS.[LANGUAGE].FORMAT.x264-GROUP
// Country.YYYY.Event.BROADCASTER.FEED.TAGS.[LANGUAGE].FORMAT.x264-GROUP

// TV HD
// https://scenerules.org/t.html?id=2007_TV_X264.nfo
// Show.Name.SXXEXX.720p.HDTV.x(X)264-GROUP
// Show.Name.YYYY-MM-DD.720p.HDTV.x(X)264-GROUP
// http://scenenotice.org/details.php?id=2098
// Show.Name.SXXEXX.Episode.Title.1080p.HDTV.x264-GROUP
// Show.Name.YYYY.MM.DD.Guest.Name.1080p.HDTV.x264-GROUP
// Episode title and guest name are optional
// Show.Name.PartXX.1080p.HDTV.x264-GROUP for miniseries
// League.YYYY.MM.DD.Event.EXTRA.TAGS.1080p.HDTV.x264-GROUP
// Competition.YYYY-MM.Event.EXTRA.TAGS.1080p.HDTV.x264-GROUP
// EPL.2010.01.01.Manchester.United.vs.Arsenal.1080p.HDTV.x264-GROUP
// TNA.Impact.2010.03.02.1080p.HDTV.x264-GROUP
// WWE.WrestleMania.2010.PPV.1080p.HDTV.x264-GROUP
// Tennis.US.Open.2011.Final.Player1.vs.Player2.1080p.HDTV.x264-GROUP



///////////////////////
// Private Variables //
///////////////////////

var videoFileExtensions = map[string]struct{}{
	"m4v": struct{}{},
	"3gp": struct{}{},
	"nsv": struct{}{},
	"ts": struct{}{},
	"ty": struct{}{},
	"strm": struct{}{},
	"rm": struct{}{},
	"rmvb": struct{}{},
	"m3u": struct{}{},
	"ifo": struct{}{},
	"mov": struct{}{},
	"qt": struct{}{},
	"divx": struct{}{},
	"xvid": struct{}{},
	"bivx": struct{}{},
	"vob": struct{}{},
	"nrg": struct{}{},
	"img": struct{}{},
	"iso": struct{}{},
	"pva": struct{}{},
	"wmv": struct{}{},
	"asf": struct{}{},
	"asx": struct{}{},
	"ogm": struct{}{},
	"m2v": struct{}{},
	"avi": struct{}{},
	"bin": struct{}{},
	"dat": struct{}{},
	"dvr-ms": struct{}{},
	"mpg": struct{}{},
	"mpeg": struct{}{},
	"mp4": struct{}{},
	"mkv": struct{}{},
	"avc": struct{}{},
	"vp3": struct{}{},
	"svq3": struct{}{},
	"nuv": struct{}{},
	"viv": struct{}{},
	"dv": struct{}{},
	"fli": struct{}{},
	"flv": struct{}{},
	"rar": struct{}{},
	"001": struct{}{},
	"wpl": struct{}{},
	"zip": struct{}{},
}