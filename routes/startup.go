package routes

import "github.com/gorilla/mux"

func newStartupRouter(router *mux.Router) {
	/*
		/Startup
	*/
	startupRouter := router.PathPrefix("/Startup").Subrouter()

	//[Route("/Startup/Complete", "POST", Summary = "Reports that the startup wizard has been completed")]
	startupRouter.HandleFunc("/Complete", notYetImplemented).Methods("POST")

	//[Route("/Startup/Info", "GET", Summary = "Gets initial server info")]
	startupRouter.HandleFunc("/Complete", notYetImplemented).Methods("GET")

	//[Route("/Startup/Configuration", "GET", Summary = "Gets initial server configuration")]
	startupRouter.HandleFunc("/Configuration", notYetImplemented).Methods("GET")

	//[Route("/Startup/Configuration", "POST", Summary = "Updates initial server configuration")]
	startupRouter.HandleFunc("/Configuration", notYetImplemented).Methods("POST")

	//[Route("/Startup/User", "GET", Summary = "Gets initial user info")]
	startupRouter.HandleFunc("/User", notYetImplemented).Methods("GET")

	//[Route("/Startup/User", "POST", Summary = "Updates initial user info")]
	startupRouter.HandleFunc("/User", notYetImplemented).Methods("POST")
}
