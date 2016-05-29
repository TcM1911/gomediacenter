package routes

import "github.com/gorilla/mux"

func newLiveTvRouter(router *mux.Router) {
	/*
		/LiveTv
	*/
	liveTvRoute := router.PathPrefix("/LiveTv").Subrouter()

	//[Route("/LiveTv/Info", "GET", Summary = "Gets available live tv services.")]
	//[Authenticated]
	liveTvRoute.HandleFunc("/Info", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Channels", "GET", Summary = "Gets available live tv channels.")]
	//[Authenticated]
	//[ApiMember(Name = "Type", Description = "Optional filter by channel type.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional filter by user and attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsFavorite", Description = "Filter by channels that are favorites, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsLiked", Description = "Filter by channels that are liked, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//ApiMember(Name = "IsDisliked", Description = "Filter by channels that are disliked, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableFavoriteSorting", Description = "Incorporate favorite and like status into channel sorting.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "AddCurrentProgram", Description = "Optional. Adds current program info to each channel", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Channels", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Channels/{Id}", "GET", Summary = "Gets a live tv channel")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Channel Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Channels/{id}", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Recordings", "GET", Summary = "Gets live tv recordings")]
	//[Authenticated]
	//[ApiMember(Name = "ChannelId", Description = "Optional filter by channel id.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional filter by user and attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "GroupId", Description = "Optional filter by recording group.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Status", Description = "Optional filter by recording status.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Status", Description = "Optional filter by recordings that are in progress, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "SeriesTimerId", Description = "Optional filter by recordings belonging to a series timer", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	liveTvRoute.HandleFunc("/Recordings", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Recordings/Groups", "GET", Summary = "Gets live tv recording groups")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional filter by user and attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Recordings/Groups", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Recordings/{Id}", "GET", Summary = "Gets a live tv recording")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Recording Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Recordings/{id}", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Tuners/{Id}/Reset", "POST", Summary = "Resets a tv tuner")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Tuner Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	liveTvRoute.HandleFunc("/Recordings/{id}/Reset", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/Timers/{Id}", "GET", Summary = "Gets a live tv timer")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Timer Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	liveTvRoute.HandleFunc("/Timers/{id}", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Timers/Defaults", "GET", Summary = "Gets default values for a new timer")]
	//[Authenticated]
	//[ApiMember(Name = "ProgramId", Description = "Optional, to attach default values based on a program.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Timers/Defaults", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Timers", "GET", Summary = "Gets live tv timers")]
	//[Authenticated]
	//[ApiMember(Name = "ChannelId", Description = "Optional filter by channel id.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "SeriesTimerId", Description = "Optional filter by timers belonging to a series timer", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Timers", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Programs", "GET,POST", Summary = "Gets available live tv epgs..")]
	//[Authenticated]
	//[ApiMember(Name = "ChannelIds", Description = "The channels to return guide information for.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "UserId", Description = "Optional filter by user id.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "MinStartDate", Description = "Optional. The minimum premiere date. Format = ISO", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "HasAired", Description = "Optional. Filter by programs that have completed airing, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "MaxStartDate", Description = "Optional. The maximum premiere date. Format = ISO", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "MinEndDate", Description = "Optional. The minimum premiere date. Format = ISO", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "MaxEndDate", Description = "Optional. The maximum premiere date. Format = ISO", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "IsMovie", Description = "Optional filter for movies.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "IsKids", Description = "Optional filter for kids.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "IsSports", Description = "Optional filter for sports.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "SortBy", Description = "Optional. Specify one or more sort orders, comma delimeted. Options: Name, StartDate", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "SortOrder", Description = "Sort Order - Ascending,Descending", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Genres", Description = "The genres to return guide information for.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	liveTvRoute.HandleFunc("/Programs", notYetImplemented).Methods("GET")
	liveTvRoute.HandleFunc("/Programs", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/Programs/Recommended", "GET", Summary = "Gets available live tv epgs..")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional filter by user id.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsAiring", Description = "Optional. Filter by programs that are currently airing, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "HasAired", Description = "Optional. Filter by programs that have completed airing, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsSports", Description = "Optional filter for sports.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "IsMovie", Description = "Optional filter for movies.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsKids", Description = "Optional filter for kids.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	liveTvRoute.HandleFunc("/Programs/Recommended", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Programs/{Id}", "GET", Summary = "Gets a live tv program")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Program Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Programs/{id}", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Recordings/{Id}", "DELETE", Summary = "Deletes a live tv recording")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Recording Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	liveTvRoute.HandleFunc("/Recordings/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/LiveTv/Timers/{Id}", "DELETE", Summary = "Cancels a live tv timer")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Timer Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	liveTvRoute.HandleFunc("/Timers/{id}", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Timers/{Id}", "POST", Summary = "Updates a live tv timer")]
	//[Authenticated]
	liveTvRoute.HandleFunc("/Timers/{id}", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/Timers", "POST", Summary = "Creates a live tv timer")]
	//[Authenticated]
	liveTvRoute.HandleFunc("/Timers", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/SeriesTimers/{Id}", "GET", Summary = "Gets a live tv series timer")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Timer Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	liveTvRoute.HandleFunc("/SeriesTimers/{id}", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/SeriesTimers", "GET", Summary = "Gets live tv series timers")]
	//[Authenticated]
	//[ApiMember(Name = "SortBy", Description = "Optional. Sort by SortName or Priority", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	//[ApiMember(Name = "SortOrder", Description = "Optional. Sort in Ascending or Descending order", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET,POST")]
	liveTvRoute.HandleFunc("/SeriesTimers", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/SeriesTimers/{Id}", "DELETE", Summary = "Cancels a live tv series timer")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Timer Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	liveTvRoute.HandleFunc("/SeriesTimers/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/LiveTv/SeriesTimers/{Id}", "POST", Summary = "Updates a live tv series timer")]
	//[Authenticated]
	liveTvRoute.HandleFunc("/SeriesTimers/{id}", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/SeriesTimers", "POST", Summary = "Creates a live tv series timer")]
	//[Authenticated]
	liveTvRoute.HandleFunc("/SeriesTimers", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/Recordings/Groups/{Id}", "GET", Summary = "Gets a recording group")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Recording group Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	liveTvRoute.HandleFunc("/Recordings/Groups/{id}", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/GuideInfo", "GET", Summary = "Gets guide info")]
	//[Authenticated]
	liveTvRoute.HandleFunc("/GuideInfo", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Folder", "GET", Summary = "Gets the users live tv folder, along with configured images")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Folder", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/TunerHosts", "POST", Summary = "Adds a tuner host")]
	//[Authenticated]
	liveTvRoute.HandleFunc("/TunerHosts", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/TunerHosts", "DELETE", Summary = "Deletes a tuner host")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Tuner host id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	liveTvRoute.HandleFunc("/TunerHosts", notYetImplemented).Methods("DELETE")

	//[Route("/LiveTv/ListingProviders", "POST", Summary = "Adds a listing provider")]
	//[Authenticated(AllowBeforeStartupWizard = true)]
	liveTvRoute.HandleFunc("/ListingProviders", notYetImplemented).Methods("POST")

	//[Route("/LiveTv/ListingProviders", "DELETE", Summary = "Deletes a listing provider")]
	//[Authenticated(AllowBeforeStartupWizard = true)]
	//[ApiMember(Name = "Id", Description = "Provider id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	liveTvRoute.HandleFunc("/ListingProviders", notYetImplemented).Methods("DELETE")

	//[Route("/LiveTv/ListingProviders/Lineups", "GET", Summary = "Gets available lineups")]
	//[Authenticated(AllowBeforeStartupWizard = true)]
	//[ApiMember(Name = "Id", Description = "Provider id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Type", Description = "Provider Type", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Location", Description = "Location", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Country", Description = "Country", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/ListingProviders/Lineups", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/ListingProviders/SchedulesDirect/Countries", "GET", Summary = "Gets available lineups")]
	//[Authenticated(AllowBeforeStartupWizard = true)]
	liveTvRoute.HandleFunc("/ListingProviders/SchedulesDirect/Countries", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/Registration", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "ChannelId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ProgramId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Feature", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	liveTvRoute.HandleFunc("/Registration", notYetImplemented).Methods("GET")

	//[Route("/LiveTv/TunerHosts/Satip/IniMappings", "GET", Summary = "Gets available mappings")]
	//[Authenticated(AllowBeforeStartupWizard = true)]
	liveTvRoute.HandleFunc("/TunerHosts/Satip/IniMappings", notYetImplemented).Methods("GET")
}