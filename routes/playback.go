package routes

import "github.com/gorilla/mux"

func newPlaybackRouter(router *mux.Router) {
	/*
		/Playback
	*/
	playbackRouter := router.PathPrefix("/Playback").Subrouter()

	//[Route("/Playback/BitrateTest", "GET")]
	//[ApiMember(Name = "Size", Description = "Size", IsRequired = true, DataType = "int", ParameterType = "query", Verb = "GET")]
	playbackRouter.HandleFunc("/BitrateTest", notYetImplemented).Methods("GET")
}
