package routes

import "github.com/gorilla/mux"

func newDevicesRouter(router *mux.Router) {
	/*
		/Devices
	*/
	devicesRouter := router.PathPrefix("/Devices").Subrouter()

	//[Route("/Devices", "GET", Summary = "Gets all devices")]
	//[Authenticated(Roles = "Admin")]
	router.HandleFunc("/Devices", notYetImplemented).Methods("GET")

	//[Route("/Devices", "DELETE", Summary = "Deletes a device")]
	//[ApiMember(Name = "Id", Description = "Device Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	router.HandleFunc("/Devices", notYetImplemented).Methods("DELETE")

	//[Route("/Devices/CameraUploads", "GET", Summary = "Gets camera upload history for a device")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Device Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	devicesRouter.HandleFunc("/CameraUploads", notYetImplemented).Methods("GET")

	//[Route("/Devices/CameraUploads", "POST", Summary = "Uploads content")]
	//[Authenticated]
	//[ApiMember(Name = "DeviceId", Description = "Device Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Album", Description = "Album", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Name", Description = "Name", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	devicesRouter.HandleFunc("/CameraUploads", notYetImplemented).Methods("POST")

	//[Route("/Devices/Info", "GET", Summary = "Gets device info")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Device Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	devicesRouter.HandleFunc("/Info", notYetImplemented).Methods("GET")

	//[Route("/Devices/Capabilities", "GET", Summary = "Gets device capabilities")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Device Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	devicesRouter.HandleFunc("/Capabilities", notYetImplemented).Methods("GET")

	//[Route("/Devices/Options", "POST", Summary = "Updates device options")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Device Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	devicesRouter.HandleFunc("/Options", notYetImplemented).Methods("POST")
}
