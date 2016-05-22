package routes

import "github.com/gorilla/mux"

func newPackagesRouter(router *mux.Router) {
	/*
		/Packages
	*/
	packagesRouter := router.PathPrefix("/Packages").Subrouter()

	//[Route("/Packages/Reviews/{Id}", "POST", Summary = "Creates or updates a package review")]
	//[ApiMember(Name = "Id", Description = "Package Id", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Rating", Description = "The rating value (1-5)", IsRequired = true, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Recommend", Description = "Whether or not this review recommends this item", IsRequired = true, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Title", Description = "Optional short description of review.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Review", Description = "Optional full review.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	packagesRouter.HandleFunc("/Reviews/{id}", notYetImplemented).Methods("POST")

	//[Route("/Packages/{Id}/Reviews", "GET", Summary = "Gets reviews for a package")]
	//[ApiMember(Name = "Id", Description = "Package Id", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "MaxRating", Description = "Retrieve only reviews less than or equal to this", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "MinRating", Description = "Retrieve only reviews greator than or equal to this", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ForceTitle", Description = "Whether or not to restrict results to those with a title", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Limit the result to this many reviews (ordered by latest)", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	packagesRouter.HandleFunc("/Reviews/{id}", notYetImplemented).Methods("GET")
}
