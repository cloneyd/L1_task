package main

import "fmt"

func main() {
	data := []int64{33, 10, 15, 7}

	fmt.Println(quicksort(data))
}

func quicksort(array []int64) []int64 {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]

		less := []int64{}
		greater := []int64{}

		for i := 1; i < len(array); i++ {
			if array[i] > pivot {
				greater = append(greater, array[i])
			} else {
				less = append(less, array[i])
			}
		}

		res := append(quicksort(less), pivot)

		return append(res, quicksort(greater)...)
	}
}
