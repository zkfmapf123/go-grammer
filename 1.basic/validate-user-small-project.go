package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	email, password := "zkfmapf123", "dl1ehd2rb3!"

	eErr := validEmail(email)

	if eErr != nil {
		panic(eErr)
	}

	pErr := valiePassword(password)

	if pErr != nil {
		panic(pErr)
	}

	fmt.Println("Success")

}

func validEmail(email string) error {
	if len(email) < 3 {
		return errors.New("email length < 3")
	}
	return nil
}

func valiePassword(password string) error {

	if !strings.Contains(password, "!") {
		return errors.New("password containe !")
	}

	return nil
}
