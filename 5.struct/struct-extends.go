package main

import "fmt"

/*
	[is-a relationship]
	Text
	Book extends Text
	Article extends Text
*/

type Text struct {
	title string
	words int
}

type Book struct {
	Text // Composition
	isbn string
}

type Article struct {
	Text  // Composition
	paper string
}

func main() {

	socialBook := Book{
		Text{title: "social", words: 300},
		"129381203asdfav",
	}

	fmt.Println(socialBook)
}
