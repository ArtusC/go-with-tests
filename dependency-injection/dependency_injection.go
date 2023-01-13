package dependency_injection

import (
	"fmt"
	"io"
	"net/http"
)

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// To run and test the greeter handler working in the localhost, run in the terminal: `go run main.go`
