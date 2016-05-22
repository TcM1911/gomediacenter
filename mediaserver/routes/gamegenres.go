package routes

import "github.com/gorilla/mux"

func newGameGenresRouter(router *mux.Router) {
	/*
		/GameGenres
	*/
	gameGenresRouter := router.PathPrefix("/GameGenres").Subrouter()

	//[Route("/GameGenres", "GET", Summary = "Gets all Game genres from a given item, folder, or the entire library")]
	router.HandleFunc("/GameGenres", notYetImplemented).Methods("GET")

	//[Route("/GameGenres/{Name}", "GET", Summary = "Gets a Game genre, by name")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	gameGenresRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")

	//[Route("/GameGenres/{Name}/Images/{Type}", "GET")]
	//[Route("/GameGenres/{Name}/Images/{Type}", "HEAD")]
	gameGenresRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/GameGenres/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/GameGenres/{Name}/Images/{Type}/{Index}", "HEAD")]
	gameGenresRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")
}
