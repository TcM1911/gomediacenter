package gomediacenter

import "time"

// Movie is the data struct for movie files.
type Movie struct {
	Name     string
	SortName string
	ID       ID
	Path     string
	Video
	VideoSource
	ExtraType    string
	Awards       string
	MetaScore    int
	PremiereDate time.Time
	ImdbID       string
	MovieDB      string
	Sync
	People          []Person
	CriticalRating  int
	CriticReview    string
	Rating          string
	Overview        string
	ShortReview     string
	Taglines        []string
	Genre           []string
	CommunityRating int
	Votes           int
	PlayAccess      string
	Year            int
	PlaceHolder     bool
	Trailers        []Trailer
	Studios         []Studio
	HD              bool
	AFolder         bool
	ParentID        string
	Type            string
}

// MovieDTO is the struct that's encoded to JSON.
type MovieDTO struct {
	Name     string `json:"Name"`
	SortName string `json:"SortName"`
	ID       string `json:"Id"`
	Path     string `json:"Path"`
	Video
	VideoSource
	ExtraType    string    `json:"ExtraType,omitempty"`
	Awards       string    `json:"AwardSummary"`
	MetaScore    int       `json:"Metascore"`
	PremiereDate time.Time `json:"PremiereDate"`
	ImdbID       string    `json:"ImdbId"`
	MovieDB      string
	Sync
	People          []PersonDTO `json:"People,array"`
	CriticalRating  int         `json:"CriticRating"`
	CriticReview    string      `json:"CriticRatingSummary"`
	Rating          string      `json:"OfficialRating"`
	Overview        string      `json:"Overview"`
	ShortReview     string      `json:"ShortOverview"`
	Taglines        []string    `json:"Taglines,array"`
	Genre           []string    `json:"Genres,array"`
	CommunityRating int         `json:"CommunityRating"`
	Votes           int         `json:"VoteCount"`
	PlayAccess      string      `json:"PlayAccess"`
	Year            int         `json:"ProductionYear"`
	PlaceHolder     bool        `json:"IsPlaceHolder"`
	Trailers        []Trailer   `json:"RemoteTrailers,array"`
	Studios         []Studio    `json:"Studios"`
	HD              bool        `json:"IsHD"`
	AFolder         bool        `json:"IsFolder"`
	ParentID        string      `json:"ParentId"`
	Type            string      `json:"Type"`
}

// MovieToDTO is used to transform the movie struct to the DTO struct.
func MovieToDTO(m *Movie) (*MovieDTO, error) {
	peopleDTO := make([]PersonDTO, len(m.People))
	for i, p := range m.People {
		dto := PersonToDTO(p)
		peopleDTO[i] = dto
	}
	return &MovieDTO{
		Name:            m.Name,
		SortName:        m.SortName,
		ID:              m.ID.String(),
		Path:            m.Path,
		Video:           m.Video,
		VideoSource:     m.VideoSource,
		ExtraType:       m.ExtraType,
		Awards:          m.Awards,
		MetaScore:       m.MetaScore,
		PremiereDate:    m.PremiereDate,
		ImdbID:          m.ImdbID,
		MovieDB:         m.MovieDB,
		Sync:            m.Sync,
		People:          peopleDTO,
		CriticalRating:  m.CriticalRating,
		CriticReview:    m.CriticReview,
		Rating:          m.Rating,
		Overview:        m.Overview,
		ShortReview:     m.ShortReview,
		Taglines:        m.Taglines,
		Genre:           m.Genre,
		CommunityRating: m.CommunityRating,
		Votes:           m.Votes,
		PlayAccess:      m.PlayAccess,
		Year:            m.Year,
		PlaceHolder:     m.PlaceHolder,
		Trailers:        m.Trailers,
		Studios:         m.Studios,
		HD:              m.HD,
		AFolder:         m.AFolder,
		ParentID:        m.ParentID,
		Type:            m.Type,
	}, nil
}

// Trailer is the data struct for trailers with all its data.
type Trailer struct {
	URL  string `json:"Url"`
	Name string `json:"Name"`
	Size string `json:"VideoSize"`
}
