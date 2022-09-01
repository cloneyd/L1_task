package main

import "fmt"

func main() {
	a := 5
	b := 10

	switchXOR(&a, &b)

	fmt.Println(a, b)
}

// Работает только с целыми числами!
func switchXOR(first, second *int) {
	*first ^= *second
	*second ^= *first
	*first ^= *second
}
