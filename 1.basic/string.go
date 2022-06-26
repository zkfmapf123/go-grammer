package main

import "fmt"

func main() {

	// string literal ""
	stringLiteral := "hi there"

	// raw string literal
	// multiline ``
	rawStringLiteral := `
		what is your name?
		hi my name is leedonggyu
	`

	fmt.Println(stringLiteral)
	fmt.Println(rawStringLiteral)

}
