package routes

import "github.com/gorilla/mux"

func newDisplayPreferenceRouter(router *mux.Router) {
	/*
		/DisplayPreferences
	*/
	displayPrefrenceRouter := router.PathPrefix("/DisplayPreferences").Subrouter()

	//[Route("/DisplayPreferences/{DisplayPreferencesId}", "POST", Summary = "Updates a user's display preferences for an item")]
	//[ApiMember(Name = "DisplayPreferencesId", Description = "DisplayPreferences Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	displayPrefrenceRouter.HandleFunc("/{id}", notYetImplemented).Methods("POST")

	//[Route("/DisplayPreferences/{Id}", "GET", Summary = "Gets a user's display preferences for an item")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Client", Description = "Client", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	displayPrefrenceRouter.HandleFunc("/{id}", notYetImplemented).Methods("GET")
}
