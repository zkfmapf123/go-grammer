package main

import "fmt"

func main() {

	citySwitch("Seoul")
	rangeSwitch(10)
	fallThroughSwitch(10) // 10 , 1, 13
}

// single condition
// multiple condition
func citySwitch(city string) string {

	switch city {
	case "Incheon", "Busan":
		return "i lived the ib"
	case "Seoul":
		return "i lived the seoul"

	default:
		return "i don't have home"
	}
}

// bool expression
func rangeSwitch(num int) {

	switch {
	case num > 10:
		//
		fallthrough // pass 하는데 -> 다음조건을 무시한다
	case num < 10:
		//
		break // break
	case num == 10:

	}
}

func fallThroughSwitch(num int) {
	switch {
	case num >= 10:
		fmt.Println("10")
		fallthrough
	case num < 1:
		fmt.Println("1")
		fallthrough
	case num == 3:
		fmt.Println("13")
		break
	}
}
