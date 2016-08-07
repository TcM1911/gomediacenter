package gomediacenter

// Video is the struct used to hold video data.
type Video struct {
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

// Studio holds data of a production studio.
type Studio struct {
	Name string `json:"Name"`
	ID   string `json:"Id"`
}

// VideoSource holds information of the video source file.
type VideoSource struct {
	Protocol              string        `json:"Protocol"`
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
	RequiredHTTPHeaders   []interface{} `json:"RequiredHttpHeaders"`
	DefaultAudioStream    int           `json:"DefaultAudioStreamIndex"`
	Chapters              []Chapter     `json:"Chapters,array"`
}

// VideoStream holds information about the video stream in the file.
type VideoStream struct {
	Codec                  string  `json:"Codec"`
	Lang                   string  `json:"Language"`
	Interlaced             bool    `json:"IsInterlaced"`
	BitRate                int64   `json:"BitRate"`
	BitDepth               int     `json:"BitDepth"`
	RefFrames              int     `json:"RefFrames"`
	Default                bool    `json:"IsDefault"`
	Forced                 bool    `json:"IsForced"`
	Height                 int     `json:"Height"`
	Width                  int     `json:"Width"`
	AverageFrameRate       float64 `json:"AverageFrameRate"`
	FrameRate              float64 `json:"RealFrameRate"`
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

// Chapter holds data of each chapter.
type Chapter struct {
	StartPos int64  `json:"StartPositionTicks"`
	Name     string `json:"Name"`
}

// Intro holds data for intros to be played before a media is being played.
type Intro struct {
	Name string
	Path string
}
