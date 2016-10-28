package routes

import "github.com/gorilla/mux"

func newConnectRouter(router *mux.Router) {
	/*
		/Connect
	*/
	connectRouter := router.PathPrefix("/Connect").Subrouter()

	//[Route("/Connect/Invite", "POST", Summary = "Creates a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "ConnectUsername", Description = "Connect username", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "SendingUserId", Description = "Sending User Id", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "EnabledLibraries", Description = "EnabledLibraries", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "EnabledChannels", Description = "EnabledChannels", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "EnableLiveTv", Description = "EnableLiveTv", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	connectRouter.HandleFunc("/Invite", notYetImplemented).Methods("POST")

	//[Route("/Connect/Pending", "GET", Summary = "Creates a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	connectRouter.HandleFunc("/Pending", notYetImplemented).Methods("GET")

	//[Route("/Connect/Pending", "DELETE", Summary = "Deletes a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Authorization Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	connectRouter.HandleFunc("/Pending", notYetImplemented).Methods("DELETE")

	//[Route("/Connect/Exchange", "GET", Summary = "Gets the corresponding local user from a connect user id")]
	//[Authenticated]
	//[ApiMember(Name = "ConnectUserId", Description = "ConnectUserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	connectRouter.HandleFunc("/Exchange", notYetImplemented).Methods("GET")
}
