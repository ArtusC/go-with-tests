package main

import (
	"log"
	"net/http"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v2/server"
)

var port = ":5000"

func main() {
	handler := http.HandlerFunc(srv.PlayerServer)
	log.Println("Connection established ont the port: ", port)
	log.Fatal((http.ListenAndServe(port, handler)))
}
