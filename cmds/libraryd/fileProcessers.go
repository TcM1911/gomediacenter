package main

import (
	"log"

	"github.com/tcm1911/gomediacenter/db"
	"github.com/tcm1911/gomediacenter/library"
)

func movieFileProcessor(job fileJob) {
	/*

	 */

	parentId := job.library.Id

	relativePath, err := library.GetRelativePath(job.file, job.libPath)
	if err != nil {
		log.Println("Failed to get the relative path, aborting", err)
		return
	}

	// Check if the file has previously been process and is in the database.
	// If it is, skip the file.
	inDB, _ := db.IsMovieFileAlreadyKnown(relativePath, parentId)
	if inDB {
		return
	}

	movieName, year := library.ParseMovieInfo(relativePath)
	metadataFetcher := new(library.MetadataDownloader)

	go func() {
		movie, err := library.DownloadIMDBMetadata(movieName, year, *metadataFetcher)
	}()
	go func() {
		output, err := library.ParseFFprobeOutput(job.file)
	}()

	// Get results

	// Add to db.
}
