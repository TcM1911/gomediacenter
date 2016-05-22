package routes

import "github.com/gorilla/mux"

func newLibraryRouter(router *mux.Router) {
	/*
		/Library router
	*/
	libraryRouter := router.PathPrefix("/Library").Subrouter()

	//[Route("/Library/FileOrganization", "GET", Summary = "Gets file organization results")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	libraryRouter.HandleFunc("/FileOrganization", notYetImplemented).Methods("GET")

	//[Route("/Library/FileOrganizations", "DELETE", Summary = "Clears the activity log")]
	//[Route("/Library/FileOrganizations/{Id}/File", "DELETE", Summary = "Deletes the original file of a organizer result")]
	//[ApiMember(Name = "Id", Description = "Result Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	libraryRouter.HandleFunc("/FileOrganization", notYetImplemented).Methods("DELETE")

	//[Route("/Library/FileOrganizations/{Id}/Organize", "POST", Summary = "Performs an organization")]
	//[ApiMember(Name = "Id", Description = "Result Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	libraryRouter.HandleFunc("/FileOrganization/{id}/Organize", notYetImplemented).Methods("POST")

	//[Route("/Library/FileOrganizations/{Id}/Episode/Organize", "POST", Summary = "Performs organization of a tv episode")]
	//[ApiMember(Name = "Id", Description = "Result Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "SeriesId", Description = "Series Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SeasonNumber", IsRequired = true, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "EpisodeNumber", IsRequired = true, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "EndingEpisodeNumber", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "RememberCorrection", Description = "Whether or not to apply the same correction to future episodes of the same series.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "NewSeriesProviderIds", Description = "A list of provider IDs identifying a new series.", IsRequired = false, DataType = "Dictionary<string, string>", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "NewSeriesName", Description = "Name of a series to add.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "NewSeriesYear", Description = "Year of a series to add.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "TargetFolder", Description = "Target Folder", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	libraryRouter.HandleFunc("/FileOrganization/{id}/Episode/Organize", notYetImplemented).Methods("POST")

	//[Route("/Library/FileOrganizations/SmartMatches", "GET", Summary = "Gets smart match entries")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	libraryRouter.HandleFunc("/FileOrganization/SmartMatches", notYetImplemented).Methods("GET")

	//[Route("/Library/FileOrganizations/SmartMatches/Delete", "POST", Summary = "Deletes a smart match entry")]
	//[ApiMember(Name = "Entries", Description = "SmartMatch Entry", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	libraryRouter.HandleFunc("/FileOrganization/SmartMatches/Delete", notYetImplemented).Methods("POST")

	//[Route("/Library/MediaFolders", "GET", Summary = "Gets all user media folders.")]
	//[Authenticated]
	//[ApiMember(Name = "IsHidden", Description = "Optional. Filter by folders that are marked hidden, or not.", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	libraryRouter.HandleFunc("/MediaFolders", notYetImplemented).Methods("GET")

	//[Route("/Library/Movies/Added", "POST", Summary = "Reports that new movies have been added by an external source")]
	//[Authenticated]
	//[ApiMember(Name = "TmdbId", Description = "Tmdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "ImdbId", Description = "Imdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	libraryRouter.HandleFunc("/Movies/Added", notYetImplemented).Methods("POST")

	//[Route("/Library/Movies/Updated", "POST", Summary = "Reports that new movies have been added by an external source")]
	//[Authenticated]
	//[ApiMember(Name = "TmdbId", Description = "Tmdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "ImdbId", Description = "Imdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	libraryRouter.HandleFunc("/Movies/Updated", notYetImplemented).Methods("POST")

	//[Route("/Library/PhysicalPaths", "GET", Summary = "Gets a list of physical paths from virtual folders")]
	//[Authenticated(Roles = "Admin")]
	libraryRouter.HandleFunc("/PhysicalPaths", notYetImplemented).Methods("GET")

	//[Route("/Library/Refresh", "POST", Summary = "Starts a library scan")]
	//[Authenticated(Roles = "Admin")]
	libraryRouter.HandleFunc("/Refresh", notYetImplemented).Methods("POST")

	//[Route("/Library/Series/Added", "POST", Summary = "Reports that new episodes of a series have been added by an external source")]
	//[Authenticated]
	//[ApiMember(Name = "TvdbId", Description = "Tvdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	libraryRouter.HandleFunc("/Series/Added", notYetImplemented).Methods("POST")

	//[Route("/Library/Series/Updated", "POST", Summary = "Reports that new episodes of a series have been added by an external source")]
	//[Authenticated]
	//[ApiMember(Name = "TvdbId", Description = "Tvdb Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	libraryRouter.HandleFunc("/Series/Updated", notYetImplemented).Methods("POST")

	//[Route("/Library/VirtualFolders", "GET")]
	libraryRouter.HandleFunc("/VirtualFolders", notYetImplemented).Methods("GET")

	//[Route("/Library/VirtualFolders", "POST")]
	libraryRouter.HandleFunc("/VirtualFolders", notYetImplemented).Methods("POST")

	//[Route("/Library/VirtualFolders", "DELETE")]
	libraryRouter.HandleFunc("/VirtualFolders", notYetImplemented).Methods("DELETE")

	//[Route("/Library/VirtualFolders/Name", "POST")]
	libraryRouter.HandleFunc("/VirtualFolders/Name", notYetImplemented).Methods("POST")

	//[Route("/Library/VirtualFolders/Paths", "POST")]
	libraryRouter.HandleFunc("/VirtualFolders/Paths", notYetImplemented).Methods("POST")

	//[Route("/Library/VirtualFolders/Paths", "DELETE")]
	libraryRouter.HandleFunc("/VirtualFolders/Paths", notYetImplemented).Methods("DELETE")
}
