package repo

import "fmt"

type book struct {
	title string
	price int
}

func NewBook(title string, price int) book {
	b := book{title, price}
	return b
}

func (b book) GetPrint() {
	fmt.Println(b.title, b.price)
}

// point receiver
func (b *book) SetTitle(title string) {
	b.title = title
}

// point receiver
func (b *book) SetPrice(price int) {
	b.price = price
}
