package main

import (
	"fmt"
	"oop/repo"
)

type game struct {
	title string
	price int
}

func main() {
	// basicOop()
	pointOop()
}

// basic
func basicOop() {
	var (
		mobydick  = repo.NewBook("mobydick", 3000)
		minecraft = repo.NewBook("mincraft", 3000)
		tetris    = repo.NewBook("tetris", 3000)
	)

	mobydick.GetPrint()
	minecraft.GetPrint()
	tetris.GetPrint()

	minecraft.SetTitle("new_minecraft")
	minecraft.SetPrice(5000)

	tetris.SetTitle("new_tetris")
	tetris.SetPrice(6000)

	fmt.Println("--------------------")
	mobydick.GetPrint()
	minecraft.GetPrint()
	tetris.GetPrint()
}

// oop
func pointOop() {

	var (
		callOfDuty = game{"callofDuty", 30000}
		hailo      = game{"hailo", 50000}
	)

	var gamePack []*game
	gamePack = append(gamePack, &callOfDuty, &hailo)

	for _, it := range gamePack {
		fmt.Println(*it)
	}

}
