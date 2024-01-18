package main

import (
	"log"
	"net/http"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v3/server"
)

var port = ":5000"

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return 123
}

func main() {
	server := &srv.PlayerServer{&InMemoryPlayerStore{}}
	log.Println("Connection established ont the port: ", port)
	log.Fatal((http.ListenAndServe(port, server)))
}

// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#run-the-application
