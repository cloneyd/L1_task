package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "sun dog snow"

	fmt.Println(reverseWords(str, " "))
}

func reverseWords(src string, delim string) string {
	var sb strings.Builder
	words := strings.Split(src, delim)

	for i, j := 0, len(words)-1; i < len(words)/2; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	sb.WriteString(words[0])

	for i := 1; i < len(words); i++ {
		sb.WriteString(" ")
		sb.WriteString(words[i])
	}

	return sb.String()
}
