package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{ 1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
