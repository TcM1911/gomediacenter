package routes

import "github.com/gorilla/mux"

func newGenresRouter(router *mux.Router) {
	/*
		/Genres
	*/
	genresRouter := router.PathPrefix("/Genres").Subrouter()

	//[Route("/Genres", "GET", Summary = "Gets all genres from a given item, folder, or the entire library")]
	router.HandleFunc("/Genres", notYetImplemented).Methods("GET")

	//[Route("/Genres/{Name}", "GET", Summary = "Gets a genre, by name")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	genresRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")

	//[Route("/Genres/{Name}/Images/{Type}", "GET")]
	//[Route("/Genres/{Name}/Images/{Type}", "HEAD")]
	genresRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Genres/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Genres/{Name}/Images/{Type}/{Index}", "HEAD")]
	genresRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")
}
