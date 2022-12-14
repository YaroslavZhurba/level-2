package main

import (
	"fmt"
	"os"
	"time"
	"dev01/clock"
)

var (
	host   = "0.beevik-ntp.pool.ntp.org"
	// host   = "bad"
	format = time.UnixDate
)

func main() {
	t, err := clock.GetDate(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid response:", err)
		os.Exit(1)
	}

	fmt.Println(t.Format(format))
}