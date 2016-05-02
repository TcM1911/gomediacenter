package gomediacenter

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Video
	VideoSource
	ExtraType    string    `json:"ExtraType,omitempty"`
	Awards       string    `json:"AwardSummary"`
	MetaScore    int       `json:"Metascore"`
	PremiereDate time.Time `json:"PremiereDate"`
	ImdbId       string
	MovieDB      string
	Sync
	People          []Person      `json:"People,array"`
	CriticalRating  int           `json:"CriticRating"`
	CriticReview    string        `json:"CriticRatingSummary"`
	Rating          string        `json:"OfficialRating"`
	Overview        string        `json:"Overview"`
	ShortReview     string        `json:"ShortOverview"`
	Taglines        []string      `json:"Taglines,array"`
	Genre           []string      `json:"Genres,array"`
	CommunityRating int           `json:"CommunityRating"`
	Votes           int           `json:"VoteCount"`
	PlayAccess      string        `json:"PlayAccess"`
	Year            int           `json:"ProductionYear"`
	PlaceHolder     bool          `json:"IsPlaceHolder"`
	Trailers        []Trailer     `json:"RemoteTrailers,array"`
	Studios         []Studio      `json:"Studios"`
	HD              bool          `json:"IsHD"`
	AFolder         bool          `json:"IsFolder"`
	ParentId        bson.ObjectId `json:"ParentId"`
	Type            string        `json:"Type"`
}

type Trailer struct {
	URL  string `json:"Url"`
	Name string `json:"Name"`
	Size string `json:"VideoSize"`
}
