package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Unicode -> 1~4 byte
	name := "carl"  // 1byte
	imoji := "😀😀😀😀" // 4byte

	fmt.Println(len(name))  // 4
	fmt.Println(len(imoji)) //16

	// Solution
	utfNameLen := utf8.RuneCountInString(name)
	utfImojiLen := utf8.RuneCountInString(imoji)

	fmt.Println(utfNameLen)  // 4
	fmt.Println(utfImojiLen) // 4
}
