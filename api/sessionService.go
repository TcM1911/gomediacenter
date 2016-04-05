package api

//[Route("/Sessions", "GET", Summary = "Gets a list of sessions")]
//[Authenticated]
//[ApiMember(Name = "ControllableByUserId", Description = "Optional. Filter by sessions that a given user is allowed to remote control.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "DeviceId", Description = "Optional. Filter by device id.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "GET")]

//[Route("/Sessions/{Id}/Viewing", "POST", Summary = "Instructs a session to browse to an item or view")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "ItemType", Description = "The type of item to browse to.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "ItemId", Description = "The Id of the item.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "ItemName", Description = "The name of the item.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]

//[Route("/Sessions/{Id}/Playing", "POST", Summary = "Instructs a session to play an item")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "ItemIds", Description = "The ids of the items to play, comma delimited", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST", AllowMultiple = true)]
//[ApiMember(Name = "StartPositionTicks", Description = "The starting position of the first item.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "PlayCommand", Description = "The type of play command to issue (PlayNow, PlayNext, PlayLast). Clients who have not yet implemented play next and play last may play now.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]

//[Route("/Sessions/{Id}/Playing/{Command}", "POST", Summary = "Issues a playstate command to a client")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "SeekPositionTicks", Description = "The position to seek to.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "Command", Description = "The command to send - stop, pause, unpause, nexttrack, previoustrack, seek, fullscreen.", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sessions/{Id}/System/{Command}", "POST", Summary = "Issues a system command to a client")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Command", Description = "The command to send.", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sessions/{Id}/Command/{Command}", "POST", Summary = "Issues a system command to a client")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Command", Description = "The command to send.", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sessions/{Id}/Command", "POST", Summary = "Issues a system command to a client")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sessions/{Id}/Message", "POST", Summary = "Issues a command to a client to display a message to the user")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Text", Description = "The message text.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "Header", Description = "The message header.", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
//[ApiMember(Name = "TimeoutMs", Description = "The message timeout. If omitted the user will have to confirm viewing the message.", IsRequired = false, DataType = "string", ParameterType = "query", Verb = "POST")]

//[Route("/Sessions/{Id}/Users/{UserId}", "POST", Summary = "Adds an additional user to a session")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "UserId", Description = "UserId Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sessions/{Id}/Users/{UserId}", "DELETE", Summary = "Removes an additional user from a session")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "UserId", Description = "UserId Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

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

//[Route("/Sessions/Capabilities/Full", "POST", Summary = "Updates capabilities for a device")]
//[Authenticated]
//[ApiMember(Name = "Id", Description = "Session Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Sessions/Logout", "POST", Summary = "Reports that a session has ended")]
//[Authenticated]

//[Route("/Auth/Keys", "GET")]
//[Authenticated(Roles = "Admin")]

//[Route("/Auth/Keys/{Key}", "DELETE")]
//[Authenticated(Roles = "Admin")]
//[ApiMember(Name = "Key", Description = "Auth Key", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Auth/Keys", "POST")]
//[Authenticated(Roles = "Admin")]
//[ApiMember(Name = "App", Description = "App", IsRequired = true, DataType = "string", ParameterType = "query", Verb = "POST")]
