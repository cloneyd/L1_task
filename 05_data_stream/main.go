package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const numSeconds = 5

	ctx, cancel := context.WithTimeout(context.Background(), numSeconds*time.Second)
	defer cancel()

	stream := make(chan int)

	go subscriber(ctx, stream)
	publisher(ctx, stream, numSeconds)
}

func publisher(ctx context.Context, stream chan int, numSeconds int) {
	defer close(stream)

	rand.Seed(42)
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for {
		select {
		case <-ctx.Done():
			return
		case stream <- data[rand.Intn(len(data))]:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func subscriber(ctx context.Context, stream chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-stream:
			fmt.Println(val)
		}
	}
}
