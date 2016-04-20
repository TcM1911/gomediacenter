package gomediacenter

import "gopkg.in/mgo.v2/bson"

type Video struct {
	Name         string        `json:"Name"`
	SortName     string        `json:"SortName"`
	Id           bson.ObjectId `json:"Id"`
	Etage        string        `json:"Etage,omitempty"`
	CanDelete    bool          `json:"CanDelete"`
	CanDownload  bool          `json:"CanDownload"`
	MediaSources []interface{} `json:"MediaSources,array"`
	LocationType string        `json:"LocationType"`
	MediaType    string        `json:"MediaType"`
	HomePage     string        `json:"HomePageUrl"`
	LockedFields []string      `json:"LockedFields"`
	LockedData   bool          `json:"LockData"`
}

type Studio struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

type VideoSource struct {
	Protocol              string        `json:"Protocol"`
	Path                  string        `json:"Path"`
	Type                  string        `json:"Type"`
	Container             string        `json:"Container"`
	Name                  string        `json:"Name"`
	RuntimeTicks          int           `json:"RunTimeTicks"`
	ReadAtNativeFramerate bool          `json:"ReadAtNativeFramerate"`
	SupportTranscodeing   bool          `json:"SupportsTranscoding"`
	SupportDirectStream   bool          `json:"SupportsDirectStream"`
	SupportDirectPlay     bool          `json:"SupportsDirectPlay"`
	RequireOpening        bool          `json:"RequiresOpening"`
	RequiresClosing       bool          `json:"RequiresClosing"`
	VideoType             string        `json:"VideoType"`
	MediaStreams          []interface{} `json:"MediaStreams"`
	StreamFileNames       []string      `json:"PlayableStreamFileNames"`
	Formats               []string      `json:"Formats"`
	RequiredHttpHeaders   []interface{} `json:"RequiredHttpHeaders"`
	DefaultAudioStream    int           `json:"DefaultAudioStreamIndex"`
	Chapters              []Chapter     `json:"Chapters,array"`
}

type VideoStream struct {
	Codec                  string  `json:"Codec"`
	Lang                   string  `json:"Language"`
	Interlaced             bool    `json:"IsInterlaced"`
	BitRate                int     `json:"BitRate"`
	BitDepth               int     `json:"BitDepth"`
	RefFrames              int     `json:"RefFrames"`
	Default                bool    `json:"IsDefault"`
	Forced                 bool    `json:"IsForced"`
	Height                 int     `json:"Height"`
	Width                  int     `json:"Width"`
	AverageFrameRate       float32 `json:"AverageFrameRate"`
	FrameRate              float32 `json:"RealFrameRate"`
	Profile                string  `json:"Profile"`
	Type                   string  `json:"Type"`
	AspectRatio            string  `json:"AspectRatio"`
	Index                  int     `json:"Index"`
	External               bool    `json:"IsExternal"`
	TextSubtitleStream     bool    `json:"IsTextSubtitleStream"`
	SupportsExternalStream bool    `json:"SupportsExternalStream"`
	PixelFormat            string  `json:"PixelFormat"`
	Level                  int     `json:"Level"`
	Anamorphic             bool    `json:"IsAnamorphic"`
	Cabac                  bool    `json:"IsCabac"`
}

type Chapter struct {
	StartPos int    `json:"StartPositionTicks"`
	Name     string `json:"Name"`
}

type Intro struct {
	Name string
	Path string
}
