package api

//[Route("/Sync/Jobs/{Id}", "DELETE", Summary = "Cancels a sync job.")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Sync/Jobs/{Id}", "GET", Summary = "Gets a sync job.")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Sync/Jobs/{Id}", "POST", Summary = "Updates a sync job.")]

//[Route("/Sync/JobItems", "GET", Summary = "Gets sync job items.")]

//[Route("/Sync/JobItems/{Id}/Enable", "POST", Summary = "Enables a cancelled or queued sync job item")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sync/JobItems/{Id}/MarkForRemoval", "POST", Summary = "Marks a job item for removal")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sync/JobItems/{Id}/UnmarkForRemoval", "POST", Summary = "Unmarks a job item for removal")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sync/JobItems/{Id}", "DELETE", Summary = "Cancels a sync job item")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Sync/{TargetId}/Items", "DELETE", Summary = "Cancels items from a sync target")]
//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "Items")]
//[ApiMember(Name = "ItemIds", Description = "ItemIds", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "Items")]

//[Route("/Sync/Jobs", "GET", Summary = "Gets sync jobs.")]

//[Route("/Sync/Jobs", "POST", Summary = "Gets sync jobs.")]

//[Route("/Sync/Targets", "GET", Summary = "Gets a list of available sync targets.")]
//[ApiMember(Name = "UserId", Description = "UserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Sync/Options", "GET", Summary = "Gets a list of available sync targets.")]
//[ApiMember(Name = "UserId", Description = "UserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "ItemIds", Description = "ItemIds", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "ParentId", Description = "ParentId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "Category", Description = "Category", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Sync/JobItems/{Id}/Transferred", "POST", Summary = "Reports that a sync job item has successfully been transferred.")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sync/JobItems/{Id}/File", "GET", Summary = "Gets a sync job item file")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Sync/JobItems/{Id}/AdditionalFiles", "GET", Summary = "Gets a sync job item file")]
//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "Name", Description = "Name", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Sync/OfflineActions", "POST", Summary = "Reports an action that occurred while offline.")]

//[Route("/Sync/Items/Ready", "GET", Summary = "Gets ready to download sync items.")]
//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Sync/Data", "POST", Summary = "Syncs data between device and server")]
