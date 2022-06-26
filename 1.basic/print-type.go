package main

import "fmt"

func main() {

	var (
		speed = 10
		heat  = 223.13
		off   = false
		brand = "str"
	)

	// %T === typeof
	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)
}
