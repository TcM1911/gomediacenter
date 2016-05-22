package routes

import "github.com/gorilla/mux"

func newArtistsRouter(router *mux.Router) {
	/*
		/Artists router
	*/
	artistsRouter := router.PathPrefix("/Artists").Subrouter()

	//[Route("/Artists/{Id}/Similar", "GET", Summary = "Finds albums similar to a given album.")]
	artistsRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Artists/{Name}/Images/{Type}", "GET")]
	//[Route("/Artists/{Name}/Images/{Type}", "HEAD")]
	artistsRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Artists/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Artists/{Name}/Images/{Type}/{Index}", "HEAD")]
	artistsRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Artists/{Name}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given artist")]
	//[ApiMember(Name = "Name", Description = "The artist name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	artistsRouter.HandleFunc("/{name}/InstantMix", notYetImplemented).Methods("GET")

	//[Route("/Artists/InstantMix", "GET", Summary = "Creates an instant playlist based on a given artist")]
	//[ApiMember(Name = "Id", Description = "The artist Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	artistsRouter.HandleFunc("/InstantMix", notYetImplemented).Methods("GET")
}
