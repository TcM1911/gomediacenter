package gomediacenter

import "time"

/////////////
// Structs //
/////////////

// UserItemData holds data for an item with regards to a user. For example:
// how many times the item has been played, if it's a favorite.
type ItemUserData struct {
	Id               string `json:"-" bson:"id"`
	Uid              string `json:"-" bson:"uid"`
	PlayedPercentage float32 `json:"PlayedPercentage" bson:"percent"`
	PlaybackPosTicks int `json:"PlaybackPositionTicks" bson:"pos"`
	PlayCount        int `json:"PlayCount" bson:"count"`
	Favorite         bool  `json:"IsFavorite" bson:"favorite"`
	LastPlayedDate   time.Time `json:"LastPlayedDate" bson:"lastPlayed"`
	Played           bool `json:"Played" bson:"played"`
	Key              int `json:"Key" bson:"key"`
}

////////////
// Public //
////////////

func NewItemUserData(id, uid string) *ItemUserData {
	itemUserData := new(ItemUserData)
	itemUserData.Id = id
	itemUserData.Uid = uid
	return itemUserData
}

