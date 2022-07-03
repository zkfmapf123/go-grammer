package main

import "fmt"

func main() {

	var any interface{}

	any = []int{1, 2, 3}
	fmt.Println(any)

	any = map[int]bool{1: true, 2: false}
	fmt.Print(any)

	any = "hello"
	fmt.Println(any)

	any = 3
	fmt.Println(any)

	any = any.(int) * 2
	fmt.Println(any)
}
