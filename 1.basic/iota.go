package main

import "fmt"

const (
	mon = iota // 0
	tue        // 1
	wed = 3    // 2
	thu        // 3
	fri        // 3
	sat = 5    // 5
	sun        // 5
)

const (
	a = iota + 1
	b
	c
	d
)

const (
	EST = -(5 + iota)
	_
	MST
	PST
)

func main() {
	fmt.Println(mon, tue, wed, thu, fri, sat, sun)

	fmt.Println(a, b, c, d)

	fmt.Println(EST, MST, PST)
}
