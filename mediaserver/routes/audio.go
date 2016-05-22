package routes

import "github.com/gorilla/mux"

func newAudioRouter(router *mux.Router) {
	/*
		/Audio
	*/
	audioRouter := router.PathPrefix("/Audio").Subrouter()

	//[Route("/Audio/{Id}/stream.{Container}", "GET", Summary = "Gets an audio stream")]
	//[Route("/Audio/{Id}/stream", "GET", Summary = "Gets an audio stream")]
	//[Route("/Audio/{Id}/stream.{Container}", "HEAD", Summary = "Gets an audio stream")]
	//[Route("/Audio/{Id}/stream", "HEAD", Summary = "Gets an audio stream")]
	//[ApiMember(Name = "Container", Description = "Container", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	audioRouter.HandleFunc("/{id}/stream", notYetImplemented).Methods("GET")

	//[Route("/Audio/{Id}/master.m3u8", "GET", Summary = "Gets an audio stream using HTTP live streaming.")]
	//[Route("/Audio/{Id}/master.m3u8", "HEAD", Summary = "Gets an audio stream using HTTP live streaming.")]
	//[Route("/Audio/{Id}/main.m3u8", "GET", Summary = "Gets an audio stream using HTTP live streaming.")]
	audioRouter.HandleFunc("/{id}/{stream}.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Audio/{Id}/hls/{SegmentId}/stream.mp3", "GET")]
	//[Route("/Audio/{Id}/hls/{SegmentId}/stream.aac", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	audioRouter.HandleFunc("/{id}/hls/{segmentId}/{stream}", notYetImplemented).Methods("GET")

	//[Route("/Audio/{Id}/hls1/{PlaylistId}/{SegmentId}.aac", "GET")]
	//[Route("/Audio/{Id}/hls1/{PlaylistId}/{SegmentId}.ts", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	audioRouter.HandleFunc("/{id}/hls1/{playlistId}/{segmentId}", notYetImplemented).Methods("GET")

}
