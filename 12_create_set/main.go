package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	res := createSet(data)
	fmt.Println(res)
}

func createSet(vals []string) []string {
	set := []string{}
	hash := make(map[string]struct{})

	for _, key := range vals {
		if _, ok := hash[key]; !ok {
			hash[key] = struct{}{}
		}
	}

	for key := range hash {
		set = append(set, key)
	}

	return set
}
