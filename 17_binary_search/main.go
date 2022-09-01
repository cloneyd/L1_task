package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	testBinSearch(data, 1)
	testBinSearch(data, 4)
	testBinSearch(data, 7)
	testBinSearch(data, 10)
}

func testBinSearch(data []int, item int) {
	if res, ok := binSearch(data, item); ok {
		fmt.Printf("Index of %d: %d\n", item, res)
	} else {
		fmt.Printf("Can't find %d in array\n", item)
	}
}

// работает только с отсортированными массивами
func binSearch(data []int, item int) (int, bool) {
	left := 0
	right := len(data) - 1

	for left <= right {
		mid := left + (right-left)/2

		if data[mid] == item {
			return mid, true
		} else if data[mid] > item {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return 0, false
}
