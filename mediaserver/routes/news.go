package routes

import "github.com/gorilla/mux"

func newNewsRouter(router *mux.Router) {
	/*
		/News
	*/
	newsRouter := router.PathPrefix("/News").Subrouter()

	//[Route("/News/Product", "GET", Summary = "Gets the latest product news.")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	newsRouter.HandleFunc("/Product", notYetImplemented).Methods("GET")
}
