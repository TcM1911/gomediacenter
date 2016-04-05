package api

//[Route("/Users", "GET", Summary = "Gets a list of users")]
//[Authenticated]
//[ApiMember(Name = "IsHidden", Description = "Optional filter by IsHidden=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "IsDisabled", Description = "Optional filter by IsDisabled=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]
//[ApiMember(Name = "IsGuest", Description = "Optional filter by IsGuest=true or false", IsRequired = false, DataType = "bool", ParameterType = "query", Verb = "GET")]

//[Route("/Users/Public", "GET", Summary = "Gets a list of publicly visible users for display on a login screen.")]

//[Route("/Users/{Id}", "GET", Summary = "Gets a user by Id")]
//[Authenticated(EscapeParentalControl = true)]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Users/{Id}/Offline", "GET", Summary = "Gets an offline user record by Id")]
//[Authenticated]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Users/{Id}", "DELETE", Summary = "Deletes a user")]
//[Authenticated(Roles = "Admin")]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Users/{Id}/Authenticate", "POST", Summary = "Authenticates a user")]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Password", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/AuthenticateByName", "POST", Summary = "Authenticates a user")]
//[ApiMember(Name = "Username", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
//[ApiMember(Name = "Password", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]
//[ApiMember(Name = "PasswordMd5", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/{Id}/Password", "POST", Summary = "Updates a user's password")]
//[Authenticated]

//[Route("/Users/{Id}/EasyPassword", "POST", Summary = "Updates a user's easy password")]
//[Authenticated]

//[Route("/Users/{Id}", "POST", Summary = "Updates a user")]
//[Authenticated]

//[Route("/Users/{Id}/Policy", "POST", Summary = "Updates a user policy")]
//[Authenticated(Roles = "admin")]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Users/{Id}/Configuration", "POST", Summary = "Updates a user configuration")]
//[Authenticated]
//[ApiMember(Name = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Users/New", "POST", Summary = "Creates a user")]
//[Authenticated(Roles = "Admin")]
//[ApiMember(Name = "Name", IsRequired = true, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/ForgotPassword", "POST", Summary = "Initiates the forgot password process for a local user")]
//[ApiMember(Name = "EnteredUsername", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]

//[Route("/Users/ForgotPassword/Pin", "POST", Summary = "Redeems a forgot password pin")]
//[ApiMember(Name = "Pin", IsRequired = false, DataType = "string", ParameterType = "body", Verb = "POST")]
