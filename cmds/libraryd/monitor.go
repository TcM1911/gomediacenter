package main

import (
	"log"
	"time"

	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2/bson"
)

func monitorFolder(libId bson.ObjectId, folder string, shutdownChannel chan struct{}) {
	ticker := time.Tick(time.Duration(*interval) * time.Minute)
monitorLoop:
	for {
		select {
		case <-ticker:
			if !ActiveScans() {
				completeChan := make(chan struct{}, 1)
				log.Println("Scanning folder:", folder)
				scan := &Scan{channel: completeChan, id: time.Now().UnixNano()}
				AddScanToPoolAndWatch(scan)
				err := db.UpdateLibraryLastScannedTime(libId, time.Now())
				if err != nil {
					log.Println("Failed to update last scanned time", err)
					close(completeChan)
					return
				}
				scanFolder(libId, folder, completeChan)
			} else {
				log.Println("Another scan is already active.")
			}
		case <-shutdownChannel:
			break monitorLoop
		}
	}
	log.Println("Monitor shutdowned.")
}
