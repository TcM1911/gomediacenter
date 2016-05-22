package routes

import "github.com/gorilla/mux"

func newProvidersRouter(router *mux.Router) {
	/*
		/Providers
	*/
	providersRouter := router.PathPrefix("/Providers").Subrouter()

	//[Route("/Providers/Chapters", "GET")]
	//[Authenticated]
	providersRouter.HandleFunc("/Chapters", notYetImplemented).Methods("GET")

	//[Route("/Providers/Subtitles/Subtitles/{Id}", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	providersRouter.HandleFunc("/Subtitles/Subtitles/{id}", notYetImplemented).Methods("GET")
}
