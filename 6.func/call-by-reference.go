package main

import "fmt"

func main() {
	local := 10

	incre(local)
	fmt.Println(local, &local) // 10

	increByRefercen(&local)
	fmt.Println(local, &local)
}

// call by value
func incre(n int) {
	n++
}

// call by refernces
func increByRefercen(n *int) {
	fmt.Printf("address : %p\n", n)
	*n++
}
