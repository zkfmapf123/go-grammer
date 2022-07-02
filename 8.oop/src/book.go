package src

import "fmt"

type book struct {
	title string
	price int
}

func NewBook(title string, price int) book {
	b := book{title, price}
	return b
}

func (b book) setTitle(t string) {
	b.title = t
}

func (b book) setPrice(p int) {
	b.price = p
}

func (b book) getPrint() {
	fmt.Println(b.title, b.price)
}
