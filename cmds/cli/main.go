package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/mongo"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		errors.New("invalid usage: must specify command")
		return
	}

	db := mongo.DB{}
	db.Connect("localhost")

	switch args[0] {
	case "newLibrary":
		name := args[1]
		var libType gomediacenter.MEDIATYPE
		switch args[2] {
		case "movie":
			libType = gomediacenter.MOVIE
		}
		library, err := db.NewLibrary(name, libType)
		if err != nil {
			panic(err)
		}
		fmt.Println(`Put the library id in a file called ".library_id" in the root folder of the library.`)
		fmt.Println("Library id:", library.ID.Hex())
	}

}
