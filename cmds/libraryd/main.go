package main

import (
	"bufio"
	"flag"
	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

var dbAdr *string = flag.String("dbaddr", "localhost:27017", "Address to the database.")
var mediaFolder *string = flag.String("path", "pwd", "Path to library folder to watch.")
var rescan *bool = flag.Bool("rescan", true, "Rescan folder at startup.")
var verbose *bool = flag.Bool("verbose", false, "More logging.")
var numWorkers *int = flag.Int("workers", 10, "The number of workers used to process the media files.")

// This variable indicates if the daemon is still running or should be shutdown.
// This should be checked from time to time during long process and abort if changed to false.
var running bool = true
var runningMutex sync.RWMutex

func main() {
	flag.Parse()

	if *mediaFolder == "pwd" {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		*mediaFolder = pwd
	}

	// Connect to the database.
	log.Println("Connecting to the database.")
	db.Connect(*dbAdr)
	defer db.Disconnect()

	// Get the library information from the database.
	idFile := filepath.Join(*mediaFolder, ".library_id")
	libraryId := readLibrayIdFromFile(idFile)
	log.Println("Monitoring library:", getCopyOfLibrary(libraryId).Name)

	if *rescan == true {
		completeChan := make(chan struct{}, 1)
		log.Println("Scanning folder:", *mediaFolder)
		go func() {
			err := db.UpdateLibraryLastScannedTime(libraryId, time.Now())
			if err != nil {
				log.Println("[Error] Faild to update last scanned time", err)
			}
		}()
		go scanFolder(libraryId, *mediaFolder, completeChan)
		initialScan := &Scan{channel: completeChan, id: time.Now().UnixNano()}
		AddScanToPoolAndWatch(initialScan)
	}

	// Shutdown.
	shutdownComplete := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		log.Println("Stopping...")
		runningMutex.Lock()
		running = false
		runningMutex.Unlock()
		log.Println("Waiting for processes to shutdown...")

		// Waiting for all scans to shutdown.
		for ActiveScans() {
			time.Sleep(500 * time.Millisecond)
		}

		shutdownComplete <- struct{}{}
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownComplete
	time.Sleep(2 * time.Second)
	log.Println("Halted..")
}

func readLibrayIdFromFile(filePath string) bson.ObjectId {
	var id bson.ObjectId
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalln("Could not read the library id file", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id = bson.ObjectIdHex(scanner.Text())
		err := fetchLibraryDataFromDB(id)
		if err != nil {
			log.Panic(err)
		}
		return id
	}
	return id
}
