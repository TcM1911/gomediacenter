package routes

import "github.com/gorilla/mux"

func newPersonsRouter(router *mux.Router) {
	/*
		/Persons
	*/
	personsRouter := router.PathPrefix("/Persons").Subrouter()

	//[Route("/Persons/{Name}/Images/{Type}", "GET")]
	//[Route("/Persons/{Name}/Images/{Type}", "HEAD")]
	personsRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Persons/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Persons/{Name}/Images/{Type}/{Index}", "HEAD")]
	personsRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Persons", "GET", Summary = "Gets all persons from a given item, folder, or the entire library")]
	router.HandleFunc("/Persons", notYetImplemented).Methods("GET")

	//[Route("/Persons/{Name}", "GET", Summary = "Gets a person, by name")]
	//[ApiMember(Name = "Name", Description = "The person name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	personsRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")
}
