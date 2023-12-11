//go:build unit
// +build unit

package context_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	cp "github.com/ArtusC/go-with-tests/context/v3"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		/*
			We are simulating a slow process where we build the result slowly by appending the string, character by character in a goroutine.
			When the goroutine finishes its work it writes the string to the data channel.
			The goroutine listens for the ctx.Done and will stop the work if a signal is sent in that channel.
			Finally the code uses another select to wait for that goroutine to finish its work or for the cancellation to occur
		*/
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("cancelled fetch request")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("test if server is work returning data", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		svr := cp.Server(store)

		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() == data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}
	})

	t.Run("test to check if the request is cancelled when the work is cancelled", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		svr := cp.Server(store)

		request := httptest.NewRequest("GET", "/", nil)

		/*
			What we do is derive a new cancellingCtx from our request which returns us a cancel function.
			We then schedule that function to be called in 5 milliseconds by using time.AfterFunc.
			Finally we use this new context in our request by calling request.WithContext.
		*/
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		fmt.Println(response)

		if response.written {
			t.Error("a response should not been written")
		}
	})
}

func TestServer_2(t *testing.T) {
	data := "hello world"
	store := &SpyStore{response: data, t: t}
	svr := cp.Server(store)

	request := httptest.NewRequest("GET", "/", nil)

	/*
		What we do is derive a new cancellingCtx from our request which returns us a cancel function.
		We then schedule that function to be called in 5 milliseconds by using time.AfterFunc.
		Finally we use this new context in our request by calling request.WithContext.
	*/
	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(5*time.Millisecond, cancel)
	request = request.WithContext(cancellingCtx)

	response := &SpyResponseWriter{}

	svr.ServeHTTP(response, request)

	fmt.Println(response)

	if response.written {
		t.Error("a response should not been written")
	}
}
