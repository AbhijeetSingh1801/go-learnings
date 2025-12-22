package dependencyinjection

import (
	"fmt"
	"io"
	"time"
)

type SpySleeper struct {
	Calls int
}
type Sleeper interface {
	Sleep()
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(out, finalWord)
}