package library

import "strings"

// IsVideoFile determines if a file is a video file based on the file extension.
// The function will only return true if the file extension matches predetermined
// file extensions.
func IsVideoFile(fileName string) bool {
	fileNameSlice := strings.Split(fileName, ".")
	ext := fileNameSlice[len(fileNameSlice) - 1]

	_, ok := videoFileExtensions[ext]
	return ok
}

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