package main

import (
	"oop/interfaces"
	"oop/types"
)

func main() {

	getPrinter()
	getDiscount()
}

// print
func getPrinter() {
	var (
		mobydick   = types.NewBook("moby dick", 10)
		callOfDuty = types.NewGame("call of duty", 10)
		hailo      = types.NewGame("hailo", 10)
	)

	var store interfaces.List
	store = append(store, &mobydick, &callOfDuty, &hailo)
	store.Print()
}

// dynamic discount
func getDiscount() {
	var (
		mobydick   = types.NewBook("moby dick", 10)
		callOfDuty = types.NewGame("call of duty", 10)
		hailo      = types.NewGame("hailo", 10)
		haligali   = types.NewPuzzle("hali gali", 50)
	)

	var store interfaces.List
	store = append(store, &mobydick, &callOfDuty, &hailo, &haligali)
	store.Discount(10)
	store.Print() // hali gali price - 10
}
