package gomediacenter

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/////////////
// Structs //
/////////////

// The user struct holds all the information about a user.
type User struct {
	Name                string        `json:"Name"`
	Id                  bson.ObjectId `json:"id" bson:"_id"`
	HasPasswd           bool          `json:"HasPassword"`
	HasConfiguredPasswd bool          `json:"HasConfiguredPassword"`
	HasConfigEasyPasswd bool          `json:"HasConfiguredEasyPassword"`
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
	Hidden               bool          `json:"IsHidden"`
	Disabled             bool          `json:"IsDisabled"`
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
