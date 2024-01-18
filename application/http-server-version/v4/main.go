package main

import (
	"log"
	"net/http"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v4/server"
	store "github.com/ArtusC/go-with-tests/application/http-server-version/v4/store"
)

var port = ":5000"

func main() {
	server := &srv.PlayerServer{store.NewInMemoryPlayerStore()}
	log.Println("Connection established ont the port: ", port)
	log.Fatal((http.ListenAndServe(port, server)))
}
