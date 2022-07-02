package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	LOG = iota
	COUNT
)

func main() {
	b := getReadLogs()
	fields := byteToStrings(b)
	sum := parse(fields)
	fmt.Println(sum)
}

func getReadLogs() []byte {
	b, err := ioutil.ReadFile("domain.txt")

	if err != nil {
		panic(err)
	}

	return b
}

func readLogs() []byte {
	in := bufio.NewScanner(os.Stdin)
	enter := "Write Log Parser"

	var inputs []byte

	for fmt.Println(enter); in.Scan(); fmt.Println(enter) {
		inputs = append(inputs, in.Bytes()...)

		if line := in.Text(); len(line) == 0 {
			break
		}
	}

	return inputs
}

func byteToStrings(b []byte) []string {
	byteStr := string(b)
	fields := strings.Split(byteStr, "\n")

	return fields
}

func parse(fields []string) int {

	var sum int = 0
	for _, v := range fields {

		str := strings.Fields(v)
		fmt.Printf("log : %s count : %s\n", str[LOG], str[COUNT])

		i, err := strconv.Atoi(str[COUNT])

		if err != nil {
			panic(err)
		}

		sum += i
	}

	return sum
}
