package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	const numWorkers = 5

	var wg sync.WaitGroup
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background()) // Контекст с возможностью отмены
	defer cancel()

	jobs := make(chan int)
	defer close(jobs)

	wg.Add(numWorkers) // Обновляем счетчик ожидаемых горутин (+ количество воркеров)
	startWorkers(ctx, &wg, jobs, numWorkers)

	signalChan := make(chan os.Signal, 1)   // Канал системных сигналов
	signal.Notify(signalChan, os.Interrupt) // В случае получения сигнала SIGINT (отправляется при нажатии Ctrl+C), отправляет его значения в канал

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for {
		select {
		case <-signalChan:
			return
		case jobs <- data[rand.Intn(len(data))]: // Записываем рандомное число в поток с задержкой
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func startWorkers(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go outputWorker(ctx, wg, jobs)
	}
}

func outputWorker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done() // По окончанию горутины обновляем счетчик (-1)

	for {
		select {
		case <-ctx.Done(): // При вызове cancel в канал Done кладется значение сигнализирующее об окончании работы контекста
			return
		case val := <-jobs: // Если из канала можно получить число - выводим его
			fmt.Println(val)
		}
	}
}
