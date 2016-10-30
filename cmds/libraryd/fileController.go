package main

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/library"
)

// This function checks all items in the library against the files in the library folder.
// If an item exist in the database but no file was found, the item is removed.
// This is used to remove old items.
func purneLibrary(db gomediacenter.LibraryMaintainer, files map[string]struct{}, parentId gomediacenter.ID, mediaType gomediacenter.MEDIATYPE) {
	log.Println("Pruning dead items from the library.")

	// Get all items in the library and check if the file was walked.
	removedItems, err := db.PruneMissingItemsFromLibrary(files, parentId, mediaType)
	if err != nil {
		log.Println("Pruning the database failed because", err)
	}

	if *verbose {
		for _, item := range removedItems {
			log.Println(item, "pruned from the database.")
		}
	}
}

func movieFileProcessor(db gomediacenter.MovieMaintainer, job fileJob) {
	parentId := job.library.ID

	relativePath, err := library.GetRelativePath(job.file, job.libPath)
	if err != nil {
		log.Println("Failed to get the relative path, aborting", err)
		return
	}

	// Check if the file has previously been process and is in the database.
	// If it is, skip the file.
	inDB := db.IsMovieFileAlreadyKnown(relativePath, parentId)
	if inDB {
		if *verbose {
			log.Println("The file", relativePath, "has already been processed.")
		}
		return
	}

	// Get metadata from IMDB.
	movieName, year := library.ParseMovieInfo(relativePath)
	metadataFetcher := new(library.MetadataDownloader)
	metaDataChan := make(chan *gomediacenter.Movie)
	go fetchMetaDataFromIMDB(movieName, year, metadataFetcher, metaDataChan)

	// Get file information from ffprobe.
	ffprobChan := make(chan library.FFprobeOutput)
	go processFileWithFFProbe(job, ffprobChan)

	// Wait for the results from ffprobe and IMDB.
	movie := <-metaDataChan
	ffprobeOutput := <-ffprobChan

	// Add the results to a movie struct.
	movie.Chapters = library.ConvertFFprobeChapter(ffprobeOutput.Chapters)
	movie.MediaStreams = library.ConvertFFprobeStream(ffprobeOutput.Streams)
	movie.ParentID = parentId
	movie.Path = relativePath

	// Add the new movie to the db.
	movieId, err := db.InsertNewMovie(movie)
	if err != nil {
		log.Println("Error when adding", movie.Name, "to the database", err)
		return
	}
	log.Println(movie.Name, "(", movieId, ") added to the library.")
}

func fetchMetaDataFromIMDB(movieName string, year int, metadataFetcher *library.MetadataDownloader, res chan<- *gomediacenter.Movie) {
	movie, err := library.DownloadIMDBMetadata(movieName, year, *metadataFetcher)
	if err != nil {
		log.Println("Failed to get meta data from IMDB.", err)
	}
	res <- movie
	close(res)
}

func processFileWithFFProbe(job fileJob, res chan<- library.FFprobeOutput) {
	args := []string{"-v", "quiet",
		"-show_chapters",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		job.file}

	cmd := exec.Command("ffprobe", args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Println("Failed to run ffprobe", err)
		log.Println(stderr.String())
		return
	}
	output, err := library.ParseFFprobeOutput(stdout.Bytes())
	if err != nil {
		log.Println("Failed to get media info from FFprobe output", err)
	}
	res <- output
	close(res)
}
