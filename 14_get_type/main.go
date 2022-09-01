package main

import "fmt"

func getType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("dont know this type")
	}
}

func main() {
	getType(21)
	getType(15.1)
	getType("hello")
	getType(true)
	getType(1 + 2i)
}
