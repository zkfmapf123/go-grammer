package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	go doc bufio
*/

func main() {
	in := bufio.NewScanner(os.Stdin)
	defer os.Stdin.Close()
	fmt.Println(`
		input method 
		1. single line
		2. multip line
		3. log parser
	`)
	in.Scan()
	switch in.Text() {
	case "1":
		singleLine()
		break
	case "2":
		multiLine()
		break
	case "3":
		logParser()
		break
	default:
		panic("no")
	}
}

func singleLine() {
	in := bufio.NewScanner(os.Stdin)
	defer os.Stdin.Close()
	in.Scan()

	fmt.Println(in.Text())
}

func multiLine() {
	in := bufio.NewScanner(os.Stdin)
	defer os.Stdin.Close()
	in.Scan()

	for in.Scan() {
		fmt.Println(in.Text())
	}
}

/*
	www.naver.com 10
	www.daum.net 20
	www.amazon.com 100
	www.colavo.com 240
*/
func logParser() {
	dict := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	defer os.Stdin.Close()

	for in.Scan() {
		str := strings.Fields(in.Text())
		num, err := strconv.Atoi(str[1])

		if err != nil {
			panic(err)
		}
		dict[str[0]] = num
	}

	for k, v := range dict {
		fmt.Printf("%s , %d\n", k, v)
	}
}
