package main

import "fmt"

func main() {
	a := 5.5
	b := 10.1

	switchArithm(&a, &b)

	fmt.Println(a, b)
}

func switchArithm(first, second *float64) {
	*first = *first + *second
	*second = *first - *second
	*first = *first - *second
}
