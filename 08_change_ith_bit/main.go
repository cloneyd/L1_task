package main

import "fmt"

func main() {
	num := int64(8)
	pos := int64(1)
	val := int64(1)
	fmt.Printf("%b\n", num)

	modifyBit(&num, pos, val)

	fmt.Printf("%b\n", num)
}

func modifyBit(num *int64, pos int64, val int64) {
	if !(val == 0 || val == 1) {
		return
	}

	mask := int64(1 << pos)
	*num = (*num & ^mask) | (val << pos)
}
