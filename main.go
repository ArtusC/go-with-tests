package main

import (
	"log"
	"net/http"

	di "github.com/ArtusC/go-with-tests/dependency-injection"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
}
