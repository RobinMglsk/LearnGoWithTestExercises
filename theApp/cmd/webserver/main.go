package main

import (
	"log"
	"net/http"

	poker "github.com/robinmglsk/server"
)

const dbFileName = "game.db.json"
const addr = "0.0.0.0:5000"

func main(){
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	
	server := poker.NewPlayerServer(store)
	log.Printf("starting server listing on %s", addr)

	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("could not listen on %s %v", addr, err)
	}
}