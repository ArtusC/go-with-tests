package mocking

import (
	"fmt"
	"io"
	"time"
)

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration        time.Duration
	SleeperDuration func(time.Duration)
}

const finalWord = "Go!"
const countDownStart = 3

func (c *ConfigurableSleeper) Sleep() {
	c.SleeperDuration(c.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
