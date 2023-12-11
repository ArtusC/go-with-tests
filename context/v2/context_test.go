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
	t         *testing.T
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("request was not cancelled")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("it should not have cancelled")
	}
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("test if server is work returning data", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("test to check if the request is cancelled when the work is cancelled", func(t *testing.T) {

		data := "hello world"
		store := &SpyStore{response: data, t: t}
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

		store.assertWasCancelled()
	})
}
