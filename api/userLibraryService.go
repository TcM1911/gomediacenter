package api

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"labix.org/v2/mgo"
)

/////////////
// Structs //
/////////////

type movieResponse struct {
	Movie        *gomediacenter.Movie
	ItemUserData *gomediacenter.ItemUserData `json:"UserData"`
}

type introResponse struct {
	Intros *[]gomediacenter.Intro `json:"Items,array,omitempty"`
	Count  int                    `json:"TotalRecordCount"`
}

////////////
// Public //
////////////

// UserItemHandler gets an item from a user's library.
// Path vars are uid and id.
func UserItemHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	// TODO: Add user restriction. Need to check if the user is allowed to view this item.
	uid := pathVars["uid"]
	id := pathVars["id"]

	database := GetContextVar(r, "db").(db.ItemFinder)
	if database == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no database found"))
	}

	mediaType, media, err := database.FindItemById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error find the item"))
	}

	// Get the items user data.
	itemUserData, err := database.FindItemUserData(uid, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error finding user data for the item"))
	}

	// Cast media to the right type so we can write the correct response.
	if mediaType == gomediacenter.MOVIE {
		movie := media.(*gomediacenter.Movie)
		writeMovieResponse(w, movie, itemUserData)
	}

}

// UserItemIntrosHandler returns a list of intros to play before the main media item plays.
func UserItemIntrosHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := GetContextVar(r, "pathVars").(map[string]string)
	// TODO: Add user restriction. Need to check if the user is allowed to view this item.
	//uid := pathVars["uid"]
	id := pathVars["id"]

	database := GetContextVar(r, "db").(db.IntroFinder)
	if database == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no database found"))
	}

	res := new(introResponse)

	intros, err := database.FindItemIntro(id)
	if err == mgo.ErrNotFound {
		res.Count = 0 // Ensure it's 0
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting intros from the database", err)
	} else {
		size := len(*intros)
		res.Count = size
		res.Intros = intros
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

/////////////
// Private //
/////////////

func writeMovieResponse(w http.ResponseWriter, m *gomediacenter.Movie, u *gomediacenter.ItemUserData) {
	res := movieResponse{Movie: m, ItemUserData: u}
	json.NewEncoder(w).Encode(res)
}

//[Route("/Users/{UserId}/Items/Root", "GET", Summary = "Gets the root folder from a user's library")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Users/{UserId}/FavoriteItems/{Id}", "POST", Summary = "Marks an item as a favorite")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]

//[Route("/Users/{UserId}/FavoriteItems/{Id}", "DELETE", Summary = "Unmarks an item as a favorite")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Users/{UserId}/Items/{Id}/Rating", "DELETE", Summary = "Deletes a user's saved personal rating for an item")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "DELETE")]

//[Route("/Users/{UserId}/Items/{Id}/Rating", "POST", Summary = "Updates a user's rating for an item")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "POST")]
//[ApiMember(Name = "Likes", Description = "Whether the user likes the item or not. true/false", IsRequired = true, DataType = "boolean", ParameterType = "query", Verb = "POST")]

//[Route("/Users/{UserId}/Items/{Id}/LocalTrailers", "GET", Summary = "Gets local trailers for an item")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "Id", Description = "Item Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

//[Route("/Users/{UserId}/Items/{Id}/SpecialFeatures", "GET", Summary = "Gets special features for an item")]
//[ApiMember(Name = "UserId", Description = "User Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]
//[ApiMember(Name = "Id", Description = "Movie Id", IsRequired = true, DataType = "string", ParameterType = "path", Verb = "GET")]

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
