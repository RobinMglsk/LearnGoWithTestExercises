package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/robinmglsk/server"
)

const dbFileName = "game.db.json"
const addr = "0.0.0.0:5000"

func main(){
	db, err := os.OpenFile(dbFileName, os.O_RDWR | os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("problem opeing %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	server := poker.NewPlayerServer(store)
	log.Printf("starting server listing on %s", addr)

	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("could not listen on %s %v", addr, err)
	}
}