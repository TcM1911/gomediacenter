package gomediacenter

type Video struct {
	Name         string `json:"Name"`
	sortName     string `json:"SortName"`
	Id           string `json:"Id"`
	Etage        string `json:"Etage,omitempty"`
	CanDelete    bool `json:"CanDelete"`
	CanDownload  bool `json:"CanDownload"`
	Path         string
	RuntimeTicks int `json:"RunTimeTicks"`
}

type Studio struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

type VideoSource struct {
	Protocol              string `json:"Protocol"`
	Id                    string `json:"Id"`
	Path                  string `json:"Path"`
	Type                  string `json:"Type"`
	Container             string `json:"Container"`
	Name                  string `json:"Name"`
	RuntimeTicks          int `json:"RunTimeTicks"`
	ReadAtNativeFramerate bool `json:"ReadAtNativeFramerate"`
	SupportTranscodeing   bool `json:"SupportsTranscoding"`
	SupportDirectStream   bool `json:"SupportsDirectStream"`
	SupportDirectPlay     bool `json:"SupportsDirectPlay"`
	RequireOpening        bool `json:"RequiresOpening"`
	RequiresClosing       bool `json:"RequiresClosing"`
	VideoType             string `json:"VideoType"`
	MediaStreams          []interface{} `json:"MediaStreams"`
}

type VideoStream struct {
	Codec                  string `json:"Codec"`
	Lang                   string `json:"Language"`
	Interlaced             bool `json:"IsInterlaced"`
	BitRate                int `json:"BitRate"`
	BitDepth               int `json:"BitDepth"`
	RefFrames              int `json:"RefFrames"`
	Default                bool `json:"IsDefault"`
	Forced                 bool `json:"IsForced"`
	Height                 int `json:"Height"`
	Width                  int `json:"Width"`
	AverageFrameRate       float32 `json:"AverageFrameRate"`
	FrameRate              float32 `json:"RealFrameRate"`
	Profile                string `json:"Profile"`
	Type                   string `json:"Type"`
	AspectRatio            string `json:"AspectRatio"`
	Index                  int `json:"Index"`
	External               bool `json:"IsExternal"`
	TextSubtitleStream     bool `json:"IsTextSubtitleStream"`
	SupportsExternalStream bool `json:"SupportsExternalStream"`
	PixelFormat            string `json:"PixelFormat"`
	Level                  int `json:"Level"`
	Anamorphic             bool `json:"IsAnamorphic"`
	Cabac                  bool `json:"IsCabac"`
}
