package gomediacenter

import "time"

// UserItemData holds data for an item with regards to a user. For example:
// how many times the item has been played, if it's a favorite.
type UserItemData struct {
	PlayedPercentage float32 `json:"PlayedPercentage" bson:"percent"`
	PlaybackPosTicks int `json:"PlaybackPositionTicks" bson:"pos"`
	PlayCount        int `json:"PlayCount" bson:"count"`
	Favorite         bool  `json:"IsFavorite" bson:"favorite"`
	LastPlayedDate   time.Time `json:"LastPlayedDate" bson:"lastPlayed"`
	Played           bool `json:"Played" bson:"played"`
	Key              int `json:"Key" bson:"key"`
}
