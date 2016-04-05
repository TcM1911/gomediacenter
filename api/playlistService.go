package api

//[Route("/Playlists", "POST", Summary = "Creates a new playlist")]
//[ApiMember(Name = "Name", Description = "The name of the new playlist.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "Ids", Description = "Item Ids to add to the playlist", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "MediaType", Description = "The playlist media type", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Playlists/{Id}/Items", "POST", Summary = "Adds items to a playlist")]
//[ApiMember(Name = "Ids", Description = "Item id, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]

//[Route("/Playlists/{Id}/Items/{ItemId}/Move/{NewIndex}", "POST", Summary = "Moves a playlist item")]
//[ApiMember(Name = "ItemId", Description = "ItemId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "NewIndex", Description = "NewIndex", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]

//[Route("/Playlists/{Id}/Items", "DELETE", Summary = "Removes items from a playlist")]
//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
//[ApiMember(Name = "EntryIds", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]

//[Route("/Playlists/{Id}/Items", "GET", Summary = "Gets the original items of a playlist")]
//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
