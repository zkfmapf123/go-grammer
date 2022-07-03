package types

import "fmt"

type game struct {
	title string
	price int
}

func NewGame(title string, price int) game {
	g := game{title, price}
	return g
}

func (g *game) Print() {
	fmt.Println(g.title, g.price)
}

func (g *game) Discount(ratio float64) {
	g.price -= int(ratio)
}
