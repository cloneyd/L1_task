package main

import "fmt"

func main() {
	first := []int{4, 6, 7}
	second := []int{3, 5, 6, 7, 8}

	if len(first) > len(second) {
		fmt.Println(intersection(second, first))
	} else {
		fmt.Println(intersection(first, second))
	}
}

func intersection(first, second []int) []int {
	result := make([]int, 0, len(first))
	hash := make(map[int]struct{})

	for _, val := range first {
		hash[val] = struct{}{}
	}

	for _, val := range second {
		if _, ok := hash[val]; ok {
			result = append(result, val)
		}
	}

	return result
}
