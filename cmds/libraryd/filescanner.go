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
	medaiType gomediacenter.MEDIATYPE
	library   db.Library
	libPath   string
}

func scanFolder(libId bson.ObjectId, folder string, complete chan<- struct{}) {

	// Before we start the scan, let's make sure we have the newest config of this library
	err := fetchLibraryDataFromDB(libId)
	if err != nil {
		log.Println("Error when getting the latest library data.", err)
	}

	LIBRARY := getCopyOfLibrary(libId)

	// Create a workers pool
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

	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err, "when scanning scanning", folder)
			return err
		}
		if info.IsDir() {
			if *verbose {
				log.Println("Scanning", path)
			}
			return nil
		}

		// Only start processing another file if the daemon is still running.
		runningMutex.RLock()
		defer runningMutex.RUnlock()
		if running && library.IsVideoFile(path) {
			jobChan <- fileJob{file: path, medaiType: mediaType, libPath: folder, library: LIBRARY}
		}

		return nil
	})

	// All work has been dispatched. Tell the workers and wait for them to finish.
	if *verbose {
		log.Println("Waiting for workers to finish.")
	}
	close(jobChan)
	workers.Wait()
	log.Println("Scan of", folder, "complete.")
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
