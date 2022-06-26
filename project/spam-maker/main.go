package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	fileDesc, err := ioutil.ReadFile("urls.txt")

	if err != nil {
		panic(err)
	}

	strArr := make([]string, len(fileDesc))
	strArr = strings.Fields(string(fileDesc))

	for _, url := range strArr {
		s := strings.Split(url, ".")
		s[1] = strings.Repeat("*", len(s[1]))
		fmt.Println(strings.Join(s, "."))
	}

}
