package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func notYetImplemented(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write([]byte("Not yet implemented"))
}

// NewAPIRouter returns a new api router.
func NewAPIRouter() *(mux.Router) {
	router := mux.NewRouter()

	/*
		/Albums router
	*/
	albumsRouter := router.PathPrefix("/Albums").Subrouter()

	//[Route("/Albums/{Id}/Similar", "GET", Summary = "Finds albums similar to a given album.")]
	albumsRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Albums/{Id}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given album")]
	albumsRouter.HandleFunc("/{id}/InstantMix", notYetImplemented).Methods("GET")
	/*
		/Artists router
	*/
	artistsRouter := router.PathPrefix("/Artists").Subrouter()

	//[Route("/Artists/{Id}/Similar", "GET", Summary = "Finds albums similar to a given album.")]
	artistsRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Artists/{Name}/Images/{Type}", "GET")]
	//[Route("/Artists/{Name}/Images/{Type}", "HEAD")]
	artistsRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Artists/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Artists/{Name}/Images/{Type}/{Index}", "HEAD")]
	artistsRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Artists/{Name}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given artist")]
	//[ApiMember(Name = "Name", Description = "The artist name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	artistsRouter.HandleFunc("/{name}/InstantMix", notYetImplemented).Methods("GET")

	//[Route("/Artists/InstantMix", "GET", Summary = "Creates an instant playlist based on a given artist")]
	//[ApiMember(Name = "Id", Description = "The artist Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	artistsRouter.HandleFunc("/InstantMix", notYetImplemented).Methods("GET")

	/*
		/Audio
	*/
	audioRouter := router.PathPrefix("/Audio").Subrouter()

	//[Route("/Audio/{Id}/stream.{Container}", "GET", Summary = "Gets an audio stream")]
	//[Route("/Audio/{Id}/stream", "GET", Summary = "Gets an audio stream")]
	//[Route("/Audio/{Id}/stream.{Container}", "HEAD", Summary = "Gets an audio stream")]
	//[Route("/Audio/{Id}/stream", "HEAD", Summary = "Gets an audio stream")]
	//[ApiMember(Name = "Container", Description = "Container", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	audioRouter.HandleFunc("/{id}/stream", notYetImplemented).Methods("GET")

	//[Route("/Audio/{Id}/master.m3u8", "GET", Summary = "Gets an audio stream using HTTP live streaming.")]
	//[Route("/Audio/{Id}/master.m3u8", "HEAD", Summary = "Gets an audio stream using HTTP live streaming.")]
	//[Route("/Audio/{Id}/main.m3u8", "GET", Summary = "Gets an audio stream using HTTP live streaming.")]
	audioRouter.HandleFunc("/{id}/{stream}.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Audio/{Id}/hls/{SegmentId}/stream.mp3", "GET")]
	//[Route("/Audio/{Id}/hls/{SegmentId}/stream.aac", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	audioRouter.HandleFunc("/{id}/hls/{segmentId}/{stream}", notYetImplemented).Methods("GET")

	//[Route("/Audio/{Id}/hls1/{PlaylistId}/{SegmentId}.aac", "GET")]
	//[Route("/Audio/{Id}/hls1/{PlaylistId}/{SegmentId}.ts", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	audioRouter.HandleFunc("/{id}/hls1/{playlistId}/{segmentId}", notYetImplemented).Methods("GET")

	/*
	   /Auth
	*/
	authRouter := router.PathPrefix("/Auth").Subrouter()

	//[Route("/Auth/Keys", "GET")]
	//[Authenticated(Roles = "Admin")]
	authRouter.HandleFunc("/Keys", notYetImplemented).Methods("GET")

	//[Route("/Auth/Keys/{Key}", "DELETE")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Key", Description = "Auth Key", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	authRouter.HandleFunc("/Keys/{key}", notYetImplemented).Methods("DELETE")

	//[Route("/Auth/Keys", "POST")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "App", Description = "App", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	authRouter.HandleFunc("/Keys", notYetImplemented).Methods("POST")

	/*
		/Branding
	*/
	brandingRouter := router.PathPrefix("/Branding").Subrouter()

	//[Route("/Branding/Configuration", "GET", Summary = "Gets branding configuration")]
	brandingRouter.HandleFunc("/Configuration", notYetImplemented).Methods("GET")

	//[Route("/Branding/Css", "GET", Summary = "Gets custom css")]
	brandingRouter.HandleFunc("/Css", notYetImplemented).Methods("GET")

	/*
		/Channels
	*/
	channelRouter := router.PathPrefix("/Channels").Subrouter()

	//[Route("/Channels", "GET", Summary = "Gets available channels")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "SupportsLatestItems", Description = "Optional. Filter by channels that support getting latest items.", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	// channelRouter.HandleFunc("*", notYetImplemented).Methods("GET")
	router.HandleFunc("/Channels", notYetImplemented).Methods("GET")

	//[Route("/Channels/Features", "GET", Summary = "Gets features for a channel")]
	channelRouter.HandleFunc("/Features", notYetImplemented).Methods("GET")

	//[Route("/Channels/Folder", "GET", Summary = "Gets the users channel folder, along with configured images")]
	//[ApiMember(Name = "UserId", Description = "Optional attach user data.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	channelRouter.HandleFunc("/Folder", notYetImplemented).Methods("GET")

	//[Route("/Channels/{Id}/Features", "GET", Summary = "Gets features for a channel")]
	//[ApiMember(Name = "Id", Description = "Channel Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	channelRouter.HandleFunc("/{id}/Features", notYetImplemented).Methods("GET")

	//[Route("/Channels/{Id}/Items", "GET", Summary = "Gets channel items")]
	//[ApiMember(Name = "Id", Description = "Channel Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "FolderId", Description = "Folder Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "SortOrder", Description = "Sort Order - Ascending,Descending", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Filters", Description = "Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "SortBy", Description = "Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	channelRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("GET")

	//[Route("/Channels/Items/Latest", "GET", Summary = "Gets channel items")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Filters", Description = "Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "ChannelIds", Description = "Optional. Specify one or more channel id's, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	channelRouter.HandleFunc("/{id}/Latest", notYetImplemented).Methods("GET")

	/*
		/Collections
	*/
	collectionsRouter := router.PathPrefix("/Collections").Subrouter()

	//[Route("/Collections", "POST", Summary = "Creates a new collection")]
	//[ApiMember(Name = "IsLocked", Description = "Whether or not to lock the new collection.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Name", Description = "The name of the new collection.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ParentId", Description = "Optional - create the collection within a specific folder", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Ids", Description = "Item Ids to add to the collection", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	router.HandleFunc("/Collections", notYetImplemented).Methods("GET")

	//[Route("/Collections/{Id}/Items", "POST", Summary = "Adds items to a collection")]
	//[ApiMember(Name = "Ids", Description = "Item id, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	collectionsRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("POST")

	//[Route("/Collections/{Id}/Items", "DELETE", Summary = "Removes items from a collection")]
	//[ApiMember(Name = "Ids", Description = "Item id, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	collectionsRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("DELETE")

	/*
		/Connect
	*/
	connectRouter := router.PathPrefix("/Connect").Subrouter()

	//[Route("/Connect/Invite", "POST", Summary = "Creates a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "ConnectUsername", Description = "Connect username", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "SendingUserId", Description = "Sending User Id", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "EnabledLibraries", Description = "EnabledLibraries", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "EnabledChannels", Description = "EnabledChannels", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "EnableLiveTv", Description = "EnableLiveTv", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	connectRouter.HandleFunc("/Invite", notYetImplemented).Methods("POST")

	//[Route("/Connect/Pending", "GET", Summary = "Creates a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	connectRouter.HandleFunc("/Pending", notYetImplemented).Methods("GET")

	//[Route("/Connect/Pending", "DELETE", Summary = "Deletes a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Authorization Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	connectRouter.HandleFunc("/Pending", notYetImplemented).Methods("DELETE")

	//[Route("/Connect/Exchange", "GET", Summary = "Gets the corresponding local user from a connect user id")]
	//[Authenticated]
	//[ApiMember(Name = "ConnectUserId", Description = "ConnectUserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	connectRouter.HandleFunc("/Exchange", notYetImplemented).Methods("GET")

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

	/*
		/DisplayPreferences
	*/
	displayPrefrenceRouter := router.PathPrefix("/DisplayPreferences").Subrouter()

	//[Route("/DisplayPreferences/{DisplayPreferencesId}", "POST", Summary = "Updates a user's display preferences for an item")]
	//[ApiMember(Name = "DisplayPreferencesId", Description = "DisplayPreferences Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	displayPrefrenceRouter.HandleFunc("/{id}", notYetImplemented).Methods("POST")

	//[Route("/DisplayPreferences/{Id}", "GET", Summary = "Gets a user's display preferences for an item")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Client", Description = "Client", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	displayPrefrenceRouter.HandleFunc("/{id}", notYetImplemented).Methods("GET")

	/*
		/Dlna
	*/
	dlnaRouter := router.PathPrefix("/Dlna").Subrouter()

	//[Route("/Dlna/{UuId}/description.xml", "GET", Summary = "Gets dlna server info")]
	//[Route("/Dlna/{UuId}/description", "GET", Summary = "Gets dlna server info")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/description", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/contentdirectory/contentdirectory.xml", "GET", Summary = "Gets dlna content directory xml")]
	//[Route("/Dlna/{UuId}/contentdirectory/contentdirectory", "GET", Summary = "Gets dlna content directory xml")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/contentdirectory/contentdirectory.xml", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/connectionmanager/connectionmanager.xml", "GET", Summary = "Gets dlna connection manager xml")]
	//[Route("/Dlna/{UuId}/connectionmanager/connectionmanager", "GET", Summary = "Gets dlna connection manager xml")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/connectionmanager/connectionmanager.xml", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/mediareceiverregistrar/mediareceiverregistrar.xml", "GET", Summary = "Gets dlna mediareceiverregistrar xml")]
	//[Route("/Dlna/{UuId}/mediareceiverregistrar/mediareceiverregistrar", "GET", Summary = "Gets dlna mediareceiverregistrar xml")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/mediareceiverregistrar/mediareceiverregistrar.xml", notYetImplemented).Methods("GET")

	//[Route("/Dlna/{UuId}/contentdirectory/control", "POST", Summary = "Processes a control request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/contentdirectory/control", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/connectionmanager/control", "POST", Summary = "Processes a control request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/connectionmanager/control", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/mediareceiverregistrar/control", "POST", Summary = "Processes a control request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/mediareceiverregistrar/control", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/mediareceiverregistrar/events", Summary = "Processes an event subscription request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "SUBSCRIBE,POST")]
	dlnaRouter.HandleFunc("/{uuid}/mediareceiverregistrar/events", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/contentdirectory/events", Summary = "Processes an event subscription request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "SUBSCRIBE,POST")]
	dlnaRouter.HandleFunc("/{uuid}/contentdirectory/events", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/connectionmanager/events", Summary = "Processes an event subscription request")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "SUBSCRIBE,POST")]
	dlnaRouter.HandleFunc("/{uuid}/connectionmanager/events", notYetImplemented).Methods("POST")

	//[Route("/Dlna/{UuId}/icons/{Filename}", "GET", Summary = "Gets a server icon")]
	//[Route("/Dlna/icons/{Filename}", "GET", Summary = "Gets a server icon")]
	//[ApiMember(Name = "UuId", Description = "Server UuId", IsRequired = false, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Filename", Description = "The icon filename", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/{uuid}/icons/{filename}", notYetImplemented).Methods("GET")

	//[Route("/Dlna/ProfileInfos", "GET", Summary = "Gets a list of profiles")]
	dlnaRouter.HandleFunc("/{ProfileInfos}", notYetImplemented).Methods("GET")

	//[Route("/Dlna/Profiles/{Id}", "DELETE", Summary = "Deletes a profile")]
	//[ApiMember(Name = "Id", Description = "Profile Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	dlnaRouter.HandleFunc("/Profiles/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Dlna/Profiles/Default", "GET", Summary = "Gets the default profile")]
	dlnaRouter.HandleFunc("/Profiles/Default", notYetImplemented).Methods("GET")

	//[Route("/Dlna/Profiles/{Id}", "GET", Summary = "Gets a single profile")]
	//[ApiMember(Name = "Id", Description = "Profile Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/Profiles/{id}", notYetImplemented).Methods("GET")

	//[Route("/Dlna/Profiles/{ProfileId}", "POST", Summary = "Updates a profile")]
	//[ApiMember(Name = "ProfileId", Description = "Profile Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	dlnaRouter.HandleFunc("/Profiles/{id}", notYetImplemented).Methods("POST")

	//[Route("/Dlna/Profiles", "POST", Summary = "Creates a profile")]
	dlnaRouter.HandleFunc("/Profiles", notYetImplemented).Methods("POST")

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

	/*
		/Games
	*/
	gamesRouter := router.PathPrefix("/Games").Subrouter()

	//[Route("/Games/{Id}/Similar", "GET", Summary = "Finds games similar to a given game.")]
	gamesRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Games/SystemSummaries", "GET", Summary = "Finds games similar to a given game.")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	gamesRouter.HandleFunc("/SystemSummaries", notYetImplemented).Methods("GET")

	//[Route("/Games/PlayerIndex", "GET", Summary = "Gets an index of players (1-x) and the number of games listed under each")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	gamesRouter.HandleFunc("/PlayerIndex", notYetImplemented).Methods("GET")

	/*
		/GameGenres
	*/
	gameGenresRouter := router.PathPrefix("/GameGenres").Subrouter()

	//[Route("/GameGenres", "GET", Summary = "Gets all Game genres from a given item, folder, or the entire library")]
	router.HandleFunc("/GameGenres", notYetImplemented).Methods("GET")

	//[Route("/GameGenres/{Name}", "GET", Summary = "Gets a Game genre, by name")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	gameGenresRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")

	//[Route("/GameGenres/{Name}/Images/{Type}", "GET")]
	//[Route("/GameGenres/{Name}/Images/{Type}", "HEAD")]
	gameGenresRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/GameGenres/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/GameGenres/{Name}/Images/{Type}/{Index}", "HEAD")]
	gameGenresRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	/*
		/Genres
	*/
	genresRouter := router.PathPrefix("/Genres").Subrouter()

	//[Route("/Genres", "GET", Summary = "Gets all genres from a given item, folder, or the entire library")]
	router.HandleFunc("/Genres", notYetImplemented).Methods("GET")

	//[Route("/Genres/{Name}", "GET", Summary = "Gets a genre, by name")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	genresRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")

	//[Route("/Genres/{Name}/Images/{Type}", "GET")]
	//[Route("/Genres/{Name}/Images/{Type}", "HEAD")]
	genresRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Genres/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Genres/{Name}/Images/{Type}/{Index}", "HEAD")]
	genresRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

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

	/*
		/Items router
	*/
	itemsRouter := router.PathPrefix("/Items").Subrouter()

	//[Route("/Items/Counts", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Get counts from a specific user's library.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsFavorite", Description = "Optional. Get counts of favorite items", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/Counts", notYetImplemented).Methods("GET")

	//[Route("/Items/Filters", "GET", Summary = "Gets branding configuration")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "Specify this to localize the search to a specific item or folder. Omit to use the root", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "MediaTypes", Description = "Optional filter by MediaType. Allows multiple, comma delimited.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	itemsRouter.HandleFunc("/Filters", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}", "DELETE", Summary = "Deletes an item from the library and file system")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	itemsRouter.HandleFunc("/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Items/{Id}/Images", "GET", Summary = "Gets information about an item's images")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Images", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}", "GET")]
	//[Route("/Items/{Id}/Images/{Type}", "HEAD")]
	itemsRouter.HandleFunc("/{id}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}/{Index}", "GET")]
	//[Route("/Items/{Id}/Images/{Type}/{Index}", "HEAD")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}/{Index}/{Tag}/{Format}/{MaxWidth}/{MaxHeight}/{PercentPlayed}/{UnplayedCount}", "HEAD")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[Route("/Items/{Id}/Images/{Type}/{Index}/{Tag}/{Format}/{MaxWidth}/{MaxHeight}/{PercentPlayed}/{UnplayedCount}", "GET")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}/{tag}/{format}/{maxWidth}/{maxHeight}/{percentPlayed}/{unplayedCount}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Images/{Type}/{Index}/Index", "POST", Summary = "Updates the index for an item image")]
	//[Authenticated(Roles = "admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Type", Description = "Image Type", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Index", Description = "Image Index", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "NewIndex", Description = "The new image index", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}/Index", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Images/{Type}", "POST")]
	itemsRouter.HandleFunc("/{id}/Images/{type}", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Images/{Type}/{Index}", "POST")]
	//[Api(Description = "Posts an item image")]
	//[Authenticated(Roles = "admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/Images/{type}/{index}", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Ancestors", "GET", Summary = "Gets all parents of an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Ancestors", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/CriticReviews", "GET", Summary = "Gets critic reviews for an item")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/CriticReviews", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Download", "GET", Summary = "Downloads item media")]
	//[Authenticated(Roles = "download")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/Download", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/File", "GET", Summary = "Gets the original file of an item")]
	//[Authenticated]
	itemsRouter.HandleFunc("/{id}/File", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/Similar", "GET", Summary = "Gets similar items")]
	//[Authenticated]
	itemsRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ThemeMedia", "GET", Summary = "Gets theme videos and songs for an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ThemeMedia", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ThemeSongs", "GET", Summary = "Gets theme songs for an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ThemeSongs", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ThemeVideos", "GET", Summary = "Gets theme videos for an item")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "InheritFromParent", Description = "Determines whether or not parent items should be searched for theme media.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ThemeVideos", notYetImplemented).Methods("GET")

	//[Route("/Items/YearIndex", "GET", Summary = "Gets a year index based on an item query.")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	itemsRouter.HandleFunc("/YearIndex", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given item")]
	itemsRouter.HandleFunc("/{id}/InstantMix", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/ExternalIdInfos", "GET", Summary = "Gets external id infos for an item")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/ExternalIdInfos", notYetImplemented).Methods("GET")

	//[Route("/Items/RemoteSearch/Movie", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/Movie", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/AdultVideo", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/AdultMovie", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Series", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/Series", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Game", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/Game", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/BoxSet", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/BoxSet", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/MusicArtist", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/MusicArtist", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/MusicAlbum", "POST")]
	//[Authenticated]
	itemsRouter.HandleFunc("/RemoteSearch/MusicAlbum", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Person", "POST")]
	//[Authenticated(Roles = "Admin")]
	itemsRouter.HandleFunc("/RemoteSearch/Person", notYetImplemented).Methods("POST")

	//[Route("/Items/RemoteSearch/Image", "GET", Summary = "Gets a remote image")]
	//[ApiMember(Name = "ImageUrl", Description = "The image url", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ProviderName", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/RemoteSearch/Image", notYetImplemented).Methods("GET")

	//[Route("/Items/RemoteSearch/Apply/{Id}", "POST", Summary = "Applies search criteria to an item and refreshes metadata")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "The item id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ReplaceAllImages", Description = "Whether or not to replace all images", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	itemsRouter.HandleFunc("/RemoteSearch/Apply/{id}", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/Refresh", "POST", Summary = "Refreshes metadata for an item")]
	//[ApiMember(Name = "Recursive", Description = "Indicates if the refresh should occur recursively.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/Refresh", notYetImplemented).Methods("POST")

	//[Route("/Items/{ItemId}", "POST", Summary = "Updates an item")]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}", notYetImplemented).Methods("POST")

	//[Route("/Items/{ItemId}/MetadataEditor", "GET", Summary = "Gets metadata editor info for an item")]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/MetadataEditor", notYetImplemented).Methods("GET")

	//[Route("/Items/{ItemId}/ContentType", "POST", Summary = "Updates an item's content type")]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "ContentType", Description = "The content type of the item", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/ContentType", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/PlaybackInfo", "GET", Summary = "Gets live playback media info for an item")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/PlaybackInfo", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/PlaybackInfo", "POST", Summary = "Gets live playback media info for an item")]
	//[ApiMember(Name = "Type", Description = "The image type", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ProviderName", Description = "Optional. The image provider to use", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeAllLanguages", Description = "Optional.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/PlaybackInfo", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/RemoteImages", "GET", Summary = "Gets available remote images for an item")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteImages", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteImages/Providers", "GET", Summary = "Gets available remote image providers for an item")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Type", Description = "The image type", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ProviderName", Description = "The image provider", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageUrl", Description = "The image url", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteImages/Providers", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteImages/Download", "POST", Summary = "Downloads a remote image for an item")]
	//[Authenticated(Roles="Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteImages/Download", notYetImplemented).Methods("POST")

	//[Route("/Items/{Id}/RemoteSearch/Subtitles/{Language}", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Language", Description = "Language", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteSearch/Subtitles/{lang}", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteSearch/Subtitles/Providers", "GET")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	itemsRouter.HandleFunc("/{id}/RemoteSearch/Subtitles/Providers", notYetImplemented).Methods("GET")

	//[Route("/Items/{Id}/RemoteSearch/Subtitles/{SubtitleId}", "POST")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "SubtitleId", Description = "SubtitleId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	itemsRouter.HandleFunc("/{id}/RemoteSearch/Subtitles/{subid}", notYetImplemented).Methods("POST")

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

	/*
		/LiveStreams
	*/
	liveStreamsRoute := router.PathPrefix("/LiveStreams").Subrouter()

	//[Route("/LiveStreams/Open", "POST", Summary = "Opens a media source")]
	liveStreamsRoute.HandleFunc("/Open", notYetImplemented).Methods("POST")

	//[Route("/LiveStreams/Close", "POST", Summary = "Closes a media source")]
	//[ApiMember(Name = "LiveStreamId", Description = "LiveStreamId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	liveStreamsRoute.HandleFunc("/Close", notYetImplemented).Methods("POST")

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

	/*
		/Movies
	*/
	moviesRouter := router.PathPrefix("/Movies").Subrouter()

	//[Route("/Movies/{Id}/Similar", "GET", Summary = "Finds movies and trailers similar to a given movie.")]
	moviesRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Movies/Recommendations", "GET", Summary = "Gets movie recommendations")]
	//[ApiMember(Name = "CategoryLimit", Description = "The max number of categories to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ItemLimit", Description = "The max number of items to return per category", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "Specify this to localize the search to a specific item or folder. Omit to use the root", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	moviesRouter.HandleFunc("/Recommendations", notYetImplemented).Methods("GET")

	/*
		/MusicGenres
	*/
	musicGenresRouter := router.PathPrefix("/MusicGenres").Subrouter()

	//[Route("/MusicGenres/{Name}/Images/{Type}", "GET")]
	//[Route("/MusicGenres/{Name}/Images/{Type}", "HEAD")]
	musicGenresRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/MusicGenres/{Name}/Images/{Type}/{Index}", "HEAD")]
	musicGenresRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres/{Name}/InstantMix", "GET", Summary = "Creates an instant playlist based on a music genre")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[Route("/MusicGenres/InstantMix", "GET", Summary = "Creates an instant playlist based on a music genre")]
	//[ApiMember(Name = "Id", Description = "The genre Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	musicGenresRouter.HandleFunc("/{name}/InstantMix", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres", "GET", Summary = "Gets all music genres from a given item, folder, or the entire library")]
	router.HandleFunc("/MusicGenres", notYetImplemented).Methods("GET")

	//[Route("/MusicGenres/{Name}", "GET", Summary = "Gets a music genre, by name")]
	//[ApiMember(Name = "Name", Description = "The genre name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	musicGenresRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")

	/*
		/News
	*/
	newsRouter := router.PathPrefix("/News").Subrouter()

	//[Route("/News/Product", "GET", Summary = "Gets the latest product news.")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	newsRouter.HandleFunc("/Product", notYetImplemented).Methods("GET")

	/*
		/Notifications
	*/
	notificationsRouter := router.PathPrefix("/Notifications").Subrouter()

	//[Route("/Notifications/{UserId}", "GET", Summary = "Gets notifications")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "IsRead", Description = "An optional filter by IsRead", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	notificationsRouter.HandleFunc("/{id}", notYetImplemented).Methods("GET")

	//[Route("/Notifications/{UserId}/Summary", "GET", Summary = "Gets a notification summary for a user")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	notificationsRouter.HandleFunc("/{id}/Summary", notYetImplemented).Methods("GET")

	//[Route("/Notifications/Types", "GET", Summary = "Gets notification types")]
	notificationsRouter.HandleFunc("/Types", notYetImplemented).Methods("GET")

	//[Route("/Notifications/Services", "GET", Summary = "Gets notification types")]
	notificationsRouter.HandleFunc("/Services", notYetImplemented).Methods("GET")

	//[Route("/Notifications/Admin", "POST", Summary = "Sends a notification to all admin users")]
	//[ApiMember(Name = "Name", Description = "The notification's name", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Description", Description = "The notification's description", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ImageUrl", Description = "The notification's image url", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Url", Description = "The notification's info url", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Level", Description = "The notification level", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	notificationsRouter.HandleFunc("/Admin", notYetImplemented).Methods("POST")

	//[Route("/Notifications/{UserId}/Read", "POST", Summary = "Marks notifications as read")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Ids", Description = "A list of notification ids, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	notificationsRouter.HandleFunc("/{id}/Read", notYetImplemented).Methods("POST")

	//[Route("/Notifications/{UserId}/Unread", "POST", Summary = "Marks notifications as unread")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Ids", Description = "A list of notification ids, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	notificationsRouter.HandleFunc("/{id}/Unread", notYetImplemented).Methods("POST")

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

	/*
		/Persons
	*/
	personsRouter := router.PathPrefix("/Persons").Subrouter()

	//[Route("/Persons/{Name}/Images/{Type}", "GET")]
	//[Route("/Persons/{Name}/Images/{Type}", "HEAD")]
	personsRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Persons/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Persons/{Name}/Images/{Type}/{Index}", "HEAD")]
	personsRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Persons", "GET", Summary = "Gets all persons from a given item, folder, or the entire library")]
	router.HandleFunc("/Persons", notYetImplemented).Methods("GET")

	//[Route("/Persons/{Name}", "GET", Summary = "Gets a person, by name")]
	//[ApiMember(Name = "Name", Description = "The person name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	personsRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")

	/*
		/Playback
	*/
	playbackRouter := router.PathPrefix("/Playback").Subrouter()

	//[Route("/Playback/BitrateTest", "GET")]
	//[ApiMember(Name = "Size", Description = "Size", IsRequired = true, DataType = "int", ParameterType = "query", Verb = "GET")]
	playbackRouter.HandleFunc("/BitrateTest", notYetImplemented).Methods("GET")

	/*
		/Playlists
	*/
	playlistsRouter := router.PathPrefix("/Playlists").Subrouter()

	//[Route("/Playlists", "POST", Summary = "Creates a new playlist")]
	//[ApiMember(Name = "Name", Description = "The name of the new playlist.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Ids", Description = "Item Ids to add to the playlist", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "MediaType", Description = "The playlist media type", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	router.HandleFunc("/Playlists", notYetImplemented).Methods("POST")

	//[Route("/Playlists/{Id}/Items", "POST", Summary = "Adds items to a playlist")]
	//[ApiMember(Name = "Ids", Description = "Item id, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	playlistsRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("POST")

	//[Route("/Playlists/{Id}/Items/{ItemId}/Move/{NewIndex}", "POST", Summary = "Moves a playlist item")]
	//[ApiMember(Name = "ItemId", Description = "ItemId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "NewIndex", Description = "NewIndex", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	playlistsRouter.HandleFunc("/{id}/Items/{itemId}/Move/{newIndex}", notYetImplemented).Methods("POST")

	//[Route("/Playlists/{Id}/Items", "DELETE", Summary = "Removes items from a playlist")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "EntryIds", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	playlistsRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("DELETE")

	//[Route("/Playlists/{Id}/Items", "GET", Summary = "Gets the original items of a playlist")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	playlistsRouter.HandleFunc("/{id}/Items", notYetImplemented).Methods("GET")

	//[Route("/Playlists/{Id}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given playlist")]
	playlistsRouter.HandleFunc("/{id}/InstantMix", notYetImplemented).Methods("GET")

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

	/*
		/ScheduledTasks
	*/
	scheduledTasksRouter := router.PathPrefix("/ScheduledTasks").Subrouter()

	//[Route("/ScheduledTasks/{Id}", "GET", Summary = "Gets a scheduled task, by Id")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	scheduledTasksRouter.HandleFunc("/{id}", notYetImplemented).Methods("GET")

	//[Route("/ScheduledTasks", "GET", Summary = "Gets scheduled tasks")]
	//[ApiMember(Name = "IsHidden", Description = "Optional filter tasks that are hidden, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsEnabled", Description = "Optional filter tasks that are enabled, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	router.HandleFunc("/ScheduledTasks", notYetImplemented).Methods("GET")

	//[Route("/ScheduledTasks/Running/{Id}", "POST", Summary = "Starts a scheduled task")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	scheduledTasksRouter.HandleFunc("/Running/{id}", notYetImplemented).Methods("POST")

	//[Route("/ScheduledTasks/Running/{Id}", "DELETE", Summary = "Stops a scheduled task")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	scheduledTasksRouter.HandleFunc("/Running/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/ScheduledTasks/{Id}/Triggers", "POST", Summary = "Updates the triggers for a scheduled task")]
	//[ApiMember(Name = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	scheduledTasksRouter.HandleFunc("/{id}/Triggers", notYetImplemented).Methods("POST")

	/*
		/Search
	*/
	searchRouter := router.PathPrefix("/Search").Subrouter()

	//[Route("/Search/Hints", "GET", Summary = "Gets search hints based on a search term")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Supply a user id to search within a user's library or omit to search all.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "SearchTerm", Description = "The search term to filter on", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludePeople", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeMedia", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeGenres", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeStudios", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeArtists", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	searchRouter.HandleFunc("/Hints", notYetImplemented).Methods("GET")

	/*
		/Sessions
	*/
	sessionsRouter := router.PathPrefix("/Sessions").Subrouter()

	//[Route("/Sessions", "GET", Summary = "Gets a list of sessions")]
	//[Authenticated]
	//[ApiMember(Name = "ControllableByUserId", Description = "Optional. Filter by sessions that a given user is allowed to remote control.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "DeviceId", Description = "Optional. Filter by device id.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	router.HandleFunc("/Sessions", notYetImplemented).Methods("GET")

	//[Route("/Sessions/{Id}/Viewing", "POST", Summary = "Instructs a session to browse to an item or view")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "ItemType", Description = "The type of item to browse to.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ItemId", Description = "The Id of the item.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ItemName", Description = "The name of the item.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Viewing", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/Playing", "POST", Summary = "Instructs a session to play an item")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "ItemIds", Description = "The ids of the items to play, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	//[ApiMember(Name = "StartPositionTicks", Description = "The starting position of the first item.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "PlayCommand", Description = "The type of play command to issue (PlayNow, PlayNext, PlayLast). Clients who have not yet implemented play next and play last may play now.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Playing", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/Playing/{Command}", "POST", Summary = "Issues a playstate command to a client")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "SeekPositionTicks", Description = "The position to seek to.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Command", Description = "The command to send - stop, pause, unpause, nexttrack, previoustrack, seek, fullscreen.", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Playing/{cmd}", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/System/{Command}", "POST", Summary = "Issues a system command to a client")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Command", Description = "The command to send.", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/System/{cmd}", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/Command/{Command}", "POST", Summary = "Issues a system command to a client")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Command", Description = "The command to send.", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Command/{cmd}", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/Command", "POST", Summary = "Issues a system command to a client")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Command", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/Message", "POST", Summary = "Issues a command to a client to display a message to the user")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Text", Description = "The message text.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Header", Description = "The message header.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "TimeoutMs", Description = "The message timeout. If omitted the user will have to confirm viewing the message.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Message", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/Users/{UserId}", "POST", Summary = "Adds an additional user to a session")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "UserId", Description = "UserId Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Users/{uid}", notYetImplemented).Methods("POST")

	//[Route("/Sessions/{Id}/Users/{UserId}", "DELETE", Summary = "Removes an additional user from a session")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "UserId", Description = "UserId Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	sessionsRouter.HandleFunc("/{id}/Users/{uid}", notYetImplemented).Methods("DELETE")

	//[Route("/Sessions/Capabilities", "POST", Summary = "Updates capabilities for a device")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "PlayableMediaTypes", Description = "A list of playable media types, comma delimited. Audio, Video, Book, Game, Photo.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SupportedCommands", Description = "A list of supported remote control commands, comma delimited", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "MessageCallbackUrl", Description = "A url to post messages to, including remote control commands.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SupportsMediaControl", Description = "Determines whether media can be played remotely.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SupportsContentUploading", Description = "Determines whether camera upload is supported.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SupportsSync", Description = "Determines whether sync is supported.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SupportsPersistentIdentifier", Description = "Determines whether the device supports a unique identifier.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "POST")]
	sessionsRouter.HandleFunc("/Capabilities", notYetImplemented).Methods("POST")

	//[Route("/Sessions/Capabilities/Full", "POST", Summary = "Updates capabilities for a device")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	sessionsRouter.HandleFunc("/Capabilities/Full", notYetImplemented).Methods("POST")

	//[Route("/Sessions/Logout", "POST", Summary = "Reports that a session has ended")]
	//[Authenticated]
	sessionsRouter.HandleFunc("/Logout", notYetImplemented).Methods("POST")

	//[Route("/Sessions/Playing", "POST", Summary = "Reports playback has started within a session")]
	sessionsRouter.HandleFunc("/Playing", notYetImplemented).Methods("POST")

	//[Route("/Sessions/Playing/Progress", "POST", Summary = "Reports playback progress within a session")]
	sessionsRouter.HandleFunc("/Playing/Progress", notYetImplemented).Methods("POST")

	//[Route("/Sessions/Playing/Ping", "POST", Summary = "Pings a playback session")]
	//[ApiMember(Name = "PlaySessionId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	sessionsRouter.HandleFunc("/Playing/Ping", notYetImplemented).Methods("POST")

	//[Route("/Sessions/Playing/Stopped", "POST", Summary = "Reports playback has stopped within a session")]
	sessionsRouter.HandleFunc("/Playing/Stopped", notYetImplemented).Methods("POST")

	/*
		/Shows
	*/
	showsRouter := router.PathPrefix("/Shows").Subrouter()

	//[Route("/Shows/NextUp", "GET", Summary = "Gets a list of next up episodes")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "SeriesId", Description = "Optional. Filter by series id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "Specify this to localize the search to a specific item or folder. Omit to use the root", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	showsRouter.HandleFunc("/NextUp", notYetImplemented).Methods("GET")

	//[Route("/Shows/Upcoming", "GET", Summary = "Gets a list of upcoming episodes")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "ParentId", Description = "Specify this to localize the search to a specific item or folder. Omit to use the root", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	showsRouter.HandleFunc("/Upcoming", notYetImplemented).Methods("GET")

	//[Route("/Shows/{Id}/Similar", "GET", Summary = "Finds tv shows similar to a given one.")]
	showsRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Shows/{Id}/Episodes", "GET", Summary = "Gets episodes for a tv season")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "Id", Description = "The series id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Season", Description = "Optional filter by season number.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "SeasonId", Description = "Optional. Filter by season id", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsMissing", Description = "Optional filter by items that are missing episodes or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsVirtualUnaired", Description = "Optional filter by items that are virtual unaired episodes or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "AdjacentTo", Description = "Optional. Return items that are siblings of a supplied item.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartItemId", Description = "Optional. Skip through the list until a given item is found.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	showsRouter.HandleFunc("/{id}/Episodes", notYetImplemented).Methods("GET")

	//[Route("/Shows/{Id}/Seasons", "GET", Summary = "Gets seasons for a tv series")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "Id", Description = "The series id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsSpecialSeason", Description = "Optional. Filter by special season.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsMissing", Description = "Optional filter by items that are missing episodes or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsVirtualUnaired", Description = "Optional filter by items that are virtual unaired episodes or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "AdjacentTo", Description = "Optional. Return items that are siblings of a supplied item.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	showsRouter.HandleFunc("/{id}/Seasons", notYetImplemented).Methods("GET")

	/*
		/Social
	*/
	socialRouter := router.PathPrefix("/Social").Subrouter()

	//[Route("/Social/Shares/{Id}", "GET", Summary = "Gets a share")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/{id}", notYetImplemented).Methods("GET")

	//[Route("/Social/Shares/Public/{Id}", "GET", Summary = "Gets a share")]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/Public/{id}", notYetImplemented).Methods("GET")

	//[Route("/Social/Shares/Public/{Id}/Image", "GET", Summary = "Gets a share")]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/Public/{id}/Images", notYetImplemented).Methods("GET")

	//[Route("/Social/Shares", "POST", Summary = "Creates a share")]
	//[Authenticated]
	//[ApiMember(Name = "ItemId", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "UserId", Description = "The user id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	socialRouter.HandleFunc("/Shares", notYetImplemented).Methods("POST")

	//[Route("/Social/Shares/{Id}", "DELETE", Summary = "Deletes a share")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	socialRouter.HandleFunc("/Shares", notYetImplemented).Methods("DELETE")

	//[Route("/Social/Shares/Public/{Id}/Item", "GET", Summary = "Gets a share")]
	//[ApiMember(Name = "Id", Description = "The id of the item", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	socialRouter.HandleFunc("/Shares/Public/{id}/Item", notYetImplemented).Methods("GET")

	/*
		/Songs
	*/
	songsRouter := router.PathPrefix("/Songs").Subrouter()

	//[Route("/Songs/{Id}/InstantMix", "GET", Summary = "Creates an instant playlist based on a given song")]
	songsRouter.HandleFunc("/{id}/InstantMix", notYetImplemented).Methods("GET")

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

	/*
		/Studios
	*/
	studiosRouter := router.PathPrefix("/Studios").Subrouter()

	//[Route("/Studios/{Name}/Images/{Type}", "GET")]
	//[Route("/Studios/{Name}/Images/{Type}", "HEAD")]
	studiosRouter.HandleFunc("/{name}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Studios/{Name}/Images/{Type}/{Index}", "GET")]
	//[Route("/Studios/{Name}/Images/{Type}/{Index}", "HEAD")]
	studiosRouter.HandleFunc("/{name}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Studios", "GET", Summary = "Gets all studios from a given item, folder, or the entire library")]
	router.HandleFunc("/Studios", notYetImplemented).Methods("GET")

	//[Route("/Studios/{Name}", "GET", Summary = "Gets a studio, by name")]
	//[ApiMember(Name = "Name", Description = "The studio name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	studiosRouter.HandleFunc("/{name}", notYetImplemented).Methods("GET")

	/*
		/Sync
	*/
	syncRouter := router.PathPrefix("/Sync").Subrouter()

	//[Route("/Sync/Jobs/{Id}", "DELETE", Summary = "Cancels a sync job.")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	syncRouter.HandleFunc("/Jobs/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Sync/Jobs/{Id}", "GET", Summary = "Gets a sync job.")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	syncRouter.HandleFunc("/Jobs/{id}", notYetImplemented).Methods("GET")

	//[Route("/Sync/Jobs/{Id}", "POST", Summary = "Updates a sync job.")]
	syncRouter.HandleFunc("/Jobs/{id}", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems", "GET", Summary = "Gets sync job items.")]
	syncRouter.HandleFunc("/JobItems", notYetImplemented).Methods("GET")

	//[Route("/Sync/JobItems/{Id}/Enable", "POST", Summary = "Enables a cancelled or queued sync job item")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/Enable", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}/MarkForRemoval", "POST", Summary = "Marks a job item for removal")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/MarkForRemoval", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}/UnmarkForRemoval", "POST", Summary = "Unmarks a job item for removal")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/UnmarkForRemoval", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}", "DELETE", Summary = "Cancels a sync job item")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	syncRouter.HandleFunc("/JobItems/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Sync/{TargetId}/Items", "DELETE", Summary = "Cancels items from a sync target")]
	//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "Items")]
	//[ApiMember(Name = "ItemIds", Description = "ItemIds", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "Items")]
	syncRouter.HandleFunc("/{targetId}/Items", notYetImplemented).Methods("DELETE")

	//[Route("/Sync/Jobs", "GET", Summary = "Gets sync jobs.")]
	syncRouter.HandleFunc("/Jobs", notYetImplemented).Methods("GET")

	//[Route("/Sync/Jobs", "POST", Summary = "Gets sync jobs.")]
	syncRouter.HandleFunc("/Jobs", notYetImplemented).Methods("POST")

	//[Route("/Sync/Targets", "GET", Summary = "Gets a list of available sync targets.")]
	//[ApiMember(Name = "UserId", Description = "UserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/Targets", notYetImplemented).Methods("GET")

	//[Route("/Sync/Options", "GET", Summary = "Gets a list of available sync targets.")]
	//[ApiMember(Name = "UserId", Description = "UserId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ItemIds", Description = "ItemIds", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "ParentId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Category", Description = "Category", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/Options", notYetImplemented).Methods("GET")

	//[Route("/Sync/JobItems/{Id}/Transferred", "POST", Summary = "Reports that a sync job item has successfully been transferred.")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	syncRouter.HandleFunc("/JobItems/{id}/Transferred", notYetImplemented).Methods("POST")

	//[Route("/Sync/JobItems/{Id}/File", "GET", Summary = "Gets a sync job item file")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	syncRouter.HandleFunc("/JobItems/{id}/File", notYetImplemented).Methods("GET")

	//[Route("/Sync/JobItems/{Id}/AdditionalFiles", "GET", Summary = "Gets a sync job item file")]
	//[ApiMember(Name = "Id", Description = "Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Name", Description = "Name", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/JobItems/{id}/AdditionalFiles", notYetImplemented).Methods("GET")

	//[Route("/Sync/OfflineActions", "POST", Summary = "Reports an action that occurred while offline.")]
	syncRouter.HandleFunc("/OfflineActions", notYetImplemented).Methods("POST")

	//[Route("/Sync/Items/Ready", "GET", Summary = "Gets ready to download sync items.")]
	//[ApiMember(Name = "TargetId", Description = "TargetId", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET")]
	syncRouter.HandleFunc("/Items/Ready", notYetImplemented).Methods("GET")

	//[Route("/Sync/Data", "POST", Summary = "Syncs data between device and server")]
	syncRouter.HandleFunc("/Data", notYetImplemented).Methods("POST")

	/*
		System
	*/
	systemRouter := router.PathPrefix("/System").Subrouter()

	//[Route("/System/ActivityLog/Entries", "GET", Summary = "Gets activity log entries")]
	//[ApiMember(Name = "StartIndex", Description = "Optional. The record index to start at. All items with a lower index will be dropped from the results.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Optional. The maximum number of records to return", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "MinDate", Description = "Optional. The minimum date. Format = ISO", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	systemRouter.HandleFunc("/ActivityLog/Entries", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration", "GET", Summary = "Gets application configuration")]
	//[Authenticated]
	systemRouter.HandleFunc("/Configuration", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration", "POST", Summary = "Updates application configuration")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Configuration", notYetImplemented).Methods("POST")

	//[Route("/System/Configuration/{Key}", "GET", Summary = "Gets a named configuration")]
	//[Authenticated(AllowBeforeStartupWizard = true)]
	//[ApiMember(Name = "Key", Description = "Key", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	systemRouter.HandleFunc("/Configuration/{key}", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration/{Key}", "POST", Summary = "Updates named configuration")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Key", Description = "Key", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	systemRouter.HandleFunc("/Configuration/{key}", notYetImplemented).Methods("POST")

	//[Route("/System/Configuration/MetadataOptions/Default", "GET", Summary = "Gets a default MetadataOptions object")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Configuration/MetadataOptions/Default", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration/MetadataPlugins", "GET", Summary = "Gets all available metadata plugins")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Configuration/MetadataPlugins", notYetImplemented).Methods("GET")

	//[Route("/System/Configuration/MetadataPlugins/Autoset", "POST")]
	//[Authenticated(Roles = "Admin", AllowBeforeStartupWizard = true)]
	systemRouter.HandleFunc("/Configuration/MetadataPlugins/Autoset", notYetImplemented).Methods("POST")

	//[Route("/System/Info", "GET", Summary = "Gets information about the server")]
	//[Authenticated(EscapeParentalControl = true)]
	systemRouter.HandleFunc("/Info", notYetImplemented).Methods("GET")

	//[Route("/System/Info/Public", "GET", Summary = "Gets public information about the server")]
	systemRouter.HandleFunc("/Info/Public", notYetImplemented).Methods("GET")

	//[Route("/System/Ping", "POST")]
	systemRouter.HandleFunc("/Ping", notYetImplemented).Methods("POST")

	//[Route("/System/Restart", "POST", Summary = "Restarts the application, if needed")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Restart", notYetImplemented).Methods("POST")

	//[Route("/System/Shutdown", "POST", Summary = "Shuts down the application")]
	systemRouter.HandleFunc("/Shutdown", notYetImplemented).Methods("POST")

	//[Route("/System/Logs", "GET", Summary = "Gets a list of available server log files")]
	//[Authenticated(Roles = "Admin")]
	systemRouter.HandleFunc("/Logs", notYetImplemented).Methods("GET")

	//[Route("/System/Endpoint", "GET", Summary = "Gets information about the request endpoint")]
	//[Authenticated]
	systemRouter.HandleFunc("/Endpoint", notYetImplemented).Methods("GET")

	//[Route("/System/Logs/Log", "GET", Summary = "Gets a log file")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Name", Description = "The log file name.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	systemRouter.HandleFunc("/Logs/Log", notYetImplemented).Methods("GET")

	/*
		/Trailers
	*/
	trailerRouter := router.PathPrefix("/Trailers").Subrouter()

	//[Route("/Trailers/{Id}/Similar", "GET", Summary = "Finds movies and trailers similar to a given trailer.")]
	trailerRouter.HandleFunc("/{id}/Similar", notYetImplemented).Methods("GET")

	//[Route("/Trailers", "GET", Summary = "Finds movies and trailers similar to a given trailer.")]
	router.HandleFunc("/Trailers", notYetImplemented).Methods("GET")

	/*
		/Users
	*/
	usersRouter := router.PathPrefix("/Users").Subrouter()

	//[Route("/Users", "GET", Summary = "Gets a list of users")]
	//[Authenticated]
	//[ApiMember(Name = "IsHidden", Description = "Optional filter by IsHidden=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsDisabled", Description = "Optional filter by IsDisabled=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsGuest", Description = "Optional filter by IsGuest=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	router.HandleFunc("/Users", notYetImplemented).Methods("GET")

	//[Route("/Users/Public", "GET", Summary = "Gets a list of publicly visible users for display on a login screen.")]
	usersRouter.HandleFunc("/Public", notYetImplemented).Methods("GET")

	//[Route("/Users/{Id}", "GET", Summary = "Gets a user by Id")]
	//[Authenticated(EscapeParentalControl = true)]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}", notYetImplemented).Methods("GET")

	//[Route("/Users/{Id}/Offline", "GET", Summary = "Gets an offline user record by Id")]
	//[Authenticated]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}/Offline", notYetImplemented).Methods("GET")

	//[Route("/Users/{Id}", "DELETE", Summary = "Deletes a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	usersRouter.HandleFunc("/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{Id}/Authenticate", "POST", Summary = "Authenticates a user")]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Password", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Authenticate", notYetImplemented).Methods("POST")

	//[Route("/Users/AuthenticateByName", "POST", Summary = "Authenticates a user")]
	//[ApiMember(Name = "Username", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "Password", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	//[ApiMember(Name = "PasswordMd5", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	usersRouter.HandleFunc("/AuthenticateByName", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/Password", "POST", Summary = "Updates a user's password")]
	//[Authenticated]
	usersRouter.HandleFunc("/{id}/Password", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/EasyPassword", "POST", Summary = "Updates a user's easy password")]
	//[Authenticated]
	usersRouter.HandleFunc("/{id}/EasyPassword", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}", "POST", Summary = "Updates a user")]
	//[Authenticated]
	usersRouter.HandleFunc("/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/Policy", "POST", Summary = "Updates a user policy")]
	//[Authenticated(Roles = "admin")]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Policy", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/Configuration", "POST", Summary = "Updates a user configuration")]
	//[Authenticated]
	//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Configuration", notYetImplemented).Methods("POST")

	//[Route("/Users/New", "POST", Summary = "Creates a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Name", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/New", notYetImplemented).Methods("POST")

	//[Route("/Users/ForgotPassword", "POST", Summary = "Initiates the forgot password process for a local user")]
	//[ApiMember(Name = "EnteredUsername", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/ForgotPassword", notYetImplemented).Methods("POST")

	//[Route("/Users/ForgotPassword/Pin", "POST", Summary = "Redeems a forgot password pin")]
	//[ApiMember(Name = "Pin", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/ForgotPassword/Pin", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/Views", "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "IncludeExternalContent", Description = "Whether or not to include external views such as channels or live tv", IsRequired = true, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Views", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/SpecialViewOptions", "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}/SpecialViewOptions", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/GroupingOptions", "GET")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}/GroupingOptions", notYetImplemented).Methods("GET")

	//[Route("/Users/{Id}/Connect/Link", "POST", Summary = "Creates a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "ConnectUsername", Description = "Connect username", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{id}/Connect/Link", notYetImplemented).Methods("POST")

	//[Route("/Users/{Id}/Connect/Link", "DELETE", Summary = "Removes a Connect link for a user")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	usersRouter.HandleFunc("/{id}/Connect/Link", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{Id}/Images/{Type}", "GET")]
	//[Route("/Users/{Id}/Images/{Type}", "HEAD")]
	usersRouter.HandleFunc("/{id}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Users/{Id}/Images/{Type}/{Index}", "GET")]
	//[Route("/Users/{Id}/Images/{Type}/{Index}", "HEAD")]
	//[ApiMember(Name = "Id", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{id}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/PlayedItems/{Id}", "POST", Summary = "Marks an item as played")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "DatePlayed", Description = "The date the item was played (if any). Format = yyyyMMddHHmmss", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/PlayedItems/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/PlayedItems/{Id}", "DELETE", Summary = "Marks an item as unplayed")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	usersRouter.HandleFunc("/{uid}/PlayedItems/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{UserId}/PlayingItems/{Id}", "POST", Summary = "Reports that a user has begun playing an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "MediaSourceId", Description = "The id of the MediaSource", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "CanSeek", Description = "Indicates if the client can seek", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "QueueableMediaTypes", Description = "A list of media types that can be queued from this item, comma delimited. Audio,Video,Book,Game", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	//[ApiMember(Name = "AudioStreamIndex", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "SubtitleStreamIndex", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "PlayMethod", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "LiveStreamId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	//[ApiMember(Name = "PlaySessionId", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/PlayingItems/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/Items/{Id}", "GET", Summary = "Gets an item from a user's library")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/Items/Root", "GET", Summary = "Gets the root folder from a user's library")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/Root", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/Items/{Id}/Intros", "GET", Summary = "Gets intros to play before the main media item plays")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/Intros", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/FavoriteItems/{Id}", "POST", Summary = "Marks an item as a favorite")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/FavoriteItems/{id}", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/FavoriteItems/{Id}", "DELETE", Summary = "Unmarks an item as a favorite")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	usersRouter.HandleFunc("/{uid}/FavoriteItems/{id}", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{UserId}/Items/{Id}/Rating", "DELETE", Summary = "Deletes a user's saved personal rating for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/Rating", notYetImplemented).Methods("DELETE")

	//[Route("/Users/{UserId}/Items/{Id}/Rating", "POST", Summary = "Updates a user's rating for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
	//[ApiMember(Name = "Likes", Description = "Whether the user likes the item or not. true/false", IsRequired = true, DataType = "boolean", ParameterType = "query", Verb = "POST")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/Rating", notYetImplemented).Methods("POST")

	//[Route("/Users/{UserId}/Items/{Id}/LocalTrailers", "GET", Summary = "Gets local trailers for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/LocalTrailers", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/Items/{Id}/SpecialFeatures", "GET", Summary = "Gets special features for an item")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Movie Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/{id}/SpecialFeatures", notYetImplemented).Methods("GET")

	//[Route("/Users/{UserId}/Items/Latest", "GET", Summary = "Gets latest media")]
	//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Limit", Description = "Limit", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ParentId", Description = "Specify this to localize the search to a specific item or folder. Omit to use the root", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Fields", Description = "Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, CriticRatingSummary, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "IncludeItemTypes", Description = "Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET", AllowMultiple = true)]
	//[ApiMember(Name = "IsFolder", Description = "Filter by items that are folders, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "IsPlayed", Description = "Filter by items that are played, or not.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "GroupItems", Description = "Whether or not to group items into a parent container.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImages", Description = "Optional, include image information in output", IsRequired = false, DataType = "boolean", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "ImageTypeLimit", Description = "Optional, the max number of images to return, per image type", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EnableImageTypes", Description = "Optional. The image types to include in the output.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	usersRouter.HandleFunc("/{uid}/Items/Latest", notYetImplemented).Methods("GET")

	/*
		/Videos
	*/
	videosRouter := router.PathPrefix("/Videos").Subrouter()

	//[Route("/Videos/{Id}/master.m3u8", "GET", Summary = "Gets a video stream using HTTP live streaming.")]
	//[Route("/Videos/{Id}/master.m3u8", "HEAD", Summary = "Gets a video stream using HTTP live streaming.")]
	//[Route("/Videos/{Id}/main.m3u8", "GET", Summary = "Gets a video stream using HTTP live streaming.")]
	//[Route("/Videos/{Id}/stream.m3u8", "GET")]
	//[Api(Description = "Gets a video stream using HTTP live streaming.")]
	//[ApiMember(Name = "BaselineStreamAudioBitRate", Description = "Optional. Specify the audio bitrate for the baseline stream.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "AppendBaselineStream", Description = "Optional. Whether or not to include a baseline audio-only stream in the master playlist.", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "TimeStampOffsetMs", Description = "Optional. Alter the timestamps in the playlist by a given amount, in ms. Default is 1000.", IsRequired = false, DataType = "int", ParameterType = "query", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/{stream}.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/hls/{PlaylistId}/stream.m3u8", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	videosRouter.HandleFunc("/{id}/hls/{playlistId}/{stream}.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/ActiveEncodings", "DELETE")]
	//[Api(Description = "Stops an encoding process")]
	//[ApiMember(Name = "DeviceId", Description = "The device id of the client requesting. Used to stop encoding processes when needed.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	//[ApiMember(Name = "PlaySessionId", Description = "The play session id", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "DELETE")]
	videosRouter.HandleFunc("/ActiveEncodings", notYetImplemented).Methods("DELETE")

	//[Route("/Videos/{Id}/hls/{PlaylistId}/{SegmentId}.ts", "GET")]
	//[Api(Description = "Gets an Http live streaming segment file. Internal use only.")]
	videosRouter.HandleFunc("/{id}/hls/{playlistId}/{segmentId}.ts", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/hls1/{PlaylistId}/{SegmentId}.ts", "GET")]
	videosRouter.HandleFunc("/{id}/hls1/{playlistId}/{segmentId}.ts", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/master.mpd", "GET", Summary = "Gets a video stream using Mpeg dash.")]
	//[Route("/Videos/{Id}/master.mpd", "HEAD", Summary = "Gets a video stream using Mpeg dash.")]
	videosRouter.HandleFunc("/{id}/master.mpd", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/dash/{RepresentationId}/{SegmentId}.m4s", "GET")]
	videosRouter.HandleFunc("/{id}/dash/{repId}/{segmentId}.m4s", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/Subtitles/{Index}", "DELETE", Summary = "Deletes an external subtitle file")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	//[ApiMember(Name = "Index", Description = "The subtitle stream index", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "DELETE")]
	videosRouter.HandleFunc("/{id}/Subtitles/{index}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/Stream.{Format}", "GET", Summary = "Gets subtitles in a specified format.")]
	videosRouter.HandleFunc("/{id}/{mediaSrc}/Subtitles/{index}/{streamFormat}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/{StartPositionTicks}/Stream.{Format}", "GET", Summary = "Gets subtitles in a specified format.")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "MediaSourceId", Description = "MediaSourceId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Index", Description = "The subtitle stream index", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Format", Description = "Format", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "StartPositionTicks", Description = "StartPositionTicks", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "EndPositionTicks", Description = "EndPositionTicks", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/{mediaSrc}/Subtitles/{index}/{startPos}/{streamFormat}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/subtitles.m3u8", "GET", Summary = "Gets an HLS subtitle playlist.")]
	//[Authenticated]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "MediaSourceId", Description = "MediaSourceId", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "Index", Description = "The subtitle stream index", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "SegmentLength", Description = "The subtitle srgment length", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/{mediaSrc}/Subtitles/{index}/subtitles.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/live.m3u8", "GET")]
	//[Api(Description = "Gets a video stream using HTTP live streaming.")]
	videosRouter.HandleFunc("/{id}/live.m3u8", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/stream.mpegts", "GET")]
	//[Route("/Videos/{Id}/stream.ts", "GET")]
	//[Route("/Videos/{Id}/stream.webm", "GET")]
	//[Route("/Videos/{Id}/stream.asf", "GET")]
	//[Route("/Videos/{Id}/stream.wmv", "GET")]
	//[Route("/Videos/{Id}/stream.ogv", "GET")]
	//[Route("/Videos/{Id}/stream.mp4", "GET")]
	//[Route("/Videos/{Id}/stream.m4v", "GET")]
	//[Route("/Videos/{Id}/stream.mkv", "GET")]
	//[Route("/Videos/{Id}/stream.mpeg", "GET")]
	//[Route("/Videos/{Id}/stream.mpg", "GET")]
	//[Route("/Videos/{Id}/stream.avi", "GET")]
	//[Route("/Videos/{Id}/stream.m2ts", "GET")]
	//[Route("/Videos/{Id}/stream.3gp", "GET")]
	//[Route("/Videos/{Id}/stream.wmv", "GET")]
	//[Route("/Videos/{Id}/stream.wtv", "GET")]
	//[Route("/Videos/{Id}/stream.mov", "GET")]
	//[Route("/Videos/{Id}/stream.iso", "GET")]
	//[Route("/Videos/{Id}/stream", "GET")]
	//[Route("/Videos/{Id}/stream.ts", "HEAD")]
	//[Route("/Videos/{Id}/stream.webm", "HEAD")]
	//[Route("/Videos/{Id}/stream.asf", "HEAD")]
	//[Route("/Videos/{Id}/stream.wmv", "HEAD")]
	//[Route("/Videos/{Id}/stream.ogv", "HEAD")]
	//[Route("/Videos/{Id}/stream.mp4", "HEAD")]
	//[Route("/Videos/{Id}/stream.m4v", "HEAD")]
	//[Route("/Videos/{Id}/stream.mkv", "HEAD")]
	//[Route("/Videos/{Id}/stream.mpeg", "HEAD")]
	//[Route("/Videos/{Id}/stream.mpg", "HEAD")]
	//[Route("/Videos/{Id}/stream.avi", "HEAD")]
	//[Route("/Videos/{Id}/stream.3gp", "HEAD")]
	//[Route("/Videos/{Id}/stream.wmv", "HEAD")]
	//[Route("/Videos/{Id}/stream.wtv", "HEAD")]
	//[Route("/Videos/{Id}/stream.m2ts", "HEAD")]
	//[Route("/Videos/{Id}/stream.mov", "HEAD")]
	//[Route("/Videos/{Id}/stream.iso", "HEAD")]
	//[Route("/Videos/{Id}/stream", "HEAD")]
	//[Api(Description = "Gets a video stream")]
	videosRouter.HandleFunc("/{id}/{stream}", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/AdditionalParts", "GET", Summary = "Gets additional parts for a video.")]
	//[Authenticated]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	videosRouter.HandleFunc("/{id}/AdditionalParts", notYetImplemented).Methods("GET")

	//[Route("/Videos/{Id}/AlternateSources", "DELETE", Summary = "Removes alternate video sources.")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
	videosRouter.HandleFunc("/{id}/AlternateSources", notYetImplemented).Methods("DELETE")

	//[Route("/Videos/MergeVersions", "POST", Summary = "Merges videos into a single record")]
	//[Authenticated(Roles = "Admin")]
	//[ApiMember(Name = "Ids", Description = "Item id list. This allows multiple, comma delimited.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
	videosRouter.HandleFunc("/{id}/MergeVersions", notYetImplemented).Methods("POST")

	/*
		/Years
	*/
	yearsRouter := router.PathPrefix("/Years").Subrouter()

	//[Route("/Years", "GET", Summary = "Gets all years from a given item, folder, or the entire library")]
	router.HandleFunc("/Years", notYetImplemented).Methods("GET")

	//[Route("/Years/{Year}", "GET", Summary = "Gets a year")]
	//[ApiMember(Name = "Year", Description = "The year", IsRequired = true, DataType = "int", ParameterType = "path", Verb = "GET")]
	//[ApiMember(Name = "UserId", Description = "Optional. Filter by user id, and attach user data", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
	yearsRouter.HandleFunc("/{year}", notYetImplemented).Methods("GET")

	//[Route("/Years/{Year}/Images/{Type}", "GET")]
	//[Route("/Years/{Year}/Images/{Type}", "HEAD")]
	yearsRouter.HandleFunc("/{year}/Images/{type}", notYetImplemented).Methods("GET")

	//[Route("/Years/{Year}/Images/{Type}/{Index}", "GET")]
	//[Route("/Years/{Year}/Images/{Type}/{Index}", "HEAD")]
	//[ApiMember(Name = "Name", Description = "Item name", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
	yearsRouter.HandleFunc("/{year}/Images/{type}/{index}", notYetImplemented).Methods("GET")

	return router
}
