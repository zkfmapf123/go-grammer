package repo

type puzzle struct {
	product Product
}

func NewPuzzle(title string, price int) puzzle {
	p := puzzle{
		Product{title, price},
	}

	return p
}

func (p *puzzle) discount(ratio int) {
	p.product.price -= ratio
}
