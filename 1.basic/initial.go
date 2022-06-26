package main

import "fmt"

func main() {

	// score := 0 // Dont!
	// var score int  // default 0

	var defaultInt int       // default 0
	var defaultString string // default ""
	var defaultBool bool     // default false

	fmt.Println(defaultInt, defaultString, defaultBool)

	// group declare
	var (
		video    string
		duration int
		current  int
	)

	// Use Short Declare ":="
	width, height := 100, 50

}
