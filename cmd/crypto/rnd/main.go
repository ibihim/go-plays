package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	nonceLen := 24
	nonce := make([]byte, nonceLen)

	if _, err := rand.Read(nonce); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("nonce with length", len(nonce), nonce)
}
