package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	for i := 0; i < 1000; i++ {
		nonce, err := newNonce(10)
		if err != nil {
			fmt.Println("erorr:", err)
			return
		}
		fmt.Println(nonce)
	}
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

func newNonce(n int) (string, error) {
	b, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
