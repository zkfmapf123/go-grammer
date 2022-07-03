package types

import "fmt"

type puzzle struct {
	title string
	price int
}

func NewPuzzle(title string, price int) puzzle {
	p := puzzle{title, price}
	return p
}

func (p puzzle) Print() {
	fmt.Println(p.title, p.price)
}

func (p *puzzle) Discount(num float64) {
	fmt.Println(p.price, "는", num, " 가격만큼 깎입니다.")
	p.price -= int(num)
}
