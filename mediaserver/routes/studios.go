package routes

import "github.com/gorilla/mux"

func newStudiosRouter(router *mux.Router) {
	/*
		/Studios
	*/
	studiosRouter := router.PathPrefix("/Studios").Subrouter()

	//[Route("/Studios/{Name}/Images/{Type}", "GET")]
	//[Route("/Studios/{Name}/Images/{Type}", "HEAD")]
	studiosRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Studios/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Studios/{Name}/Images/{Type}/{Index}", "HEAD")]
	studiosRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Studios", "GET", Summary = "Gets all studios from a given item, folder, or the entire library")]
	router.HandleFunc("/Studios", notYetImplemented).Methods("GET")

	//[Route("/Studios/{Name}", "GET", Summary = "Gets a studio, by name")]
	//[ApiMember(Name = "Name", Description = "The studio name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	studiosRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")
}
