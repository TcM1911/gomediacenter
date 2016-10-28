package routes

import "github.com/gorilla/mux"

func newEnvironmentRouter(router *mux.Router) {
	/*
		/Environment
	*/
	environmentRouter := router.PathPrefix("/Environment").Subrouter()

	//[Route("/Environment/DirectoryContents", "GET", Summary = "Gets the contents of a given directory in the file system")]
	//[ApiMember(Name = "Path", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeFiles", Description = "An optional filter to include or exclude files from the results. true/false", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeDirectories", Description = "An optional filter to include or exclude folders from the results. true/false", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeHidden", Description = "An optional filter to include or exclude hidden files and folders. true/false", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	environmentRouter.HandleFunc("/DirectoryContents", notYetImplemented).Methods("GET")

	//[Route("/Environment/NetworkShares", "GET", Summary = "Gets shares from a network device")]
	//[ApiMember(Name = "Path", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	environmentRouter.HandleFunc("/NetworkShares", notYetImplemented).Methods("GET")

	//[Route("/Environment/Drives", "GET", Summary = "Gets available drives from the server's file system")]
	environmentRouter.HandleFunc("/Drives", notYetImplemented).Methods("GET")

	//[Route("/Environment/NetworkDevices", "GET", Summary = "Gets a list of devices on the network")]
	environmentRouter.HandleFunc("/NetworkDevices", notYetImplemented).Methods("GET")

	//[Route("/Environment/ParentPath", "GET", Summary = "Gets the parent path of a given path")]
	//[ApiMember(Name = "Path", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	environmentRouter.HandleFunc("/ParentPath", notYetImplemented).Methods("GET")
}
