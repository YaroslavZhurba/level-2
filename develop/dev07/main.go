package main

import (
	"fmt"
	"time"
)

func receiveChan(ch chan interface{}, chans ...<-chan interface{}) {
	i := 0

Exit1:
	for {
		select {
		case <-chans[i]:
			{
				close(ch)
				break Exit1
			}
		default:
			{
				i = (i + 1) % len(chans)
				continue Exit1
			}
		}
	}
	fmt.Println("Finished")
}

func or(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go receiveChan(out, chans...)
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(4*time.Second),
		sig(8*time.Second),
		sig(6*time.Second),
		sig(3*time.Second),
		sig(5*time.Second),
	)

	fmt.Printf("Time: %v", time.Since(start))
}
