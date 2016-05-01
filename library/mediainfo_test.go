package library

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

///////////
// Tests //
///////////

func TestParseFFprobeOutputOfMKVFile(t *testing.T) {
	assert := assert.New(t)

	parsedOutput, err := ParseFFprobeOutput(ffprobeOutputMkvFile)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Format
	assert.Equal("/mnt/Videos/MKV.Movie.File.mkv", parsedOutput.Format.File)
	assert.Equal(2, parsedOutput.Format.StreamCount)
	assert.Equal("matroska,webm", parsedOutput.Format.FormatName)
	assert.Equal("Matroska/WebM", parsedOutput.Format.FormatLongName)
	assert.Equal("0.000000", parsedOutput.Format.StartTime)
	assert.Equal("6954.410000", parsedOutput.Format.Duration)
	assert.Equal("806213424", parsedOutput.Format.Size)
	assert.Equal("927426", parsedOutput.Format.BitRate)
	assert.Equal(100, parsedOutput.Format.ProbeScore)
	assert.Equal("libebmlv1.3.0+libmatroskav1.4.0", parsedOutput.Format.Tags.Encoder)

	// Chapters
	assert.Equal(20, len(parsedOutput.Chapters))
	assert.Equal(189336231, parsedOutput.Chapters[0].Id)
	assert.Equal(int64(182140000000), parsedOutput.Chapters[1].StartPos)
	assert.Equal("1/1000000000", parsedOutput.Chapters[1].TimeBase)

	// Streams
	assert.Equal(2, len(parsedOutput.Streams))
	assert.Equal("video", parsedOutput.Streams[0].Type)
	assert.Equal("audio", parsedOutput.Streams[1].Type)

	videoStream := parsedOutput.Streams[0]
	audioStream := parsedOutput.Streams[1]

	// Video stream
	assert.Equal(0, videoStream.Index)
	assert.Equal("h264", videoStream.CodecName)
	assert.Equal("High", videoStream.Profile)
	assert.Equal(720, videoStream.Width)
	assert.Equal(300, videoStream.Height)
	assert.Equal(2, videoStream.HasBFrames)
	assert.Equal("1:1", videoStream.SampleAspectRatio)
	assert.Equal("12:5", videoStream.DisplayAspectRatio)
	assert.Equal("yuv420p", videoStream.PixFmt)
	assert.Equal(31, videoStream.Level)
	assert.Equal("left", videoStream.ChromaLocation)
	assert.Equal(5, videoStream.Refs)
	assert.Equal("true", videoStream.Avc)
	assert.Equal("13978/583", videoStream.FrameRate)
	assert.Equal("13978/583", videoStream.AvgFrameRate)
	assert.Equal("1/1000", videoStream.TimeBase)
	assert.Equal(0, videoStream.StartPts)
	assert.Equal("0.000000", videoStream.StartTime)

	// Audio stream
	assert.Equal(1, audioStream.Index)
	assert.Equal("aac", audioStream.CodecName)
	assert.Equal("LC", audioStream.Profile)
	assert.Equal("audio", audioStream.Type)
	assert.Equal("fltp", audioStream.SampleFmt)
	assert.Equal("48000", audioStream.SampleRate)
	assert.Equal(2, audioStream.Channels)
	assert.Equal("stereo", audioStream.ChannelLayout)
	assert.Equal("1/1000", audioStream.TimeBase)
}

func TestParseFFprobeOutputOfAviFile(t *testing.T) {
	assert := assert.New(t)

	parsedOutput, err := ParseFFprobeOutput(ffprobeOutputAviFile)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Format
	assert.Equal("/mnt/Videos/AVI.Movie.File.avi", parsedOutput.Format.File)
	assert.Equal(2, parsedOutput.Format.StreamCount)
	assert.Equal("avi", parsedOutput.Format.FormatName)
	assert.Equal("AVI(AudioVideoInterleaved)", parsedOutput.Format.FormatLongName)
	assert.Equal("0.000000", parsedOutput.Format.StartTime)
	assert.Equal("8699.000000", parsedOutput.Format.Duration)
	assert.Equal("1464607688", parsedOutput.Format.Size)
	assert.Equal("1346920", parsedOutput.Format.BitRate)
	assert.Equal(100, parsedOutput.Format.ProbeScore)

	// Chapters
	assert.Equal(0, len(parsedOutput.Chapters))

	// Streams
	assert.Equal(2, len(parsedOutput.Streams))
	assert.Equal("video", parsedOutput.Streams[0].Type)
	assert.Equal("audio", parsedOutput.Streams[1].Type)

	videoStream := parsedOutput.Streams[0]
	audioStream := parsedOutput.Streams[1]

	// Video stream
	assert.Equal(0, videoStream.Index)
	assert.Equal("mpeg4", videoStream.CodecName)
	assert.Equal("SimpleProfile", videoStream.Profile)
	assert.Equal(608, videoStream.Width)
	assert.Equal(272, videoStream.Height)
	assert.Equal(0, videoStream.HasBFrames)
	assert.Equal("1:1", videoStream.SampleAspectRatio)
	assert.Equal("38:17", videoStream.DisplayAspectRatio)
	assert.Equal("yuv420p", videoStream.PixFmt)
	assert.Equal(0, videoStream.Level)
	assert.Equal("left", videoStream.ChromaLocation)
	assert.Equal(1, videoStream.Refs)
	assert.Equal("false", videoStream.DivxPacked)
	assert.Equal("false", videoStream.QuarterSample)
	assert.Equal("25/1", videoStream.FrameRate)
	assert.Equal("25/1", videoStream.AvgFrameRate)
	assert.Equal("1/25", videoStream.TimeBase)
	assert.Equal(0, videoStream.StartPts)
	assert.Equal("0.000000", videoStream.StartTime)

	// Audio stream
	assert.Equal(1, audioStream.Index)
	assert.Equal("mp3", audioStream.CodecName)
	assert.Equal("audio", audioStream.Type)
	assert.Equal("s16p", audioStream.SampleFmt)
	assert.Equal("48000", audioStream.SampleRate)
	assert.Equal(2, audioStream.Channels)
	assert.Equal("stereo", audioStream.ChannelLayout)
	assert.Equal("3/125", audioStream.TimeBase)
}

