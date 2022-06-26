package main

import "fmt"

type FriendType = map[string]string

var friendDict FriendType

func main() {
	friendDict = make(FriendType)
	names := []string{"leedonggyu", "limjeahyodck", "sinjunghyeon", "kimhyeoncheol"}
	jobs := []string{"backend", "frontend", "soccer", "police"}

	for i := 0; i < len(names); i++ {
		friendDict[names[i]] = jobs[i]
	}

	// map use range -> k, v
	for k, v := range friendDict {
		fmt.Printf("name : %s, job : %s\n", k, v)
	}

	// getKey
	// _, ok := friendDict["asdf"]
	// if !ok {
	// 	panic("asdf is not key")
	// }

	// map use delete
	delete(friendDict, "leedonggyu")

	fmt.Println(friendDict)

}
