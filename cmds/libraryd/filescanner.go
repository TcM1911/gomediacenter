package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"github.com/tcm1911/gomediacenter/library"
	"gopkg.in/mgo.v2/bson"
)

type fileJob struct {
	file      string
	mediaType gomediacenter.MEDIATYPE
	library   db.Library
	libPath   string
}

func scanFolder(libId bson.ObjectId, folder string, complete chan<- struct{}) {
	// Before we start the scan, let's make sure we have the newest config for this library.
	err := fetchLibraryDataFromDB(libId)
	if err != nil {
		log.Println("Error when getting the latest library data.", err)
	}
	LIBRARY := getCopyOfLibrary(libId)

	// Create a worker pool
	var workers sync.WaitGroup
	jobChan := make(chan fileJob, *numWorkers)

	if *verbose {
		log.Println("Starting up", *numWorkers, "workers.")
	}
	for i := 1; i <= *numWorkers; i++ {
		workers.Add(1)
		go scanWorker(&workers, jobChan)
	}

	// Determine the folder type: Movies, TV shows, Music, etc.
	mediaType := LIBRARY.Type

	// Walk the library folder and process each file.
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err, "when scanning scanning", folder)
			return err
		}
		// If we find a folder, log if we are being verbose before continue with the
		// next file.
		if info.IsDir() {
			if *verbose {
				log.Println("Scanning", path)
			}
			return nil
		}

		// If we have a file, only start processing
		// if the daemon is still running and not in shutdown mode.
		runningMutex.RLock()
		defer runningMutex.RUnlock()
		if running && library.IsVideoFile(path) {
			// Wait for a worker to become available and send the job.
			jobChan <- fileJob{file: path, mediaType: mediaType,
				libPath: folder, library: LIBRARY}
		}
		return nil
	})

	// All work has been dispatched. Tell the workers and wait for them to finish.
	if *verbose {
		log.Println("Waiting for workers to finish processing files.")
	}
	close(jobChan)
	workers.Wait()
	log.Println("Scanning for new items in", folder, "complete.")
	// TODO: Start purge scan to remove removed files from the library.
	complete <- struct{}{}
}

func scanWorker(pool *sync.WaitGroup, jobs <-chan fileJob) {
	if *verbose {
		log.Println("Worker started.")
	}
	defer pool.Done()

	for job := range jobs {
		if *verbose {
			log.Println("Processing", job.file)
		}
		// Use the movieFileProcessor if we are processing files in a movie library.
		if job.library.Type == gomediacenter.MOVIE {
			movieFileProcessor(job)
		} else {
			processFile(job.file)
		}
	}
}

func processFile(path string) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(path)
}
