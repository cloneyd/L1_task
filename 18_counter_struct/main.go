package main

import (
	"fmt"
	"sync"
)

type concurrentCounter struct {
	mu  sync.Mutex
	val int64
}

func main() {
	var wg sync.WaitGroup
	var cc concurrentCounter

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			cc.Add()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(cc.Value())
}

func (cc *concurrentCounter) Add() {
	defer cc.mu.Unlock()

	cc.mu.Lock()
	cc.val++
}

func (cc *concurrentCounter) Value() int64 {
	return cc.val
}
