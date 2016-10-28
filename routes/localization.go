package routes

import "github.com/gorilla/mux"

func newLocalizationRouter(router *mux.Router) {
	/*
		/Localization
	*/
	localizationRouter := router.PathPrefix("/Localization").Subrouter()

	//[Route("/Localization/Cultures", "GET", Summary = "Gets known cultures")]
	localizationRouter.HandleFunc("/Cultures", notYetImplemented).Methods("GET")

	//[Route("/Localization/Countries", "GET", Summary = "Gets known countries")]
	localizationRouter.HandleFunc("/Countries", notYetImplemented).Methods("GET")

	//[Route("/Localization/ParentalRatings", "GET", Summary = "Gets known parental ratings")]
	localizationRouter.HandleFunc("/ParentalRatings", notYetImplemented).Methods("GET")

	//[Route("/Localization/Options", "GET", Summary = "Gets localization options")]
	localizationRouter.HandleFunc("/Options", notYetImplemented).Methods("GET")
}
