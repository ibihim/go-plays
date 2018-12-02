package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	nonceLen := 24
	nonce := make([]byte, nonceLen)

	if _, err := rand.Read(nonce); err != nil {
		fmt.Println(err)
		return
	}

	res := pbkdf2.Key([]byte("iamspiderman"), nonce, 4096, 32, sha256.New)

	fmt.Println("pbkdf2 key with len", len(res), res)
}
