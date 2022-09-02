package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration) {
	end := time.Now().Add(dur)

	for time.Now().Before(end) {
		continue
	}
}

func main() {
	start := time.Now()
	dur := time.Duration(500 * time.Millisecond)

	sleep(dur)

	fmt.Println("Elapsed", time.Since(start))
}
