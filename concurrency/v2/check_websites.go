package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	/*
		Now when we iterate over the urls, instead of writing to the map directly
		we're sending a result struct for each call to wc to the resultChannel
		with a send statement.
		This uses the <- operator, taking a channel on the left and a value on the right.
	*/
	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	/*
		The next for loop iterates once for each of the urls.
		Inside we're using a receive expression, which assigns a value received from a channel to a variable.
		This also uses the <- operator, but with the two operands now reversed:
			the channel is now on the right and
			the variable that we're assigning to is on the left.
	*/
	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
