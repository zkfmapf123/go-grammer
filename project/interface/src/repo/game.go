package repo

type game struct {
	Product
}

func NewGame(title string, price int) game {
	g := game{
		Product{title, price},
	}
	return g
}
