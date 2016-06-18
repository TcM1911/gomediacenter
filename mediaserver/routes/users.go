package routes

import (
	"github.com/gorilla/mux"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers"
	"github.com/tcm1911/gomediacenter/mediaserver/middleware"
)

func newUsersRouter(router *mux.Router) {
	/*
		/Users
	*/
	usersRouter := router.PathPrefix("/Users").Subrouter()

	// GET to the route "/Users" gets a list of users. This action required the user
	// to be logged in. The list can be filtered by the query parameters: IsHidden,
	// IsDisabled, and IsGuest.
	router.HandleFunc("/Users", middleware.WithContext(
		middleware.WithQueryVars(
			middleware.WithDB(
				controllers.GetAllUsers)))).Methods("GET")

	// GET to the route"/Users/Public" gets a list of publicly visible users
	// for display on a login screen.
	usersRouter.HandleFunc("/Public", middleware.WithContext(
		middleware.WithDB(
			controllers.GetAllUsersPublic))).Methods("GET")

	// POST to the route /Users/New creates a new user. This action requires
	// admin status. The username is sent in the body with the parameter name of 'Name'.
	usersRouter.HandleFunc("/New", middleware.WithContext(
		middleware.Admin(
			middleware.WithDB(
				controllers.NewUser)))).Methods("POST")

	// GET to the route /Users/{uid} returns the data about a user. This actions
	// requires admin status or being login as the user.
	usersRouter.HandleFunc("/{uid}", middleware.WithContext(
		middleware.WithPathVars(
			middleware.AdminOrLoggedInUser(
				middleware.WithDB(
					controllers.GetUserById))))).Methods("GET")

	// GET to the route "/Users/{Id}/Offline" gets an offline user record by Id.
	usersRouter.HandleFunc("/{uid}/Offline", middleware.WithContext(
		middleware.WithPathVars(
			middleware.AdminOrLoggedInUser(
				middleware.WithDB(
					controllers.GetOfflineUserById))))).Methods("GET")

	//A DELETE to /Users/{uid} deletes a user and all it's item data.
	// This action requires admin rights.
	usersRouter.HandleFunc("/{uid}", middleware.WithContext(
		middleware.WithPathVars(
			middleware.Admin(
				middleware.WithDB(
					controllers.DeleteUser))))).Methods("DELETE")

	// A POST to /Users/{uid}/Authenticate authenticates a user.
	// The password is past in the body in the parameter password.
	usersRouter.HandleFunc("/{uid}/Authenticate",
		middleware.WithContext(
			middleware.WithPathVars(
				middleware.VerifyIds([]string{"uid"},
					middleware.WithDB(
						controllers.Authenticate))))).Methods("POST")

	// A POST to /Users/AuthenticateByName authenticates a user.
	// Username and password is past in the body as the parameters Username and password.
	usersRouter.HandleFunc("/AuthenticateByName", middleware.WithContext(
		middleware.WithDB(
			controllers.AuthenticateByName))).Methods("POST")

	// A POST to /Users/{uid}/Logout logs the user out.
	usersRouter.HandleFunc("/{uid}/Logout",
		middleware.WithContext(
			middleware.WithPathVars(
				middleware.VerifyIds([]string{"uid"},
					controllers.LogoutUser)))).Methods("POST")

	// A POST to /Users/{Id}/Password updates a user's password.
	// New password and current password are past as body parameters
	// newPassword and currentPassword.
	usersRouter.HandleFunc("/{uid}/Password", middleware.WithContext(
		middleware.WithPathVars(
			middleware.VerifyIds([]string{"uid"},
				middleware.AdminOrLoggedInUser(
					middleware.WithQueryVars(
						middleware.WithDB(
							controllers.ChangeUserPassword))))))).Methods("POST")

	//[Route("/Users/{Id}/EasyPassword", "POST", Summary = "Updates a user's easy password")]
	//[Authenticated]
	usersRouter.HandleFunc("/{id}/EasyPassword", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}", "POST", Summary = "Updates a user")]
	//[Authenticated]
	usersRouter.HandleFunc("/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/Policy", "POST", Summary = "Updates a user policy")]
	//[Authenticated(Roles = "admin")]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Policy", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/Configuration", "POST", Summary = "Updates a user configuration")]
	//[Authenticated]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Configuration", notYetImplemented).Methods("POST")

	//[Route("/Users/ForgotPassword", "POST", Summary = "Initiates the forgot password process for a local user")]
	//[ApiMember(Name = "EnteredUsername", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/ForgotPassword", notYetImplemented).Methods("POST")

	//[Route("/Users/ForgotPassword/Pin", "POST", Summary = "Redeems a forgot password pin")]
	//[ApiMember(Name = "Pin", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/ForgotPassword/Pin", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/Views", "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "IncludeExternalContent", Description = "Whether or not to include external views such as channels or live tv", IsRequired = true, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Views", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/SpecialViewOptions", "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}/SpecialViewOptions", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/GroupingOptions", "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}/GroupingOptions", notYetImplemented).Methods("GET")

	//[Route("/Users/{Id}/Connect/Link", "POST", Summary = "Creates a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ConnectUsername", Description = "Connect username", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Connect/Link", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/Connect/Link", "DELETE", Summary = "Removes a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	usersRouter.HandleFunc("/{id}/Connect/Link", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{Id}/Images/{Type}", "GET")]
	//[Route("/Users/{Id}/Images/{Type}", "HEAD")]
	usersRouter.HandleFunc("/{id}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Users/{Id}/Images/{Type}/{Index}", "GET")]
	//[Route("/Users/{Id}/Images/{Type}/{Index}", "HEAD")]
	//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/PlayedItems/{Id}", "POST", Summary = "Marks an item as played")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "DatePlayed", Description = "The date the item was played (if any). Format = yyyyMMddHHmmss", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/PlayedItems/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/PlayedItems/{Id}", "DELETE", Summary = "Marks an item as unplayed")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	usersRouter.HandleFunc("/{uid}/PlayedItems/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{UserId}/PlayingItems/{Id}", "POST", Summary = "Reports that a user has begun playing an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "MediaSourceId", Description = "The id of the MediaSource", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "CanSeek", Description = "Indicates if the client can seek", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "QueueableMediaTypes", Description = "A list of media types that can be queued from this item, comma delimited. Audio,Video,Book,Game", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	//[ApiMember(Name = "AudioStreamIndex", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SubtitleStreamIndex", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "PlayMethod", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "LiveStreamId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "PlaySessionId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/PlayingItems/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/Items/{Id}", "GET", Summary = "Gets an item from a user's library")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}",
		middleware.WithContext(
			middleware.WithPathVars(
				middleware.AdminOrLoggedInUser(
					middleware.WithDB(
						controllers.UserItemHandler))))).Methods("GET")

	//[Route("/Users/{UserId}/Items/Root", "GET", Summary = "Gets the root folder from a user's library")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/Root", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/Items/{Id}/Intros", "GET", Summary = "Gets intros to play before the main media item plays")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/Intros", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/FavoriteItems/{Id}", "POST", Summary = "Marks an item as a favorite")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/FavoriteItems/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/FavoriteItems/{Id}", "DELETE", Summary = "Unmarks an item as a favorite")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	usersRouter.HandleFunc("/{uid}/FavoriteItems/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{UserId}/Items/{Id}/Rating", "DELETE", Summary = "Deletes a user's saved personal rating for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/Rating", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{UserId}/Items/{Id}/Rating", "POST", Summary = "Updates a user's rating for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Likes", Description = "Whether the user likes the item or not. true/false", IsRequired = true, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/Rating", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/Items/{Id}/LocalTrailers", "GET", Summary = "Gets local trailers for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/LocalTrailers", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/Items/{Id}/SpecialFeatures", "GET", Summary = "Gets special features for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Movie Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/SpecialFeatures", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/Items/Latest", "GET", Summary = "Gets latest media")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Limit", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "Specify this to localize the search to a specific item or folder. Omit to use the root", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "IsFolder", Description = "Filter by items that are folders, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsPlayed", Description = "Filter by items that are played, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "GroupItems", Description = "Whether or not to group items into a parent container.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/Latest", notYetImplemented).Methods("GET")
}
