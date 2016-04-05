package api

//[Route("/Items/{Id}/Images", "GET", Summary = "Gets information about an item's images")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Items/{Id}/Images/{Type}", "GET")]
//[Route("/Items/{Id}/Images/{Type}/{Index}", "GET")]
//[Route("/Items/{Id}/Images/{Type}", "HEAD")]
//[Route("/Items/{Id}/Images/{Type}/{Index}", "HEAD")]
//[Route("/Items/{Id}/Images/{Type}/{Index}/{Tag}/{Format}/{MaxWidth}/{MaxHeight}/{PercentPlayed}/{UnplayedCount}", "GET")]

//[Route("/Items/{Id}/Images/{Type}/{Index}/{Tag}/{Format}/{MaxWidth}/{MaxHeight}/{PercentPlayed}/{UnplayedCount}", "HEAD")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Items/{Id}/Images/{Type}/{Index}/Index", "POST", Summary = "Updates the index for an item image")]
//[Authenticated(Roles = "admin")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "Type", Description = "Image Type", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Index", Description = "Image Index", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "NewIndex", Description = "The new image index", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Artists/{Name}/Images/{Type}", "GET")]
//[Route("/Artists/{Name}/Images/{Type}/{Index}", "GET")]
//[Route("/Genres/{Name}/Images/{Type}", "GET")]
//[Route("/Genres/{Name}/Images/{Type}/{Index}", "GET")]
//[Route("/GameGenres/{Name}/Images/{Type}", "GET")]
//[Route("/GameGenres/{Name}/Images/{Type}/{Index}", "GET")]
//[Route("/MusicGenres/{Name}/Images/{Type}", "GET")]
//[Route("/MusicGenres/{Name}/Images/{Type}/{Index}", "GET")]
//[Route("/Persons/{Name}/Images/{Type}", "GET")]
//[Route("/Persons/{Name}/Images/{Type}/{Index}", "GET")]
//[Route("/Studios/{Name}/Images/{Type}", "GET")]
//[Route("/Studios/{Name}/Images/{Type}/{Index}", "GET")]
//[Route("/Years/{Year}/Images/{Type}", "GET")]
//[Route("/Years/{Year}/Images/{Type}/{Index}", "GET")]
//[Route("/Artists/{Name}/Images/{Type}", "HEAD")]
//[Route("/Artists/{Name}/Images/{Type}/{Index}", "HEAD")]
//[Route("/Genres/{Name}/Images/{Type}", "HEAD")]
//[Route("/Genres/{Name}/Images/{Type}/{Index}", "HEAD")]
//[Route("/GameGenres/{Name}/Images/{Type}", "HEAD")]
//[Route("/GameGenres/{Name}/Images/{Type}/{Index}", "HEAD")]
//[Route("/MusicGenres/{Name}/Images/{Type}", "HEAD")]
//[Route("/MusicGenres/{Name}/Images/{Type}/{Index}", "HEAD")]
//[Route("/Persons/{Name}/Images/{Type}", "HEAD")]
//[Route("/Persons/{Name}/Images/{Type}/{Index}", "HEAD")]
//[Route("/Studios/{Name}/Images/{Type}", "HEAD")]
//[Route("/Studios/{Name}/Images/{Type}/{Index}", "HEAD")]
//[Route("/Years/{Year}/Images/{Type}", "HEAD")]

//[Route("/Years/{Year}/Images/{Type}/{Index}", "HEAD")]
//[ApiMember(Name = "Name", Description = "Item name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Users/{Id}/Images/{Type}", "GET")]
//[Route("/Users/{Id}/Images/{Type}/{Index}", "GET")]
//[Route("/Users/{Id}/Images/{Type}", "HEAD")]

//[Route("/Users/{Id}/Images/{Type}/{Index}", "HEAD")]
//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Items/{Id}/Images/{Type}", "DELETE")]

//[Route("/Items/{Id}/Images/{Type}/{Index}", "DELETE")]
//[Authenticated(Roles = "admin")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Users/{Id}/Images/{Type}", "DELETE")]

//[Route("/Users/{Id}/Images/{Type}/{Index}", "DELETE")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Users/{Id}/Images/{Type}", "POST")]
//[Route("/Users/{Id}/Images/{Type}/{Index}", "POST")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Items/{Id}/Images/{Type}", "POST")]

//[Route("/Items/{Id}/Images/{Type}/{Index}", "POST")]
//[Api(Description = "Posts an item image")]
//[Authenticated(Roles = "admin")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
