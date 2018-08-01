package main

import (
	"github.com/Bowery/prompt"
		"fmt"
	)

func main() {
	result, err := ExamplePassword()
	if (err != nil) {
		fmt.Printf("Failed ExamplePassword: %v", err)
	}
	fmt.Printf("result=%s\n", result)
}
func ExamplePassword() (string, error) {
	clear, err := prompt.Password("Password")
	if err != nil {
		return "", err
	}

	return string(clear), nil
}
