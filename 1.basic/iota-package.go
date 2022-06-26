package main

import (
	"fmt"
	"strconv"
)

// go doc strconv
func main() {

	// int -> int
	s := strconv.Itoa(42)
	fmt.Println(s)

	// string -> int
	i, err := strconv.Atoi("42")

	if err != nil {
		panic(err)
	}

	fmt.Println(i)

	// string -> int
	is, err := strconv.Atoi("leedonggyu")

	if err != nil {
		panic(err)
	}

	fmt.Println(is)
}
