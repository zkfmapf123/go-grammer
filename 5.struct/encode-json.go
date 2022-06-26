package main

import (
	"encoding/json"
	"fmt"
)

// json 모델은 Pascal 케이스로 작성이 되야함
// json key 네이밍 지정가능
type user struct {
	Name     string `json:"username"`
	Password string `json:"-"` // 사용안함
	Grade    string `json:"grade,omitempty"`
}

// go run main.go > users.json
func main() {

	createUser()
}

func createUser() {
	user := []user{
		{"aaa", "1234", "best"},
		{"bbb", "1234", "good"},
		{"ccc", "1234", "better"},
		{"ddd", "1234", "worst"},
	}

	// []byte, Error
	out, err := json.MarshalIndent(user, ">>", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
