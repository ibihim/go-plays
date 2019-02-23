package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
)

func foo() []byte {
	password := []byte("iamspiderman")
	plainText := []byte("my spider sensors are tingling")

	// does not care about pbkdf2 in this example
	h := sha256.New()
	h.Write([]byte(password))
	key := h.Sum(nil)

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 1)
	}

	nonce := make([]byte, 12)
	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 1)
	}
	return aesgcm.Seal(nil, nonce, plainText, nil)
}

func main() {
	fmt.Println(string(foo()))
}
