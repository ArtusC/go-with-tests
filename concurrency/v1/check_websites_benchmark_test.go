package concurrency

import (
	"testing"
	"time"
)

func slowStubWebSiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// To run into terminal just run: go test -bench=.
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "an url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebSiteChecker, urls)
	}
}
