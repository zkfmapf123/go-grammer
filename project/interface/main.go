package main

import "interface/src/repo"

type printer interface {
	Print()
}

type list []printer

func (l list) print() {

	for _, it := range l {
		it.Print()
	}
}

func (l list) discount() {

	type discounter interface {
		Discount(num int)
	}

	for _, it := range l {
		if it, ok := it.(discounter); ok {
			it.Discount(5)
		}
	}
}

func main() {

	store := []list{
		repo.NewBook("moby dick", 2000, 12312412),
		repo.NewGame("call of duty", 50000),
		repo.NewPuzzle("hali gali", 50000),
	}

}
