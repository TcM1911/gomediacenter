package gomediacenter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-errors/errors"
	"gopkg.in/mgo.v2/bson"
)

/////////////
// Structs //
/////////////

// The user struct holds all the information about a user.
type User struct {
	Name                string        `json:"Name"`
	Id                  bson.ObjectId `json:"id" bson:"_id"`
	HasPasswd           bool          `json:"HasPassword" bson:"haspassword"`
	HasConfiguredPasswd bool          `json:"HasConfiguredPassword"`
	HasConfigEasyPasswd bool          `json:"HasConfiguredEasyPassword"`
	Password            []byte        `json:"-" bson:"password"`
	Configuration       *UserConfig   `json:"Configuration"`
	Policy              *UserPolicy   `json:"Policy"`
}

// UserConfig holds the user's preferences. These can be changed by the user.
type UserConfig struct {
	PlayDefaultAudioTrack  bool `json:"PlayDefaultAudioTrack"`
	DisplayMissingEpisodes bool `json:"DisplayMissingEpisodes"`
	DisplayUnairedEpisodes bool `json:"DisplayUnairedEpisodes"`
	GroupMoviesIntoBoxSets bool `json:"GroupMoviesIntoBoxSets"`
	//"GroupedFolders":[],
	SubtitleMode           string `json:"SubtitleMode"`
	DisplayCollectionsView bool   `json:"DisplayCollectionsView"`
	DisplayFoldersView     bool   `json:"DisplayFoldersView"`
	EnableLocalPassword    bool   `json:"EnableLocalPassword"`
	//"OrderedViews":[],
	IncludeTrailersSuggestions bool `json:"IncludeTrailersInSuggestions"`
	//"LatestItemsExcludes":[],
	//"PlainFolderViews":[],
	HidePlayedInLatest         bool `json:"HidePlayedInLatest"`
	EnableChannelView          bool `json:"EnableChannelView"`
	RememberAudioSelections    bool `json:"RememberAudioSelections"`
	RememberSubtitleSelections bool `json:"RememberSubtitleSelections"`
	EnableAutoPlay             bool `json:"EnableNextEpisodeAutoPlay"`
}

// UserPolicy holds permission configurations for the user. This data can be changed
// by an admin.
type UserPolicy struct {
	Admin                bool          `json:"IsAdministrator"`
	Hidden               bool          `json:"IsHidden" bson:"IsHidden"`
	Disabled             bool          `json:"IsDisabled" bson:"IsDisabled"`
	BlockedTags          []interface{} `json:"BlockedTags,array"`
	UserPreferenceAccess bool          `json:"EnableUserPreferenceAccess"`
	//"AccessSchedules":[],
	BlockedUnratedItems     []bson.ObjectId `json:"BlockUnratedItems,array"`
	RemoteControlOtherUsers bool            `json:"EnableRemoteControlOfOtherUsers"`
	SharedDeviceControl     bool            `json:"EnableSharedDeviceControl"`
	//"EnableLiveTvManagement":true,
	//"EnableLiveTvAccess":true,
	MediaPlayback            bool `json:"EnableMediaPlayback"`
	AudioPlaybackTranscoding bool `json:"EnableAudioPlaybackTranscoding"`
	VideoPlaybackTranscoding bool `json:"EnableVideoPlaybackTranscoding"`
	ContentDeletion          bool `json:"EnableContentDeletion"`
	ContentDownloading       bool `json:"EnableContentDownloading"`
	Sync                     bool `json:"EnableSync"`
	SyncTranscoding          bool `json:"EnableSyncTranscoding"`
	//"EnabledDevices":[],
	//"EnableAllDevices":true,
	//"EnabledChannels":[],
	//"EnableAllChannels":true,
	//"EnabledFolders":[],
	//"EnableAllFolders":true,
	InvalidLoginAttempts int  `json:"InvalidLoginAttemptCount"`
	PublicSharing        bool `json:"EnablePublicSharing"`
}

// UserItemData holds data for an item with regards to a user. For example:
// how many times the item has been played, if it's a favorite.
type ItemUserData struct {
	Id               string    `json:"-" bson:"id"`
	Uid              string    `json:"-" bson:"uid"`
	PlayedPercentage float32   `json:"PlayedPercentage" bson:"percent"`
	PlaybackPosTicks int       `json:"PlaybackPositionTicks" bson:"pos"`
	PlayCount        int       `json:"PlayCount" bson:"count"`
	Favorite         bool      `json:"IsFavorite" bson:"favorite"`
	LastPlayedDate   time.Time `json:"LastPlayedDate" bson:"lastPlayed"`
	Played           bool      `json:"Played" bson:"played"`
	Key              int       `json:"Key" bson:"key"`
}

