package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/salsa20"
)

const (
	nonceLen = 24
	keyLen   = 32
)

func main() {
	salsa20Example()
	chacha20Example()
}

func salsa20Example() {
	// io
	msg := []byte("Alice 1000 EUR Bob")
	cipherText := make([]byte, len(msg))
	password := "iamspiderman"

	// Create nonce for Salsa20
	nonce := make([]byte, nonceLen)
	if _, err := rand.Read(nonce); err != nil {
		fmt.Println(err)
		return
	}

	// Transform Key in right format
	key := new([keyLen]byte)
	tmpKey, _, err := hashPassword(password)
	if err != nil {
		fmt.Println(err)
		return
	}
	copy(key[:], tmpKey)

	// Encryption
	salsa20.XORKeyStream(cipherText, msg, nonce, key)
	fmt.Println("salsa20 encrypted message string:", string(cipherText))
	fmt.Println("salsa20 encrypted message bytes: ", cipherText)

	// Decryption. Encrypting an Decrypting is the same.
	plainText := make([]byte, len(msg))
	salsa20.XORKeyStream(plainText, cipherText, nonce, key)

	fmt.Println("salsa20 plain text:", string(plainText))

	// Salsa20 is not protected against message malfroming!

	// Malform message.
	// Attempt to translate into a particular user (Krz). It will not work as key stream
	// with different IVs makes the end result unpredictable.
	cipherText[len(cipherText)-3] = cipherText[len(cipherText)-3] - 9
	cipherText[len(cipherText)-2] = cipherText[len(cipherText)-2] - 3
	cipherText[len(cipherText)-1] = cipherText[len(cipherText)-1] - 24

	modifiedPlainText := make([]byte, len(msg))
	salsa20.XORKeyStream(modifiedPlainText, cipherText, nonce, key)

	fmt.Println("salsa20 plain text:", string(modifiedPlainText))
}

func chacha20Example() {
	// io
	msg := []byte("Alice 1000 EUR Bob")
	password := "iamspiderman"

	// Transform Key in right format.
	key, _, err := hashPassword(password)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Creates an authenticated encryption with associated data (AEAD).
	// Stream cipher with message authentication code (MAC).
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		fmt.Println("Failed to instantiate XChaCha20-Poly1305:", err)
		return
	}

	// Encryption.
	nonce := make([]byte, nonceLen)
	if _, err := rand.Read(nonce); err != nil {
		fmt.Println(err)
		return
	}
	cipherText := aead.Seal(nil, nonce, []byte(msg), nil)

	fmt.Println("chacha20poly1305 cipher text (string):", string(cipherText))
	fmt.Println("chacha20poly1305 cipher text (bytes): ", cipherText)

	// Decryption.
	plaintext, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(plaintext)

	// Malforming text.
	cipherText[len(cipherText)-3] = cipherText[len(cipherText)-3] - 9
	cipherText[len(cipherText)-2] = cipherText[len(cipherText)-2] - 3
	cipherText[len(cipherText)-1] = cipherText[len(cipherText)-1] - 24

	// Attempt to decrypt malformed text.
	_, err = aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		fmt.Println("Couldn't decrypt malformed cipher text!")
		return
	}
	fmt.Println(fmt.Errorf("could decrypt malformed plaintext"))
}

// hashPassword returns a key of 32 bytes so it satisfies Salsa20 input requirements.
// The second byte slice is the nonce used.
func hashPassword(password string) ([]byte, []byte, error) {
	n := make([]byte, 24)

	if _, err := rand.Read(n); err != nil {
		return nil, nil, err
	}

	key := pbkdf2.Key([]byte(password), n, 4096, keyLen, sha256.New)

	return key, n, nil
}
