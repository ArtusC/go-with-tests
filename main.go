package main

import (
	"os"
	"time"

	mo "github.com/ArtusC/go-with-tests/mocking"
)

func main() {
	// Chapter dependency_injection
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))

	// Chapter mocking
	sleeper := &mo.ConfigurableSleeper{Duration: 1 * time.Second, SleeperDuration: time.Sleep}
	mo.Countdown(os.Stdout, sleeper)
}
