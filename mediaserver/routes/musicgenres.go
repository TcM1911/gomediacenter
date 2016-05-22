package routes

import "github.com/gorilla/mux"

func newMusicGenresRouter(router *mux.Router) {
	/*
		/MusicGenres
	*/
	musicGenresRouter := router.PathPrefix("/MusicGenres").Subrouter()

	//[Route("/MusicGenres/{Name}/Images/{Type}", "GET")]
	//[Route("/MusicGenres/{Name}/Images/{Type}", "HEAD")]
	musicGenresRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/MusicGenres/{Name}/Images/{Type}/{Index}", "HEAD")]
	musicGenresRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres/{Name}/InstantMix", "GET", Summary = "Creates an instant playlist based on a music genre")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[Route("/MusicGenres/InstantMix", "GET", Summary = "Creates an instant playlist based on a music genre")]
	//[ApiMember(Name = "Id", Description = "The genre Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	musicGenresRouter.HandleFunc("/{name}/InstantMix", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres", "GET", Summary = "Gets all music genres from a given item, folder, or the entire library")]
	router.HandleFunc("/MusicGenres", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres/{Name}", "GET", Summary = "Gets a music genre, by name")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	musicGenresRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")
}
