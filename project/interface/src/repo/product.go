package repo

import "fmt"

type Product struct {
	title string
	price int
}

func (p *Product) Print() {
	fmt.Println(p.title, p.price)
}

func (p *Product) discount(ratio float64) {
	p.price -= int(ratio)
}
