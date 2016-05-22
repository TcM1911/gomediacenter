package routes

import "github.com/gorilla/mux"

func newItemsRouter(router *mux.Router) {
	/*
		/Items router
	*/
	itemsRouter := router.PathPrefix("/Items").Subrouter()

	//[Route("/Items/Counts", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Get counts from a specific user's library.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsFavorite", Description = "Optional. Get counts of favorite items", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/Counts", notYetImplemented).Methods("GET")

	//[Route("/Items/Filters", "GET", Summary = "Gets branding configuration")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "Specify this to localize the search to a specific item or folder. Omit to use the root", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "MediaTypes", Description = "Optional filter by MediaType. Allows multiple, comma delimited.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	itemsRouter.HandleFunc("/Filters", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}", "DELETE", Summary = "Deletes an item from the library and file system")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	itemsRouter.HandleFunc("/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Items/{Id}/Images", "GET", Summary = "Gets information about an item's images")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Images", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}", "GET")]
	//[Route("/Items/{Id}/Images/{Type}", "HEAD")]
	itemsRouter.HandleFunc("/{id}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}/{Index}", "GET")]
	//[Route("/Items/{Id}/Images/{Type}/{Index}", "HEAD")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}/{Index}/{Tag}/{Format}/{MaxWidth}/{MaxHeight}/{PercentPlayed}/{UnplayedCount}", "HEAD")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[Route("/Items/{Id}/Images/{Type}/{Index}/{Tag}/{Format}/{MaxWidth}/{MaxHeight}/{PercentPlayed}/{UnplayedCount}", "GET")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}/{tag}/{format}/{maxWidth}/{maxHeight}/{percentPlayed}/{unplayedCount}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}/{Index}/Index", "POST", Summary = "Updates the index for an item image")]
	//[Authenticated(Roles = "admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Type", Description = "Image Type", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Index", Description = "Image Index", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "NewIndex", Description = "The new image index", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}/Index", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Images/{Type}", "POST")]
	itemsRouter.HandleFunc("/{id}/Images/{type}", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Images/{Type}/{Index}", "POST")]
	//[Api(Description = "Posts an item image")]
	//[Authenticated(Roles = "admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Ancestors", "GET", Summary = "Gets all parents of an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Ancestors", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/CriticReviews", "GET", Summary = "Gets critic reviews for an item")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/CriticReviews", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Download", "GET", Summary = "Downloads item media")]
	//[Authenticated(Roles = "download")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Download", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/File", "GET", Summary = "Gets the original file of an item")]
	//[Authenticated]
	itemsRouter.HandleFunc("/{id}/File", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Similar", "GET", Summary = "Gets similar items")]
	//[Authenticated]
	itemsRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ThemeMedia", "GET", Summary = "Gets theme videos and songs for an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ThemeMedia", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ThemeSongs", "GET", Summary = "Gets theme songs for an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ThemeSongs", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ThemeVideos", "GET", Summary = "Gets theme videos for an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ThemeVideos", notYetImplemented).Methods("GET")

	//[Route("/Items/YearIndex", "GET", Summary = "Gets a year index based on an item query.")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	itemsRouter.HandleFunc("/YearIndex", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given item")]
	itemsRouter.HandleFunc("/{id}/InstantMix", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ExternalIdInfos", "GET", Summary = "Gets external id infos for an item")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ExternalIdInfos", notYetImplemented).Methods("GET")

	//[Route("/Items/RemoteSearch/Movie", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/Movie", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/AdultVideo", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/AdultMovie", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Series", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/Series", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Game", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/Game", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/BoxSet", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/BoxSet", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/MusicArtist", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/MusicArtist", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/MusicAlbum", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/MusicAlbum", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Person", "POST")]
	//[Authenticated(Roles = "Admin")]
	itemsRouter.HandleFunc("/RemoteSearch/Person", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Image", "GET", Summary = "Gets a remote image")]
	//[ApiMember(Name = "ImageUrl", Description = "The image url", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ProviderName", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/RemoteSearch/Image", notYetImplemented).Methods("GET")

	//[Route("/Items/RemoteSearch/Apply/{Id}", "POST", Summary = "Applies search criteria to an item and refreshes metadata")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "The item id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ReplaceAllImages", Description = "Whether or not to replace all images", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	itemsRouter.HandleFunc("/RemoteSearch/Apply/{id}", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Refresh", "POST", Summary = "Refreshes metadata for an item")]
	//[ApiMember(Name = "Recursive", Description = "Indicates if the refresh should occur recursively.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/Refresh", notYetImplemented).Methods("POST")

	//[Route("/Items/{ItemId}", "POST", Summary = "Updates an item")]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}", notYetImplemented).Methods("POST")

	//[Route("/Items/{ItemId}/MetadataEditor", "GET", Summary = "Gets metadata editor info for an item")]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/MetadataEditor", notYetImplemented).Methods("GET")

	//[Route("/Items/{ItemId}/ContentType", "POST", Summary = "Updates an item's content type")]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "ContentType", Description = "The content type of the item", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/ContentType", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/PlaybackInfo", "GET", Summary = "Gets live playback media info for an item")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/PlaybackInfo", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/PlaybackInfo", "POST", Summary = "Gets live playback media info for an item")]
	//[ApiMember(Name = "Type", Description = "The image type", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ProviderName", Description = "Optional. The image provider to use", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeAllLanguages", Description = "Optional.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/PlaybackInfo", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/RemoteImages", "GET", Summary = "Gets available remote images for an item")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteImages", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteImages/Providers", "GET", Summary = "Gets available remote image providers for an item")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Type", Description = "The image type", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ProviderName", Description = "The image provider", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageUrl", Description = "The image url", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteImages/Providers", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteImages/Download", "POST", Summary = "Downloads a remote image for an item")]
	//[Authenticated(Roles="Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteImages/Download", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/RemoteSearch/Subtitles/{Language}", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Language", Description = "Language", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteSearch/Subtitles/{lang}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteSearch/Subtitles/Providers", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteSearch/Subtitles/Providers", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteSearch/Subtitles/{SubtitleId}", "POST")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "SubtitleId", Description = "SubtitleId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/RemoteSearch/Subtitles/{subid}", notYetImplemented).Methods("POST")
}
