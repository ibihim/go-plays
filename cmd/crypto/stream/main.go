package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/salsa20"
)

const (
	nonceLen = 24
	keyLen   = 32
)

func main() {
	// io
	msg := []byte("Hello World")
	cipherText := make([]byte, len(msg))
	password := "iamspiderman"

	// Create nonce for Salsa20
	nonce := make([]byte, nonceLen)
	if _, err := rand.Read(nonce); err != nil {
		fmt.Println(err)
		return
	}

	// Transform Key in right format
	n := make([]byte, nonceLen)
	if _, err := rand.Read(n); err != nil {
		fmt.Println(err)
		return
	}
	tmpKey := pbkdf2.Key([]byte(password), n, 4096, keyLen, sha256.New)
	key := new([keyLen]byte)
	copy(key[:], tmpKey)

	// Stream cipher in action
	salsa20.XORKeyStream(cipherText, msg, nonce, key)
	fmt.Println("salsa20 encrypted msg", string(cipherText))

	// decrypt
	plainText := make([]byte, len(msg))
	salsa20.XORKeyStream(plainText, cipherText, nonce, key)

	fmt.Println("salsa20 plain text", string(plainText))
}
