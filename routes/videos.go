package routes

import "github.com/gorilla/mux"

func newVideosRouter(router *mux.Router) {
	/*
		/Videos
	*/
	videosRouter := router.PathPrefix("/Videos").Subrouter()

	//[Route("/Videos/{Id}/master.m3u8", "GET", Summary = "Gets a video stream using HTTP live streaming.")]
	//[Route("/Videos/{Id}/master.m3u8", "HEAD", Summary = "Gets a video stream using HTTP live streaming.")]
	//[Route("/Videos/{Id}/main.m3u8", "GET", Summary = "Gets a video stream using HTTP live streaming.")]
	//[Route("/Videos/{Id}/stream.m3u8", "GET")]
	//[Api(Description = "Gets a video stream using HTTP live streaming.")]
	//[ApiMember(Name = "BaselineStreamAudioBitRate", Description = "Optional. Specify the audio bitrate for the baseline stream.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "AppendBaselineStream", Description = "Optional. Whether or not to include a baseline audio-only stream in the master playlist.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "TimeStampOffsetMs", Description = "Optional. Alter the timestamps in the playlist by a given amount, in ms. Default is 1000.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/{stream}.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/hls/{PlaylistId}/stream.m3u8", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	videosRouter.HandleFunc("/{id}/hls/{playlistId}/{stream}.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/ActiveEncodings", "DELETE")]
	//[Api(Description = "Stops an encoding process")]
	//[ApiMember(Name = "DeviceId", Description = "The device id of the client requesting. Used to stop encoding processes when needed.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	//[ApiMember(Name = "PlaySessionId", Description = "The play session id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	videosRouter.HandleFunc("/ActiveEncodings", notYetImplemented).Methods("DELETE")

	//[Route("/Videos/{Id}/hls/{PlaylistId}/{SegmentId}.ts", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	videosRouter.HandleFunc("/{id}/hls/{playlistId}/{segmentId}.ts", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/hls1/{PlaylistId}/{SegmentId}.ts", "GET")]
	videosRouter.HandleFunc("/{id}/hls1/{playlistId}/{segmentId}.ts", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/master.mpd", "GET", Summary = "Gets a video stream using Mpeg dash.")]
	//[Route("/Videos/{Id}/master.mpd", "HEAD", Summary = "Gets a video stream using Mpeg dash.")]
	videosRouter.HandleFunc("/{id}/master.mpd", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/dash/{RepresentationId}/{SegmentId}.m4s", "GET")]
	videosRouter.HandleFunc("/{id}/dash/{repId}/{segmentId}.m4s", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/Subtitles/{Index}", "DELETE", Summary = "Deletes an external subtitle file")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Index", Description = "The subtitle stream index", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "DELETE")]
	videosRouter.HandleFunc("/{id}/Subtitles/{index}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/Stream.{Format}", "GET", Summary = "Gets subtitles in a specified format.")]
	videosRouter.HandleFunc("/{id}/{mediaSrc}/Subtitles/{index}/{streamFormat}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/{StartPositionTicks}/Stream.{Format}", "GET", Summary = "Gets subtitles in a specified format.")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "MediaSourceId", Description = "MediaSourceId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Index", Description = "The subtitle stream index", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Format", Description = "Format", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "StartPositionTicks", Description = "StartPositionTicks", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EndPositionTicks", Description = "EndPositionTicks", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/{mediaSrc}/Subtitles/{index}/{startPos}/{streamFormat}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/subtitles.m3u8", "GET", Summary = "Gets an HLS subtitle playlist.")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "MediaSourceId", Description = "MediaSourceId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Index", Description = "The subtitle stream index", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "SegmentLength", Description = "The subtitle srgment length", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/{mediaSrc}/Subtitles/{index}/subtitles.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/live.m3u8", "GET")]
	//[Api(Description = "Gets a video stream using HTTP live streaming.")]
	videosRouter.HandleFunc("/{id}/live.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/stream.mpegts", "GET")]
	//[Route("/Videos/{Id}/stream.ts", "GET")]
	//[Route("/Videos/{Id}/stream.webm", "GET")]
	//[Route("/Videos/{Id}/stream.asf", "GET")]
	//[Route("/Videos/{Id}/stream.wmv", "GET")]
	//[Route("/Videos/{Id}/stream.ogv", "GET")]
	//[Route("/Videos/{Id}/stream.mp4", "GET")]
	//[Route("/Videos/{Id}/stream.m4v", "GET")]
	//[Route("/Videos/{Id}/stream.mkv", "GET")]
	//[Route("/Videos/{Id}/stream.mpeg", "GET")]
	//[Route("/Videos/{Id}/stream.mpg", "GET")]
	//[Route("/Videos/{Id}/stream.avi", "GET")]
	//[Route("/Videos/{Id}/stream.m2ts", "GET")]
	//[Route("/Videos/{Id}/stream.3gp", "GET")]
	//[Route("/Videos/{Id}/stream.wmv", "GET")]
	//[Route("/Videos/{Id}/stream.wtv", "GET")]
	//[Route("/Videos/{Id}/stream.mov", "GET")]
	//[Route("/Videos/{Id}/stream.iso", "GET")]
	//[Route("/Videos/{Id}/stream", "GET")]
	//[Route("/Videos/{Id}/stream.ts", "HEAD")]
	//[Route("/Videos/{Id}/stream.webm", "HEAD")]
	//[Route("/Videos/{Id}/stream.asf", "HEAD")]
	//[Route("/Videos/{Id}/stream.wmv", "HEAD")]
	//[Route("/Videos/{Id}/stream.ogv", "HEAD")]
	//[Route("/Videos/{Id}/stream.mp4", "HEAD")]
	//[Route("/Videos/{Id}/stream.m4v", "HEAD")]
	//[Route("/Videos/{Id}/stream.mkv", "HEAD")]
	//[Route("/Videos/{Id}/stream.mpeg", "HEAD")]
	//[Route("/Videos/{Id}/stream.mpg", "HEAD")]
	//[Route("/Videos/{Id}/stream.avi", "HEAD")]
	//[Route("/Videos/{Id}/stream.3gp", "HEAD")]
	//[Route("/Videos/{Id}/stream.wmv", "HEAD")]
	//[Route("/Videos/{Id}/stream.wtv", "HEAD")]
	//[Route("/Videos/{Id}/stream.m2ts", "HEAD")]
	//[Route("/Videos/{Id}/stream.mov", "HEAD")]
	//[Route("/Videos/{Id}/stream.iso", "HEAD")]
	//[Route("/Videos/{Id}/stream", "HEAD")]
	//[Api(Description = "Gets a video stream")]
	videosRouter.HandleFunc("/{id}/{stream}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/AdditionalParts", "GET", Summary = "Gets additional parts for a video.")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/AdditionalParts", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/AlternateSources", "DELETE", Summary = "Removes alternate video sources.")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	videosRouter.HandleFunc("/{id}/AlternateSources", notYetImplemented).Methods("DELETE")

	//[Route("/Videos/MergeVersions", "POST", Summary = "Merges videos into a single record")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Ids", Description = "Item id list. This allows multiple, comma delimited.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	videosRouter.HandleFunc("/{id}/MergeVersions", notYetImplemented).Methods("POST")
}
