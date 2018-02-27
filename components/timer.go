package components

import (
	"time"
)

func (f *FlowsBuilder) Timer(d time.Duration) <-chan int {
	ch := make(chan int, 0)

	go func() {
		i := 0
		for {
			time.Sleep(d)
			ch <- i
		}
	}()

	return ch
}
