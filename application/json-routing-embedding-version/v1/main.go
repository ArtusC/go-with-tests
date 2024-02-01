package main

import (
	"log"
	"net/http"

	srv "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v1/server"
	store "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v1/store"
)

var port = ":5000"

func main() {
	server := srv.NewPlayerServer(store.NewInMemoryPlayerStore())
	log.Println("Connection established ont the port: ", port)
	log.Fatal((http.ListenAndServe(port, server.Handler)))
}
