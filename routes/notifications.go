package routes

import "github.com/gorilla/mux"

func newNotificationsRouter(router *mux.Router) {
	/*
		/Notifications
	*/
	notificationsRouter := router.PathPrefix("/Notifications").Subrouter()

	//[Route("/Notifications/{UserId}", "GET", Summary = "Gets notifications")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "IsRead", Description = "An optional filter by IsRead", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	notificationsRouter.HandleFunc("/{id}", notYetImplemented).Methods("GET")

	//[Route("/Notifications/{UserId}/Summary", "GET", Summary = "Gets a notification summary for a user")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	notificationsRouter.HandleFunc("/{id}/Summary", notYetImplemented).Methods("GET")

	//[Route("/Notifications/Types", "GET", Summary = "Gets notification types")]
	notificationsRouter.HandleFunc("/Types", notYetImplemented).Methods("GET")

	//[Route("/Notifications/Services", "GET", Summary = "Gets notification types")]
	notificationsRouter.HandleFunc("/Services", notYetImplemented).Methods("GET")

	//[Route("/Notifications/Admin", "POST", Summary = "Sends a notification to all admin users")]
	//[ApiMember(Name = "Name", Description = "The notification's name", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Description", Description = "The notification's description", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ImageUrl", Description = "The notification's image url", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Url", Description = "The notification's info url", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Level", Description = "The notification level", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	notificationsRouter.HandleFunc("/Admin", notYetImplemented).Methods("POST")

	//[Route("/Notifications/{UserId}/Read", "POST", Summary = "Marks notifications as read")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Ids", Description = "A list of notification ids, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	notificationsRouter.HandleFunc("/{id}/Read", notYetImplemented).Methods("POST")

	//[Route("/Notifications/{UserId}/Unread", "POST", Summary = "Marks notifications as unread")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Ids", Description = "A list of notification ids, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	notificationsRouter.HandleFunc("/{id}/Unread", notYetImplemented).Methods("POST")
}