////////////
// Public //
////////////

// NewUser creates a new User instance with the default configuration.
func NewUser(name string) *User {

	config := &UserConfig{
		PlayDefaultAudioTrack:      true,
		IncludeTrailersSuggestions: true,
		HidePlayedInLatest:         true,
		RememberAudioSelections:    true,
		RememberSubtitleSelections: true,
		EnableAutoPlay:             true,
		SubtitleMode:               "Default",
	}

	policy := &UserPolicy{
		UserPreferenceAccess:     true,
		SharedDeviceControl:      true,
		MediaPlayback:            true,
		AudioPlaybackTranscoding: true,
		VideoPlaybackTranscoding: true,
		ContentDownloading:       true,
		Sync:                     true,
		SyncTranscoding:          true,
		PublicSharing:            true,
	}

	id := bson.NewObjectId()
	return &User{
		Name:          name,
		Id:            id,
		Configuration: config,
		Policy:        policy,
	}
}

// NewItemUserData returns a default ItemUserData.
func NewItemUserData(id, uid string) *ItemUserData {
	itemUserData := new(ItemUserData)
	itemUserData.Id = id
	itemUserData.Uid = uid
	return itemUserData
}

// AuthenticateUserByNameReqest creates and sends a login request to the API.
func AuthenticateUserByNameReqest(name, passwd, apiURL, authHeader string) (*AuthUserResponse, int, error) {
	reqBody := LoginRequest{Name: name, Password: passwd}
	url := apiURL + "/Users/AuthenticateByName"
	return sendAuthenticationRequest(reqBody, url, authHeader)
}

// AuthenticateUserByIdReqest creates and sends a login request to the API.
func AuthenticateUserByIdReqest(id, passwd, apiURL, authHeader string) (*AuthUserResponse, int, error) {
	reqBody := LoginRequest{Name: "", Password: passwd}
	url := apiURL + "/Users/" + id + "/Authenticate"
	return sendAuthenticationRequest(reqBody, url, authHeader)
}

func sendAuthenticationRequest(body LoginRequest, url, header string) (*AuthUserResponse, int, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, 0, errors.New("error when creating login request: " + err.Error())
	}
	req.Header.Add(SESSIION_AUTH_HEADER, header)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()

	var decoded AuthUserResponse
	return &decoded, resp.StatusCode, json.NewDecoder(resp.Body).Decode(&decoded)
}

// LogoutUserReq sends a logout request to the api server. It returns true if the request
// was successful.
func LogoutUserReq(uid, token, apiServer string) (bool, error) {

	if uid == "" || token == "" || apiServer == "" {
		return false, errors.New("arguments can't be empty")
	}
	url := fmt.Sprintf("%s/Users/%s/Logout", apiServer, uid)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add(SESSION_KEY_HEADER, token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, errors.New(
		fmt.Sprintf("logout request failed with status code: %d", resp.StatusCode))
}

// GetUser gets the user data from the api server.
func GetUser(uid, token, apiServer string) (*User, int, error) {
	if uid == "" || token == "" || apiServer == "" {
		return nil, 0, errors.New("arguments can't be empty")
	}
	url := fmt.Sprintf("%s/Users/%s", apiServer, uid)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}
	setHeader(req, token)

	resp, err := http.DefaultClient.Do(req)
	code := resp.StatusCode
	if err != nil {
		return nil, code, err
	}
	defer resp.Body.Close()

	if code != http.StatusOK {
		return nil, code, errors.New(
			fmt.Sprintf("Request failed with return code: %d\n", resp.StatusCode))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return &user, 0, err
	}
	return &user, code, err
}

func CreateUser(name, token, apiServer string) (*User, error) {
	if name == "" {
		return nil, errors.New("empty username")
	}
	url := apiServer + "/Users/New"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(NewUserRequest{Name: name})
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, err
	}
	setHeader(req, token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request returned wrong status code")
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func ChangePassword(current, new, token, uid, apiServer string) (int, error) {
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(PasswordRequest{New: new, Current: current})

	url := fmt.Sprintf("%s/Users/%s/Password", apiServer, uid)
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return 0, err
	}
	setHeader(req, token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func DeleteUser(uid, token, apiServer string) (int, error) {
	url := fmt.Sprintf("%s/Users/%s", apiServer, uid)
	r, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return 0, err
	}
	setHeader(r, token)
	resp, err := http.DefaultClient.Do(r)
	return resp.StatusCode, err
}

func setHeader(r *http.Request, token string) {
	r.Header.Add(SESSION_KEY_HEADER, token)
	r.Header.Add("Content-Type", "application/json")
}
