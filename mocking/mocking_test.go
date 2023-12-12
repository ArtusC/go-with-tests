//go:build unit
// +build unit

package mocking_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	mock "github.com/ArtusC/go-with-tests/mocking"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountDown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleepPrinter := &SpyCountdownOperations{}
		mock.Countdown(buffer, spySleepPrinter)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		mock.Countdown(spySleepPrinter, spySleepPrinter)
		fmt.Println("spySleepPrinter = ", spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleep(t *testing.T) {

	t.Run("test SpyTime", func(t *testing.T) {
		sleepTime := (5 * time.Second)

		spyTime := &SpyTime{}
		sleeper := mock.ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
