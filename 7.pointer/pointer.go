package main

import "fmt"

// map, type -> pointer
// etc value -> not pointer

func main() {
	// printPointer(100)
	pointerBasic()
}

func printPointer(b int) {
	// b := paramter

	p := &b
	v := *p

	fmt.Println(p) // address
	fmt.Println(v) // pointer
}

func pointerBasic() {
	nums := [...]int{1, 2, 3}
	incre(nums)
	fmt.Println(nums)

	increByPtr(&nums)
	fmt.Println(nums)

}

func incre(num [3]int) {

	for i := range num {
		num[i]++
	}
}

func increByPtr(num *[3]int) {

	for i := range num {
		num[i]++
	}
}
