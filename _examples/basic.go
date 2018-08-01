package main

import (
	"github.com/Bowery/prompt"
	"net/mail"
	"fmt"
)

func main() {
	result, err := ExampleBasic()
	if (err != nil) {
		fmt.Printf("Failed ExampleBasic: %v", err)
	}
	fmt.Printf("result=%s\n", result)
}
func ExampleBasic() (string, error) {
	email, err := prompt.Basic("Email", true)
	if err != nil {
		return "", err
	}

	_, err = mail.ParseAddress(email)
	return email, err
}

