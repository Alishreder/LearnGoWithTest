package mocking

import (
	"fmt"
	"io"
	"time"
)

type (
	Sleeper interface {
		Sleep()
	}

	SpySleeper struct {
		Calls int
	}

	SpyCountdownOperations struct {
		Calls []string
	}

	ConfigurableSleeper struct {
		Duration  time.Duration
		SleepFunc func(time.Duration)
	}

	SpyTime struct {
		durationSlept time.Duration
	}
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	countSleeper   = 4
	write          = "write"
	sleep          = "sleep"
)

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

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(out, finalWord)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}
