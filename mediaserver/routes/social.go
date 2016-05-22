package routes

import "github.com/gorilla/mux"

func newSocialRouter(router *mux.Router) {
	/*
		/Social
	*/
	socialRouter := router.PathPrefix("/Social").Subrouter()

	//[Route("/Social/Shares/{Id}", "GET", Summary = "Gets a share")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/{id}", notYetImplemented).Methods("GET")

	//[Route("/Social/Shares/Public/{Id}", "GET", Summary = "Gets a share")]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/Public/{id}", notYetImplemented).Methods("GET")

	//[Route("/Social/Shares/Public/{Id}/Image", "GET", Summary = "Gets a share")]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/Public/{id}/Images", notYetImplemented).Methods("GET")

	//[Route("/Social/Shares", "POST", Summary = "Creates a share")]
	//[Authenticated]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "UserId", Description = "The user id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	socialRouter.HandleFunc("/Shares", notYetImplemented).Methods("POST")

	//[Route("/Social/Shares/{Id}", "DELETE", Summary = "Deletes a share")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	socialRouter.HandleFunc("/Shares", notYetImplemented).Methods("DELETE")

	//[Route("/Social/Shares/Public/{Id}/Item", "GET", Summary = "Gets a share")]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/Public/{id}/Item", notYetImplemented).Methods("GET")
}
