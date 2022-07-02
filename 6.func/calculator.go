package main

import "fmt"

func main() {
	idx1, idx2 := 10, 20

	result1 := add(idx1, idx2)
	result2 := min(idx1, idx2)
	result3 := mul(idx1, idx2)
	result4 := div(idx1, idx2)

	fmt.Println(result1, result2, result3, result4)
}

func add(idx1 int, idx2 int) int {
	return idx1 + idx2
}

func min(idx1 int, idx2 int) int {
	return idx1 - idx2
}

func mul(idx1 int, idx2 int) int {
	return idx1 * idx2
}

func div(idx1 int, idx2 int) int {
	return idx1 / idx2
}