///////////////////////
// Private variables //
///////////////////////

var ffprobeOutputAviFile []byte = []byte(`{"streams":[{"index":0,"codec_name":"mpeg4","codec_long_name":"MPEG-4part2","profile":"SimpleProfile","codec_type":"video","codec_time_base":"1/25","codec_tag_string":"XVID","codec_tag":"0x44495658","width":608,"height":272,"coded_width":608,"coded_height":272,"has_b_frames":0,"sample_aspect_ratio":"1:1","display_aspect_ratio":"38:17","pix_fmt":"yuv420p","level":0,"chroma_location":"left","refs":1,"quarter_sample":"false","divx_packed":"false","r_frame_rate":"25/1","avg_frame_rate":"25/1","time_base":"1/25","start_pts":0,"start_time":"0.000000","duration_ts":217475,"duration":"8699.000000","bit_rate":"1213210","nb_frames":"217475","disposition":{"default":0,"dub":0,"original":0,"comment":0,"lyrics":0,"karaoke":0,"forced":0,"hearing_impaired":0,"visual_impaired":0,"clean_effects":0,"attached_pic":0}},{"index":1,"codec_name":"mp3","codec_long_name":"MP3(MPEGaudiolayer3)","codec_type":"audio","codec_time_base":"1/48000","codec_tag_string":"U[0][0][0]","codec_tag":"0x0055","sample_fmt":"s16p","sample_rate":"48000","channels":2,"channel_layout":"stereo","bits_per_sample":0,"r_frame_rate":"0/0","avg_frame_rate":"0/0","time_base":"3/125","start_pts":0,"start_time":"0.000000","duration_ts":357431,"duration":"8578.344000","bit_rate":"122376","nb_frames":"357431","disposition":{"default":0,"dub":0,"original":0,"comment":0,"lyrics":0,"karaoke":0,"forced":0,"hearing_impaired":0,"visual_impaired":0,"clean_effects":0,"attached_pic":0}}],"chapters":[],"format":{"filename":"/mnt/Videos/AVI.Movie.File.avi","nb_streams":2,"nb_programs":0,"format_name":"avi","format_long_name":"AVI(AudioVideoInterleaved)","start_time":"0.000000","duration":"8699.000000","size":"1464607688","bit_rate":"1346920","probe_score":100,"tags":{"encoder":"MEncoderSVN-r29237-4.4.1"}}}`)

