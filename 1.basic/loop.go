package main

import (
	"fmt"
	"strings"
)

func main() {
	basicSum := basicLoop(1, 100)
	nestedSum := nestedLoop(1, 100)
	rangeSum := rangeLoop(1, 100)
	stringsLoop("hi my name is leedonggyu")

	fmt.Println(basicSum, nestedSum, rangeSum)
}

func basicLoop(start int, end int) int {
	var sum int
	for i := start; i < end; i++ {
		sum += i
	}

	return sum
}

func nestedLoop(start int, end int) int {
	var sum int

	return sum
}

// use range -> required array
// range (index, value)
func rangeLoop(start int, end int) int {
	arr := make([]int, end)
	var sum int

	// make array
	for i := start; i < end; i++ {
		arr[i] = i
	}

	for _, v := range arr {
		sum += v
	}

	return sum
}

// example
func stringsLoop(s string) {

	words := strings.Fields(s)

	for _, v := range words {
		fmt.Println(v)
	}
}
