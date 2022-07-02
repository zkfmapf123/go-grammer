package main

import "github.com/zkfmapf123/src/book"

func main() {
	mobidick := book.NewBook("mobidict", 1000)
	mobidick.getPrint()

	mobidick.setPrice(10000)
	mobidick.setTitle("call of duty")
	mobidick.getPrint()
}
