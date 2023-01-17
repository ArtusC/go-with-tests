package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

const finalWord = "Go!"
const countDownStart = 3

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
