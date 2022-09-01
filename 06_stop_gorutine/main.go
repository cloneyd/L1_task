package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stopWithChannel()  // Создаем канал в который кладем некое значение, сигнализируещее о необходимости горутине закончить работу
	stopWithTimeout()  // Завершаем горутину по истечению времени
	stopWithDeadline() // Завершаем горутину в конкретное время
	stopWithCancel()   // Завершаем горутину вызовом функции cancel
}

func channelWorker(done chan bool) {
	defer fmt.Println("Channel worker stopped")

	for {
		select {
		case <-done:
			return
		}
	}
}

func stopWithChannel() {
	done := make(chan bool, 1)

	go channelWorker(done)

	done <- true
}

func contextWorker(ctx context.Context) {
	defer fmt.Println("Context worker stopped")

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func stopWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go contextWorker(ctx)
}

func stopWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go contextWorker(ctx)
}

func stopWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	go contextWorker(ctx)
}