var ffprobeOutputMkvFile []byte = []byte(`{"streams":[{"index":0,"codec_name":"h264","codec_long_name":"H.264/AVC/MPEG-4AVC/MPEG-4part10","profile":"High","codec_type":"video","codec_time_base":"104271/5000000","codec_tag_string":"[0][0][0][0]","codec_tag":"0x0000","width":720,"height":300,"coded_width":720,"coded_height":304,"has_b_frames":2,"sample_aspect_ratio":"1:1","display_aspect_ratio":"12:5","pix_fmt":"yuv420p","level":31,"color_range":"tv","color_space":"bt709","chroma_location":"left","refs":5,"is_avc":"true","nal_length_size":"4","r_frame_rate":"13978/583","avg_frame_rate":"13978/583","time_base":"1/1000","start_pts":0,"start_time":"0.000000","bits_per_raw_sample":"8","disposition":{"default":1,"dub":0,"original":0,"comment":0,"lyrics":0,"karaoke":0,"forced":0,"hearing_impaired":0,"visual_impaired":0,"clean_effects":0,"attached_pic":0},"tags":{"language":"eng","title":"X264"}},{"index":1,"codec_name":"aac","codec_long_name":"AAC(AdvancedAudioCoding)","profile":"LC","codec_type":"audio","codec_time_base":"1/48000","codec_tag_string":"[0][0][0][0]","codec_tag":"0x0000","sample_fmt":"fltp","sample_rate":"48000","channels":2,"channel_layout":"stereo","bits_per_sample":0,"r_frame_rate":"0/0","avg_frame_rate":"0/0","time_base":"1/1000","start_pts":0,"start_time":"0.000000","disposition":{"default":1,"dub":0,"original":0,"comment":0,"lyrics":0,"karaoke":0,"forced":0,"hearing_impaired":0,"visual_impaired":0,"clean_effects":0,"attached_pic":0},"tags":{"language":"eng"}}],"chapters":[{"id":189336231,"time_base":"1/1000000000","start":0,"start_time":"0.000000","end":182140000000,"end_time":"182.140000","tags":{"title":"00:00:00.000"}},{"id":-2105106549,"time_base":"1/1000000000","start":182140000000,"start_time":"182.140000","end":505797000000,"end_time":"505.797000","tags":{"title":"00:03:02.140"}},{"id":2005430194,"time_base":"1/1000000000","start":505797000000,"start_time":"505.797000","end":746996000000,"end_time":"746.996000","tags":{"title":"00:08:25.797"}},{"id":-1787971142,"time_base":"1/1000000000","start":746996000000,"start_time":"746.996000","end":1104895000000,"end_time":"1104.895000","tags":{"title":"00:12:26.996"}},{"id":-1709341564,"time_base":"1/1000000000","start":1104895000000,"start_time":"1104.895000","end":1386176000000,"end_time":"1386.176000","tags":{"title":"00:18:24.895"}},{"id":2098974878,"time_base":"1/1000000000","start":1386176000000,"start_time":"1386.176000","end":1693442000000,"end_time":"1693.442000","tags":{"title":"00:23:06.176"}},{"id":586930054,"time_base":"1/1000000000","start":1693442000000,"start_time":"1693.442000","end":2110442000000,"end_time":"2110.442000","tags":{"title":"00:28:13.442"}},{"id":764065941,"time_base":"1/1000000000","start":2110442000000,"start_time":"2110.442000","end":2373705000000,"end_time":"2373.705000","tags":{"title":"00:35:10.442"}},{"id":1700123307,"time_base":"1/1000000000","start":2373705000000,"start_time":"2373.705000","end":2699822000000,"end_time":"2699.822000","tags":{"title":"00:39:33.705"}},{"id":-1513455956,"time_base":"1/1000000000","start":2699822000000,"start_time":"2699.822000","end":3014470000000,"end_time":"3014.470000","tags":{"title":"00:44:59.822"}},{"id":-1237939300,"time_base":"1/1000000000","start":3014470000000,"start_time":"3014.470000","end":3494533000000,"end_time":"3494.533000","tags":{"title":"00:50:14.470"}},{"id":989510569,"time_base":"1/1000000000","start":3494533000000,"start_time":"3494.533000","end":3944941000000,"end_time":"3944.941000","tags":{"title":"00:58:14.533"}},{"id":-633358443,"time_base":"1/1000000000","start":3944941000000,"start_time":"3944.941000","end":4408070000000,"end_time":"4408.070000","tags":{"title":"01:05:44.941"}},{"id":-2118768510,"time_base":"1/1000000000","start":4408070000000,"start_time":"4408.070000","end":4675129000000,"end_time":"4675.129000","tags":{"title":"01:13:28.070"}},{"id":-914335315,"time_base":"1/1000000000","start":4675129000000,"start_time":"4675.129000","end":5024561000000,"end_time":"5024.561000","tags":{"title":"01:17:55.129"}},{"id":-660535148,"time_base":"1/1000000000","start":5024561000000,"start_time":"5024.561000","end":5244656000000,"end_time":"5244.656000","tags":{"title":"01:23:44.561"}},{"id":653055159,"time_base":"1/1000000000","start":5244656000000,"start_time":"5244.656000","end":5727889000000,"end_time":"5727.889000","tags":{"title":"01:27:24.656"}},{"id":-550277732,"time_base":"1/1000000000","start":5727889000000,"start_time":"5727.889000","end":6172374000000,"end_time":"6172.374000","tags":{"title":"01:35:27.889"}},{"id":-547912316,"time_base":"1/1000000000","start":6172374000000,"start_time":"6172.374000","end":6582951000000,"end_time":"6582.951000","tags":{"title":"01:42:52.374"}},{"id":1506868920,"time_base":"1/1000000000","start":6582951000000,"start_time":"6582.951000","end":6954410000000,"end_time":"6954.410000","tags":{"title":"01:49:42.951"}}],"format":{"filename":"/mnt/Videos/MKV.Movie.File.mkv","nb_streams":2,"nb_programs":0,"format_name":"matroska,webm","format_long_name":"Matroska/WebM","start_time":"0.000000","duration":"6954.410000","size":"806213424","bit_rate":"927426","probe_score":100,"tags":{"encoder":"libebmlv1.3.0+libmatroskav1.4.0"}}}`)
