package main

import (
	"flag"
	"log"
	"time"

	"github.com/stretchr/graceful"
	"github.com/tcm1911/gomediacenter/mediaserver/controllers/auth"
	"github.com/tcm1911/gomediacenter/mediaserver/routes"
	"github.com/tcm1911/gomediacenter/mongo"
)

func main() {
	// Command line flags.
	dbAdr := flag.String("dbaddr", "localhost:27017", "Address to the database.")
	addr := flag.String("addr", ":8093", "API address.")
	//verbose := flag.Bool("verbose", false, "Verbose logging.")
	flag.Parse()

	// Connect to the database.
	log.Println("Connecting to the database...")
	db := &mongo.DB{}
	db.Connect(*dbAdr)
	log.Println("Connected to the database.")
	defer db.Close()

	log.Println("Starting sessions manager...")
	shutdownSession := auth.Run(db)

	// API server.
	mux := routes.NewAPIRouter(db, db)
	log.Println("Serving API on", *addr)
	graceful.Run(*addr, 3*time.Second, mux)

	log.Println("Stopping...")
	shutdownSession <- struct{}{}
	<-shutdownSession
}
