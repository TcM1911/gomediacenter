package routes

import "github.com/gorilla/mux"

func newLiveStreamsRouter(router *mux.Router) {
	/*
		/LiveStreams
	*/
	liveStreamsRoute := router.PathPrefix("/LiveStreams").Subrouter()

	//[Route("/LiveStreams/Open", "POST", Summary = "Opens a media source")]
	liveStreamsRoute.HandleFunc("/Open", notYetImplemented).Methods("POST")

	//[Route("/LiveStreams/Close", "POST", Summary = "Closes a media source")]
	//[ApiMember(Name = "LiveStreamId", Description = "LiveStreamId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	liveStreamsRoute.HandleFunc("/Close", notYetImplemented).Methods("POST")
}
