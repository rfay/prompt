package main

import (
	"github.com/Bowery/prompt"
	"net/mail"
	"fmt"
)

func main() {
	result, err := ExampleBasicDefault()
	if (err != nil) {
		fmt.Printf("Failed ExampleBasic: %v", err)
	}
	fmt.Printf("result=%s\n", result)
}
func ExampleBasicDefault() (string, error) {
	email, err := prompt.BasicDefault("Email", "Somedefault@nowhere.com")
	if err != nil {
		return "", err
	}

	_, err = mail.ParseAddress(email)
	return email, err
}

