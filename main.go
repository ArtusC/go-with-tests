package main

import (
	"os"

	mo "github.com/ArtusC/go-with-tests/mocking"
)

func main() {
	// Chapter dependency_injection
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))

	// Chapter mocking
	sleeper := &mo.DefaultSleeper{}
	mo.Countdown(os.Stdout, sleeper)
}
