package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main(){

	ans := area(10,10)
	fmt.Println(ans)
}