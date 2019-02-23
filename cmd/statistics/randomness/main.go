package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	fmt.Println("vim-go")
}

func newPin() (string, error) {
	pin, err := rand.Int(rand.Reader, big.NewInt(999999))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", pin.Int64()), nil
}
