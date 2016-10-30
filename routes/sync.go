package routes

import "github.com/gorilla/mux"

func newSyncRouter(router *mux.Router) {
	/*
		/Sync
	*/
	syncRouter := router.PathPrefix("/Sync").Subrouter()

	//[Route("/Sync/Jobs/{Id}", "DELETE", Summary = "Cancels a sync job.")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	syncRouter.HandleFunc("/Jobs/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Sync/Jobs/{Id}", "GET", Summary = "Gets a sync job.")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	syncRouter.HandleFunc("/Jobs/{id}", notYetImplemented).Methods("GET")

	//[Route("/Sync/Jobs/{Id}", "POST", Summary = "Updates a sync job.")]
	syncRouter.HandleFunc("/Jobs/{id}", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems", "GET", Summary = "Gets sync job items.")]
	syncRouter.HandleFunc("/JobItems", notYetImplemented).Methods("GET")

	//[Route("/Sync/JobItems/{Id}/Enable", "POST", Summary = "Enables a cancelled or queued sync job item")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/Enable", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}/MarkForRemoval", "POST", Summary = "Marks a job item for removal")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/MarkForRemoval", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}/UnmarkForRemoval", "POST", Summary = "Unmarks a job item for removal")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/UnmarkForRemoval", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}", "DELETE", Summary = "Cancels a sync job item")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	syncRouter.HandleFunc("/JobItems/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Sync/{TargetId}/Items", "DELETE", Summary = "Cancels items from a sync target")]
	//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "Items")]
	//[ApiMember(Name = "ItemIds", Description = "ItemIds", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "Items")]
	syncRouter.HandleFunc("/{targetId}/Items", notYetImplemented).Methods("DELETE")

	//[Route("/Sync/Jobs", "GET", Summary = "Gets sync jobs.")]
	syncRouter.HandleFunc("/Jobs", notYetImplemented).Methods("GET")

	//[Route("/Sync/Jobs", "POST", Summary = "Gets sync jobs.")]
	syncRouter.HandleFunc("/Jobs", notYetImplemented).Methods("POST")

	//[Route("/Sync/Targets", "GET", Summary = "Gets a list of available sync targets.")]
	//[ApiMember(Name = "UserId", Description = "UserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/Targets", notYetImplemented).Methods("GET")

	//[Route("/Sync/Options", "GET", Summary = "Gets a list of available sync targets.")]
	//[ApiMember(Name = "UserId", Description = "UserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ItemIds", Description = "ItemIds", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "ParentId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Category", Description = "Category", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/Options", notYetImplemented).Methods("GET")

	//[Route("/Sync/JobItems/{Id}/Transferred", "POST", Summary = "Reports that a sync job item has successfully been transferred.")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/Transferred", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}/File", "GET", Summary = "Gets a sync job item file")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	syncRouter.HandleFunc("/JobItems/{id}/File", notYetImplemented).Methods("GET")

	//[Route("/Sync/JobItems/{Id}/AdditionalFiles", "GET", Summary = "Gets a sync job item file")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Name", Description = "Name", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/JobItems/{id}/AdditionalFiles", notYetImplemented).Methods("GET")

	//[Route("/Sync/OfflineActions", "POST", Summary = "Reports an action that occurred while offline.")]
	syncRouter.HandleFunc("/OfflineActions", notYetImplemented).Methods("POST")

	//[Route("/Sync/Items/Ready", "GET", Summary = "Gets ready to download sync items.")]
	//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/Items/Ready", notYetImplemented).Methods("GET")

	//[Route("/Sync/Data", "POST", Summary = "Syncs data between device and server")]
	syncRouter.HandleFunc("/Data", notYetImplemented).Methods("POST")
}
