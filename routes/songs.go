package routes

import "github.com/gorilla/mux"

func newSongsRouter(router *mux.Router) {
	/*
		/Songs
	*/
	songsRouter := router.PathPrefix("/Songs").Subrouter()

	//[Route("/Songs/{Id}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given song")]
	songsRouter.HandleFunc("/{id}/InstantMix", notYetImplemented).Methods("GET")
}
