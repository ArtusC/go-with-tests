package concurrency

import "time"

type WebsiteChecker func(string) bool

/*
	Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time.
	Maps in Go don't like it when more than one thing tries to write to them at once, and so fatal error.
	This is a race condition, a bug that occurs when the output of our software is dependent on the timing
	and sequence of events that we have no control over.
	Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable
	to two goroutines writing to it at the same time.
	Go can help us to spot race conditions with its built in race detector.
	To enable this feature, run the tests with the race flag:
		go test -race .

	We can solve this data race by coordinating our goroutines using channels.
	Channels are a Go data structure that can both receive and send values.
	These operations, along with their details, allow communication between different processes.
	In this case we want to think about the communication between the parent process and each of the goroutines
	that it makes to do the work of running the WebsiteChecker function with the url.

	See v2 folder.
*/

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)

		time.Sleep(2 * time.Second)
	}

	return results
}
