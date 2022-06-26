package main

import "fmt"

type Person struct {
	job         string
	age         int
	personality string
}

var friendMap map[string]Person

func main() {

	friendMap = make(map[string]Person)

	friendMap["leedonggyu"] = Person{
		job:         "backend",
		age:         29,
		personality: "good",
	}

	// 생략가능
	friendMap["limjeahyock"] = Person{
		"frontend",
		28,
		"bad",
	}

	for k, v := range friendMap {
		fmt.Printf("name : %s\njob : %s\nage : %d\nper : %s\n", k, v.job, v.age, v.personality)
	}
}
