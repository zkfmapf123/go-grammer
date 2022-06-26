package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"zkfmapf123/json-reader/utils"
)

func main() {

	var jsonBook utils.Book
	var inputs []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {

		inputs = append(inputs, in.Bytes()...)
	}

	err := json.Unmarshal(inputs, &jsonBook)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonBook)

}
