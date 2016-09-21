package gomediacenter

// AudioStream struct holds all the data for a audio stream. This stream is used in both video files
// and audio files.
type AudioStream struct {
	Codec                  string `json:"Codec"`
	Lang                   string `json:"Language"`
	Interlaced             bool   `json:"IsInterlaced"`
	ChannelLayout          string `json:"ChannelLayout"`
	BitRate                int64  `json:"BitRate"`
	Channels               int    `json:"Channels"`
	SampleRate             int    `json:"SampleRate"`
	Default                bool   `json:"IsDefault"`
	Forced                 bool   `json:"IsForced"`
	Profile                string `json:"Profile"`
	Type                   string `json:"Type"`
	Index                  int    `json:"Index"`
	External               bool   `json:"IsExternal"`
	SupportsExternalStream bool   `json:"SupportsExternalStream"`
	Level                  int    `json:"Level"`
}
