package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	NUM = 4
)

type Item struct {
	id    int
	name  string
	price int
}

type Genre struct {
	item  Item
	genre string
}

func main() {
	genres := readTxt("game.txt")

	fmt.Println(genres)
}

func readTxt(fname string) []Genre {

	itemPtr := 0
	item := make([]Genre, NUM)
	fs, err := ioutil.ReadFile(fname)

	if err != nil {
		panic(err)
	}

	field := strings.Fields(string(fs))

	arr := make([]string, NUM)
	for _, v := range field {

		arr = append(arr, string(v))

		if checkToDataLength(arr) == NUM {

			fmt.Println(arr)
			item[itemPtr] = Genre{
				item:  Item{id: stringToInt(arr[0]), name: arr[1], price: stringToInt(arr[2])},
				genre: arr[3],
			}
			itemPtr++
		}
	}

	return item
}

func checkToDataLength(arr []string) int {

	existDataCheck := 1

	for _, v := range arr {
		if v != "" {
			existDataCheck = existDataCheck + 1
		}
	}

	return existDataCheck

}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return num
}
