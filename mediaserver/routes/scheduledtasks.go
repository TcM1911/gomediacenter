package routes

import "github.com/gorilla/mux"

func newScheduledTasksRouter(router *mux.Router) {
	/*
		/ScheduledTasks
	*/
	scheduledTasksRouter := router.PathPrefix("/ScheduledTasks").Subrouter()

	//[Route("/ScheduledTasks/{Id}", "GET", Summary = "Gets a scheduled task, by Id")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	scheduledTasksRouter.HandleFunc("/{id}", notYetImplemented).Methods("GET")

	//[Route("/ScheduledTasks", "GET", Summary = "Gets scheduled tasks")]
	//[ApiMember(Name = "IsHidden", Description = "Optional filter tasks that are hidden, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsEnabled", Description = "Optional filter tasks that are enabled, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	router.HandleFunc("/ScheduledTasks", notYetImplemented).Methods("GET")

	//[Route("/ScheduledTasks/Running/{Id}", "POST", Summary = "Starts a scheduled task")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	scheduledTasksRouter.HandleFunc("/Running/{id}", notYetImplemented).Methods("POST")

	//[Route("/ScheduledTasks/Running/{Id}", "DELETE", Summary = "Stops a scheduled task")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	scheduledTasksRouter.HandleFunc("/Running/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/ScheduledTasks/{Id}/Triggers", "POST", Summary = "Updates the triggers for a scheduled task")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	scheduledTasksRouter.HandleFunc("/{id}/Triggers", notYetImplemented).Methods("POST")
}
