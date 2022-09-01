package main

import "fmt"

func main() {
	size := 1 << 10
	justString := someFunc(size)
	fmt.Println(justString)
}

func someFunc(size int) string {
	return createHugeString(size)
}

func createHugeString(int) string {
	return ""
}
