package main

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNoInput error = errors.New("you must provide input: <tool name> <hashedSecret> <plainTextSecret>")
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(ErrNoInput)
		return
	}

	hashedSecret := os.Args[1]
	plainTextSecret := os.Args[2]

	if hashedSecret == "" || plainTextSecret == "" {
		fmt.Println(ErrNoInput)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedSecret), []byte(plainTextSecret)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok")
}
