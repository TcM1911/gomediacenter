package routes

import "github.com/gorilla/mux"

func newAlbumsRounter(router *mux.Router) {
	/*
		/Albums router
	*/
	albumsRouter := router.PathPrefix("/Albums").Subrouter()

	//[Route("/Albums/{Id}/Similar", "GET", Summary = "Finds albums similar to a given album.")]
	albumsRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Albums/{Id}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given album")]
	albumsRouter.HandleFunc("/{id}/InstantMix", notYetImplemented).Methods("GET")
}
