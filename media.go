package gomediacenter

// MEDIATYPE is an iota for the media type.
type MEDIATYPE int8

const (
	// MOVIE is the media type used for movies.
	MOVIE MEDIATYPE = iota
	// TV is the media type used for TV shows.
	TV
	// MUSIC is the media type used for music.
	MUSIC
)
