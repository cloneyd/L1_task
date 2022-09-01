package main

import (
	"errors"
	"fmt"
)

func main() {
	data := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if val, err := delete(data, 9); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}

func delete(src []int64, idx int) ([]int64, error) {
	if idx < 0 || idx > len(src)-1 {
		return nil, errors.New("index out of range")
	}

	return append(src[:idx], src[idx+1:]...), nil
}
