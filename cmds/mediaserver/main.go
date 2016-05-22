package main

import (
	"flag"
	"time"

	"github.com/stretchr/graceful"
	"github.com/tcm1911/gomediacenter/api"
	"github.com/tcm1911/gomediacenter/db"
	"log"
)

func main() {
	// Command line flags.
	dbAdr := flag.String("dbaddr", "localhost:27017", "Address to the database.")
	addr := flag.String("addr", ":8093", "API address.")
	//verbose := flag.Bool("verbose", false, "Verbose logging.")
	flag.Parse()

	// Connect to the database.
	log.Println("Connecting to the database...")
	db.Connect(*dbAdr)
	log.Println("Connected to the database.")
	defer db.Disconnect()

	// API server.
	mux := api.NewAPIRouter()
	log.Println("Serving API on", *addr)
	graceful.Run(*addr, 3*time.Second, mux)
}
