package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("css/main.css")
	fmt.Println(dir)  // css
	fmt.Println(file) // main.css

	var onlyFile string
	_, onlyFile = path.Split("path/path.html")
	fmt.Println(onlyFile)

	_, f := path.Split("path/path.html")
	fmt.Println(f)
}
