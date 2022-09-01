package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	squares := make(chan int, len(numbers))

	// Запускаем горутину в которой производится вычисление квадратов чисел
	go sqr(numbers, squares)

	fmt.Println("Squares:")
	for val := range squares {
		fmt.Println(val)
	}
}

// Вычисляет квадрат числа и кладет полученные значения в канал
func sqr(nums []int, sqrs chan<- int) {
	// По завершению выполнения функции закрываем канал
	defer close(sqrs)

	// Вычисленные значения кладём в канал
	for _, num := range nums {
		sqrs <- num * num
	}
}
