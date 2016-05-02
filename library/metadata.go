package library

import (
	"net/http"

	"github.com/StalkR/imdb"
)

type IMDBMetadataDownloader interface {
	SearchIMDBTitle(c *http.Client, title string) ([]imdb.Title, error)
	NewIMDBTitle(c *http.Client, id string) (*imdb.Title, error)
}

type MetadataDownloader struct {
}

func (d *MetadataDownloader) SearchIMDBTitle(c *http.Client, title string) ([]imdb.Title, error) {
	return imdb.SearchTitle(c, title)
}

func (d *MetadataDownloader) NewIMDBTitle(c *http.Client, id string) (*imdb.Title, error) {
	return imdb.NewTitle(c, id)
}
