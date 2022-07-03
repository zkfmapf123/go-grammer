package types

import "fmt"

type book struct {
	title string
	price int
}

func NewBook(title string, price int) book {
	b := book{title, price}
	return b
}

func (b book) Print() {
	fmt.Println(b.title, b.price)
}
