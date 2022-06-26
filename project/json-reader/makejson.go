package main

import (
	"encoding/json"
	"fmt"
	"zkfmapf123/json-reader/utils"
)

func main() {

	book := utils.Book{
		Ko: utils.Language{"안녕하세요", "수고하세요", "힘내세요"},
		En: utils.Language{"hello", "goodbye", "cheerup man"},
		Jp: utils.Language{"ohayo", "gg", "!!!!"},
	}

	b, err := json.MarshalIndent(book, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}
