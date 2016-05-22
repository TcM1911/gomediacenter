package routes

import "github.com/gorilla/mux"

func newDlnaRouter(router *mux.Router) {
	/*
		/Dlna
	*/
	dlnaRouter := router.PathPrefix("/Dlna").Subrouter()

	//[Route("/Dlna/{UuId}/description.xml", "GET", Summary = "Gets dlna server info")]
	//[Route("/Dlna/{UuId}/description", "GET", Summary = "Gets dlna server info")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/description", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/contentdirectory/contentdirectory.xml", "GET", Summary = "Gets dlna content directory xml")]
	//[Route("/Dlna/{UuId}/contentdirectory/contentdirectory", "GET", Summary = "Gets dlna content directory xml")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/contentdirectory/contentdirectory.xml", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/connectionmanager/connectionmanager.xml", "GET", Summary = "Gets dlna connection manager xml")]
	//[Route("/Dlna/{UuId}/connectionmanager/connectionmanager", "GET", Summary = "Gets dlna connection manager xml")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/connectionmanager/connectionmanager.xml", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/mediareceiverregistrar/mediareceiverregistrar.xml", "GET", Summary = "Gets dlna mediareceiverregistrar xml")]
	//[Route("/Dlna/{UuId}/mediareceiverregistrar/mediareceiverregistrar", "GET", Summary = "Gets dlna mediareceiverregistrar xml")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/mediareceiverregistrar/mediareceiverregistrar.xml", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/contentdirectory/control", "POST", Summary = "Processes a control request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/contentdirectory/control", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/connectionmanager/control", "POST", Summary = "Processes a control request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/connectionmanager/control", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/mediareceiverregistrar/control", "POST", Summary = "Processes a control request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/mediareceiverregistrar/control", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/mediareceiverregistrar/events", Summary = "Processes an event subscription request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "SUBSCRIBE,POST")]
	dlnaRouter.HandleFunc("/{uuid}/mediareceiverregistrar/events", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/contentdirectory/events", Summary = "Processes an event subscription request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "SUBSCRIBE,POST")]
	dlnaRouter.HandleFunc("/{uuid}/contentdirectory/events", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/connectionmanager/events", Summary = "Processes an event subscription request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "SUBSCRIBE,POST")]
	dlnaRouter.HandleFunc("/{uuid}/connectionmanager/events", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/icons/{Filename}", "GET", Summary = "Gets a server icon")]
	//[Route("/Dlna/icons/{Filename}", "GET", Summary = "Gets a server icon")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Filename", Description = "The icon filename", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/icons/{filename}", notYetImplemented).Methods("GET")

	//[Route("/Dlna/ProfileInfos", "GET", Summary = "Gets a list of profiles")]
	dlnaRouter.HandleFunc("/{ProfileInfos}", notYetImplemented).Methods("GET")

	//[Route("/Dlna/Profiles/{Id}", "DELETE", Summary = "Deletes a profile")]
	//[ApiMember(Name = "Id", Description = "Profile Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	dlnaRouter.HandleFunc("/Profiles/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Dlna/Profiles/Default", "GET", Summary = "Gets the default profile")]
	dlnaRouter.HandleFunc("/Profiles/Default", notYetImplemented).Methods("GET")

	//[Route("/Dlna/Profiles/{Id}", "GET", Summary = "Gets a single profile")]
	//[ApiMember(Name = "Id", Description = "Profile Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/Profiles/{id}", notYetImplemented).Methods("GET")

	//[Route("/Dlna/Profiles/{ProfileId}", "POST", Summary = "Updates a profile")]
	//[ApiMember(Name = "ProfileId", Description = "Profile Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/Profiles/{id}", notYetImplemented).Methods("POST")

	//[Route("/Dlna/Profiles", "POST", Summary = "Creates a profile")]
	dlnaRouter.HandleFunc("/Profiles", notYetImplemented).Methods("POST")
}
