package main

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("test if server is working", func(t *testing.T) {
		data := "hello world"

		svr := Server(&SpyStore{response: data})

		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}
	})

	t.Run("test to check if the request is cancelled when the work is cancelled", func(t *testing.T) {

		data := "hello world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest("GET", "/", nil)

		/*
			What we do is derive a new cancellingCtx from our request which returns us a cancel function.
			We then schedule that function to be called in 5 milliseconds by using time.AfterFunc.
			Finally we use this new context in our request by calling request.WithContext.
		*/
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("request was not cancelled")
		}

	})
}
