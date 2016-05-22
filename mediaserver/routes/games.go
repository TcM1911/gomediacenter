package routes

import "github.com/gorilla/mux"

func newGamesRouter(router *mux.Router) {
	/*
		/Games
	*/
	gamesRouter := router.PathPrefix("/Games").Subrouter()

	//[Route("/Games/{Id}/Similar", "GET", Summary = "Finds games similar to a given game.")]
	gamesRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Games/SystemSummaries", "GET", Summary = "Finds games similar to a given game.")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	gamesRouter.HandleFunc("/SystemSummaries", notYetImplemented).Methods("GET")

	//[Route("/Games/PlayerIndex", "GET", Summary = "Gets an index of players (1-x) and the number of games listed under each")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	gamesRouter.HandleFunc("/PlayerIndex", notYetImplemented).Methods("GET")
}
