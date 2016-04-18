package gomediacenter

import "time"

type UserItemData struct {
	PlayedPercentage float32 `json:"PlayedPercentage" bson:"percent"`
	PlaybackPosTicks int `json:"PlaybackPositionTicks" bson:"pos"`
	PlayCount        int `json:"PlayCount" bson:"count"`
	Favorite         bool  `json:"IsFavorite" bson:"favorite"`
	LastPlayedDate   time.Time `json:"LastPlayedDate" bson:"lastPlayed"`
	Played           bool `json:"Played" bson:"played"`
	Key              int `json:"Key" bson:"key"`
}
