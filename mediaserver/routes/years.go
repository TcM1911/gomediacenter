package routes

import "github.com/gorilla/mux"

func newYearsRouter(router *mux.Router) {
	/*
		/Years
	*/
	yearsRouter := router.PathPrefix("/Years").Subrouter()

	//[Route("/Years", "GET", Summary = "Gets all years from a given item, folder, or the entire library")]
	router.HandleFunc("/Years", notYetImplemented).Methods("GET")

	//[Route("/Years/{Year}", "GET", Summary = "Gets a year")]
	//[ApiMember(Name = "Year", Description = "The year", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	yearsRouter.HandleFunc("/{year}", notYetImplemented).Methods("GET")

	//[Route("/Years/{Year}/Images/{Type}", "GET")]
	//[Route("/Years/{Year}/Images/{Type}", "HEAD")]
	yearsRouter.HandleFunc("/{year}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Years/{Year}/Images/{Type}/{Index}", "GET")]
	//[Route("/Years/{Year}/Images/{Type}/{Index}", "HEAD")]
	//[ApiMember(Name = "Name", Description = "Item name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	yearsRouter.HandleFunc("/{year}/Images/{type}/{index}", notYetImplemented).Methods("GET")
}
