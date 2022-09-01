package main

import (
	"fmt"
	"strings"
)

func main() {
	testcases := []string{"abcd", "abCdefA", "aabcd"}

	for _, tc := range testcases {
		fmt.Println(unique(tc))
	}
}

func unique(str string) bool {
	lower := strings.ToLower(str)
	hash := make(map[rune]struct{})

	for _, sym := range lower {
		if _, ok := hash[sym]; ok {
			return false
		}

		hash[sym] = struct{}{}
	}

	return true
}
