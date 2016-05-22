package routes

import "github.com/gorilla/mux"

func newImagesRouter(router *mux.Router) {
	/*
		/Images
	*/
	imagesRouter := router.PathPrefix("/Images").Subrouter()

	//[Route("/Images/General/{Name}/{Type}", "GET", Summary = "Gets a general image by name")]
	//[ApiMember(Name = "Name", Description = "The name of the image", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Type", Description = "Image Type (primary, backdrop, logo, etc).", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	imagesRouter.HandleFunc("/General/{name}/{type}", notYetImplemented).Methods("GET")

	//[Route("/Images/Ratings/{Theme}/{Name}", "GET", Summary = "Gets a rating image by name")]
	//[ApiMember(Name = "Name", Description = "The name of the image", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Theme", Description = "The theme to get the image from", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	imagesRouter.HandleFunc("/Ratings/{theme}/{name}", notYetImplemented).Methods("GET")

	//[Route("/Images/MediaInfo/{Theme}/{Name}", "GET", Summary = "Gets a media info image by name")]
	//[ApiMember(Name = "Name", Description = "The name of the image", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Theme", Description = "The theme to get the image from", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	imagesRouter.HandleFunc("/MediaInfo/{theme}/{name}", notYetImplemented).Methods("GET")

	//[Route("/Images/MediaInfo", "GET", Summary = "Gets all media info image by name")]
	//[Authenticated]
	imagesRouter.HandleFunc("/MediaInfo", notYetImplemented).Methods("GET")

	//[Route("/Images/Ratings", "GET", Summary = "Gets all rating images by name")]
	//[Authenticated]
	imagesRouter.HandleFunc("/Ratings", notYetImplemented).Methods("GET")

	//[Route("/Images/General", "GET", Summary = "Gets all general images by name")]
	//[Authenticated]
	imagesRouter.HandleFunc("/General", notYetImplemented).Methods("GET")

	//[Route("/Images/Remote", "GET", Summary = "Gets a remote image")]
	//[ApiMember(Name = "ImageUrl", Description = "The image url", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	imagesRouter.HandleFunc("/Remote", notYetImplemented).Methods("GET")
}
