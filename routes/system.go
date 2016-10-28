package routes

import "github.com/gorilla/mux"

func newSystemRouter(router *mux.Router) {
	/*
		System
	*/
	systemRouter := router.PathPrefix("/System").Subrouter()

	//[Route("/System/ActivityLog/Entries", "GET", Summary = "Gets activity log entries")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "MinDate", Description = "Optional. The minimum date. Format = ISO", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	systemRouter.HandleFunc("/ActivityLog/Entries", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration", "GET", Summary = "Gets application configuration")]
	//[Authenticated]
	systemRouter.HandleFunc("/Configuration", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration", "POST", Summary = "Updates application configuration")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Configuration", notYetImplemented).Methods("POST")

	//[Route("/System/Configuration/{Key}", "GET", Summary = "Gets a named configuration")]
	//[Authenticated(AllowBeforeStartupWizard = true)]
	//[ApiMember(Name = "Key", Description = "Key", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	systemRouter.HandleFunc("/Configuration/{key}", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration/{Key}", "POST", Summary = "Updates named configuration")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Key", Description = "Key", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	systemRouter.HandleFunc("/Configuration/{key}", notYetImplemented).Methods("POST")

	//[Route("/System/Configuration/MetadataOptions/Default", "GET", Summary = "Gets a default MetadataOptions object")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Configuration/MetadataOptions/Default", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration/MetadataPlugins", "GET", Summary = "Gets all available metadata plugins")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Configuration/MetadataPlugins", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration/MetadataPlugins/Autoset", "POST")]
	//[Authenticated(Roles = "Admin", AllowBeforeStartupWizard = true)]
	systemRouter.HandleFunc("/Configuration/MetadataPlugins/Autoset", notYetImplemented).Methods("POST")

	//[Route("/System/Info", "GET", Summary = "Gets information about the server")]
	//[Authenticated(EscapeParentalControl = true)]
	systemRouter.HandleFunc("/Info", notYetImplemented).Methods("GET")

	//[Route("/System/Info/Public", "GET", Summary = "Gets public information about the server")]
	systemRouter.HandleFunc("/Info/Public", notYetImplemented).Methods("GET")

	//[Route("/System/Ping", "POST")]
	systemRouter.HandleFunc("/Ping", notYetImplemented).Methods("POST")

	//[Route("/System/Restart", "POST", Summary = "Restarts the application, if needed")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Restart", notYetImplemented).Methods("POST")

	//[Route("/System/Shutdown", "POST", Summary = "Shuts down the application")]
	systemRouter.HandleFunc("/Shutdown", notYetImplemented).Methods("POST")

	//[Route("/System/Logs", "GET", Summary = "Gets a list of available server log files")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Logs", notYetImplemented).Methods("GET")

	//[Route("/System/Endpoint", "GET", Summary = "Gets information about the request endpoint")]
	//[Authenticated]
	systemRouter.HandleFunc("/Endpoint", notYetImplemented).Methods("GET")

	//[Route("/System/Logs/Log", "GET", Summary = "Gets a log file")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Name", Description = "The log file name.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	systemRouter.HandleFunc("/Logs/Log", notYetImplemented).Methods("GET")
}
