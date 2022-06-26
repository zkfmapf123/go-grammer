package main

import (
	"fmt"
	"strings"
)

// go doc strings
func main() {
	msg := "helloWorld"

	fmt.Println(strings.Repeat(msg, 3))
	fmt.Println(strings.ToUpper(msg))
	fmt.Println(strings.ToLower(msg))

}
