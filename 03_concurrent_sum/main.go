package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sums := make(chan int, 1)

	// Запускаем горутину в которой производится вычисление суммы квадратов чисел
	go sqrSum(numbers, sums)

	for val := range sums {
		fmt.Println("Sum:", val)
	}
}

// Вычисляет квадраты чисел, суммирует их, и кладет полученное значения суммы в канал
func sqrSum(nums []int, sums chan<- int) {
	defer close(sums)

	sum := 0

	for _, num := range nums {
		sum += num * num
	}

	sums <- sum
}
