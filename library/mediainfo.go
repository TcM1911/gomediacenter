package library

import (
	"encoding/json"
)

// ffprobe cmd: ffprobe -v quiet -show_chapters -print_format json -show_format -show_streams file.mkv

/////////////
// Structs //
/////////////

type FFprobeOutput struct {
	Format   Format `json:"format"`
	Chapters []Chapter `json:"chapters,array"`
	Streams  []Stream `json:"streams,array"`
}

type Stream struct {
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
	StartTime          string `json:"start_time"`
	DurationTS         int `json:"duration_ts"`
	DurationInSeconds  string `json:"duration"`
	BitRate            string `json:"bit_rate"`
	NBFrames           string `json:"nb_frames"`
	Disposition        Disposition `json:"disposition"`
	Tags `json:"tags,array"`

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
	QuarterSample      string `json:"quarter_sample"`
	DivxPacked         string `json:"divx_packed"`
	Avc                string `json:"is_avc"`

	// Audio
	SampleFmt          string `json:"sample_fmt"`
	SampleRate         string `json:"sample_rate"`
	Channels           int `json:"channels"`
	ChannelLayout      string `json:"channel_layout"`
	BitsPerSample      int `json:"bits_per_sample"`
}

type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
}

type Chapter struct {
	Id        int `json:"id"`
	TimeBase  string `json:"time_base"`
	StartPos  int64 `json:"start"`
	StartTime string `json:"start_time"`
	EndPos    int64 `json:"end"`
	EndTime   string `json:"end_time"`
	Tags `json:"tags"`
}

type Format struct {
	File           string `json:"filename"`
	StreamCount    int `json:"nb_streams"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      string `json:"start_time"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int `json:"probe_score"`
	Tags `json:"tags"`
}

type Tags struct {
	Title        string `json:"title"`
	Lang         string `json:"language"`
	Encoder      string `json:"encoder"`
	CreationTime string `json:"creation_time"`
}

////////////
// Public //
////////////

// ParseFFprobeOutput parses the FFprobe output and returns a FFprobeOutput struct.
func ParseFFprobeOutput(stdout []byte) (FFprobeOutput, error) {
	var parsedData FFprobeOutput
	err := json.Unmarshal(stdout, &parsedData)
	if err != nil {
		return parsedData, err
	}
	return parsedData, nil
}