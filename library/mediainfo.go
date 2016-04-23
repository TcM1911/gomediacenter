package library

// ffprobe cmd: ffprobe -v quiet -show_chapters -print_format json -show_format -show_streams file.mkv

/////////////
// Structs //
/////////////

type FFprobeOutput struct {
	Format   format `json:"format"`
	Chapters []chapter `json:"chapters,array"`
	Streams  []stream `json:"streams,array"`
}

type stream struct {
	// Video and audio
	Index              int `json:"index"`
	CodecName          string `json:"codec_name"`
	CodecLongName      string `json:"codec_long_name"`
	Profile            string `json:"profile"`
	Type               string `json:"codec_type"`
	CodecTimeBase      string `json:"codec_time_base"`
	TagString          string `json:"codec_tag_string"`
	Tag                string `json:"codec_tag"`
	FrameRate          string `json:"r_frame_rate"`
	AvgFrameRate       string `json:"avg_frame_rate"`
	TimeBase           string `json:"time_base"`
	StartPts           int `json:"start_pts"`
	StartTime          float32 `json:"start_time"`
	DurationTS         int `json:"duration_ts"`
	DurationInSeconds  int64 `json:"duration"`
	BitRate            int `json:"bit_rate"`
	NBFrames           int `json:"nb_frames"`
	Disposition        disposition `json:"disposition"`
	Tags               tags `json:"tags"`

	// Video
	Width              int `json:"width"`
	Height             int `json:"height"`
	HasBFrames         int `json:"has_b_frames"`
	SampleAspectRatio  string `json:"sample_aspect_ratio"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	PixFmt             string `json:"pix_fmt"`
	Level              int `json:"level"`
	ChromaLocation     string `json:"chroma_location"`
	Refs               int `json:"refs"`
	QuarterSample      bool `json:"quarter_sample"`
	DivxPacked         bool `json:"divx_packed"`

	// Audio
	SampleFmt          string `json:"sample_fmt"`
	SampleRate         int `json:"sample_rate"`
	Channels           int `json:"channels"`
	ChannelLayout      string `json:"channel_layout"`
	BitsPerSample      int `json:"bits_per_sample"`
}

type disposition struct {
	Default         bool `json:"default"`
	Dub             bool `json:"dub"`
	Original        bool `json:"original"`
	Comment         bool `json:"comment"`
	Lyrics          bool `json:"lyrics"`
	Karaoke         bool `json:"karaoke"`
	Forced          bool `json:"forced"`
	HearingImpaired bool `json:"hearing_impaired"`
	VisualImpaired  bool `json:"visual_impaired"`
	CleanEffects    bool `json:"clean_effects"`
	AttachedPic     bool `json:"attached_pic"`
}

type chapter struct {
	Id        int `json:"id"`
	TimeBase  string `json:"time_base"`
	StartPos  int64 `json:"start"`
	StartTime string `json:"start_time"`
	EndPos    int64 `json:"end"`
	EndTime   string `json:"end_time"`
	Tags      tags `json:"tags"`
}

type format struct {
	File           string `json:"filename"`
	StreamCount    int `json:"nb_streams"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      float64 `json:"start_time"`
	Duration       float64 `json:"duration"`
	Size           int `json:"size"`
	BitRate        int `json:"bit_rate"`
	ProbeScore     int `json:"probe_score"`
	Tags           tags `json:"tags"`
}

type tags struct {
	Title        string `json:"title"`
	Lang         string `json:"language"`
	Encoder      string `json:"encoder"`
	CreationTime string `json:"creation_time"`
}