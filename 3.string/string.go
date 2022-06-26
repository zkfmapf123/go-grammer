package main

import "fmt"

// byte, Rune -> string
func main() {
	useByte()
}

func useByte() {
	start, stop := 'A', 'Z'

	for n := start; n <= stop; n++ {
		fmt.Printf("string : %s -> %d -> %x\n", string(n), n, n)
	}

	// useRuneCountInString
}
