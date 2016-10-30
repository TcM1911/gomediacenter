package routes

import "github.com/gorilla/mux"

func newAuthRouter(router *mux.Router) {
	/*
	   /Auth
	*/
	authRouter := router.PathPrefix("/Auth").Subrouter()

	//[Route("/Auth/Keys", "GET")]
	//[Authenticated(Roles = "Admin")]
	authRouter.HandleFunc("/Keys", notYetImplemented).Methods("GET")

	//[Route("/Auth/Keys/{Key}", "DELETE")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Key", Description = "Auth Key", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	authRouter.HandleFunc("/Keys/{key}", notYetImplemented).Methods("DELETE")

	//[Route("/Auth/Keys", "POST")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "App", Description = "App", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	authRouter.HandleFunc("/Keys", notYetImplemented).Methods("POST")
}
