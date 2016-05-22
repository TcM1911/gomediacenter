package routes

import "github.com/gorilla/mux"

func newBrandingRouter(router *mux.Router) {
	/*
		/Branding
	*/
	brandingRouter := router.PathPrefix("/Branding").Subrouter()

	//[Route("/Branding/Configuration", "GET", Summary = "Gets branding configuration")]
	brandingRouter.HandleFunc("/Configuration", notYetImplemented).Methods("GET")

	//[Route("/Branding/Css", "GET", Summary = "Gets custom css")]
	brandingRouter.HandleFunc("/Css", notYetImplemented).Methods("GET")
}
