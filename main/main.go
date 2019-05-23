package main

import (
	"LeetCode-Go/Array"
	"fmt"
)

func main() {

	array := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := Array.FindDuplicates(array)
	fmt.Println(result)
}
