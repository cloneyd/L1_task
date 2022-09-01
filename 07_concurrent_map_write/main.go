package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex // Создаем mutex, чтобы только одна горутина имела доступ к переменной в промежуток времени

	data := make(map[int]struct{})

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go set(data, i, &wg, &mutex)
	}
	wg.Wait()

	fmt.Println(data)
}

func set(data map[int]struct{}, key int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer mutex.Unlock() // По завершению работы функции разблокируем mutex
	defer wg.Done()

	mutex.Lock() // Блокируем mutex чтобы другие горутины не имели доступ к data
	if _, ok := data[key]; !ok {
		data[key] = struct{}{}
	}
}
