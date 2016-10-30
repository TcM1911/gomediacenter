package routes

import "github.com/gorilla/mux"

func newTrailersRouter(router *mux.Router) {
	/*
		/Trailers
	*/
	trailerRouter := router.PathPrefix("/Trailers").Subrouter()

	//[Route("/Trailers/{Id}/Similar", "GET", Summary = "Finds movies and trailers similar to a given trailer.")]
	trailerRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Trailers", "GET", Summary = "Finds movies and trailers similar to a given trailer.")]
	router.HandleFunc("/Trailers", notYetImplemented).Methods("GET")
}
