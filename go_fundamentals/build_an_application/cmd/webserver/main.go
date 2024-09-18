package main

import (
	"log"
	"net/http"

	poker "github.com/Vantonage/learn_go_with_tests/go_fundamentals/build_an_application"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatal(http.ListenAndServe(":5000", server))
	}

}
