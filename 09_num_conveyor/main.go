package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	vals := make(chan int, len(data))
	mults := make(chan int, len(data))

	go multiply(vals, mults)
	writeChan(vals, data)

	for val := range mults {
		fmt.Println(val)
	}
}

func writeChan(to chan int, from []int) {
	defer close(to)

	for _, val := range from {
		to <- val
	}
}

func multiply(vals <-chan int, mults chan<- int) {
	defer close(mults)

	for val := range vals {
		mults <- val * 2
	}
}
