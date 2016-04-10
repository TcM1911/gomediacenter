package api

//[Route("/Items/{Id}/File", "GET", Summary = "Gets the original file of an item")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Items/{Id}/CriticReviews", "GET", Summary = "Gets critic reviews for an item")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]

//[Route("/Items/{Id}/ThemeSongs", "GET", Summary = "Gets theme songs for an item")]
//[Authenticated]
//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Items/{Id}/ThemeVideos", "GET", Summary = "Gets theme videos for an item")]
//[Authenticated]
//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Items/{Id}/ThemeMedia", "GET", Summary = "Gets theme videos and songs for an item")]
//[Authenticated]
//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Library/Refresh", "POST", Summary = "Starts a library scan")]
//[Authenticated(Roles = "Admin")]

//[Route("/Items/{Id}", "DELETE", Summary = "Deletes an item from the library and file system")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Items", "DELETE", Summary = "Deletes an item from the library and file system")]
//[Authenticated]
//[ApiMember(Name = "Ids", Description = "Ids", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Items/Counts", "GET")]
//[Authenticated]
//[ApiMember(Name = "UserId", Description = "Optional. Get counts from a specific user's library.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "IsFavorite", Description = "Optional. Get counts of favorite items", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]

//[Route("/Items/{Id}/Ancestors", "GET", Summary = "Gets all parents of an item")]
//[Authenticated]
//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Items/YearIndex", "GET", Summary = "Gets a year index based on an item query.")]
//[Authenticated]
//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]

//[Route("/Library/PhysicalPaths", "GET", Summary = "Gets a list of physical paths from virtual folders")]
//[Authenticated(Roles = "Admin")]

//[Route("/Library/MediaFolders", "GET", Summary = "Gets all user media folders.")]
//[Authenticated]
//[ApiMember(Name = "IsHidden", Description = "Optional. Filter by folders that are marked hidden, or not.", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]

//[Route("/Library/Series/Added", "POST", Summary = "Reports that new episodes of a series have been added by an external source")]
//[Route("/Library/Series/Updated", "POST", Summary = "Reports that new episodes of a series have been added by an external source")]
//[Authenticated]
//[ApiMember(Name = "TvdbId", Description = "Tvdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Library/Movies/Added", "POST", Summary = "Reports that new movies have been added by an external source")]
//[Route("/Library/Movies/Updated", "POST", Summary = "Reports that new movies have been added by an external source")]
//[Authenticated]
//[ApiMember(Name = "TmdbId", Description = "Tmdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "ImdbId", Description = "Imdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Items/{Id}/Download", "GET", Summary = "Downloads item media")]
//[Authenticated(Roles = "download")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Items/{Id}/Similar", "GET", Summary = "Gets similar items")]
//[Authenticated]
