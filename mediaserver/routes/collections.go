package routes

import "github.com/gorilla/mux"

func newCollectionsRouter(router *mux.Router) {
	/*
		/Collections
	*/
	collectionsRouter := router.PathPrefix("/Collections").Subrouter()

	//[Route("/Collections", "POST", Summary = "Creates a new collection")]
	//[ApiMember(Name = "IsLocked", Description = "Whether or not to lock the new collection.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Name", Description = "The name of the new collection.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ParentId", Description = "Optional - create the collection within a specific folder", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Ids", Description = "Item Ids to add to the collection", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	router.HandleFunc("/Collections", notYetImplemented).Methods("GET")

	//[Route("/Collections/{Id}/Items", "POST", Summary = "Adds items to a collection")]
	//[ApiMember(Name = "Ids", Description = "Item id, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	collectionsRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("POST")

	//[Route("/Collections/{Id}/Items", "DELETE", Summary = "Removes items from a collection")]
	//[ApiMember(Name = "Ids", Description = "Item id, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	collectionsRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("DELETE")
}
