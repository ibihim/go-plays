package main

import (
	"crypto/cipher"
	"fmt"
	"hash"
	"io"
	"strings"
)

type StreamEncrypt struct {
	Source io.Reader
	Block  cipher.Block
	Stream cipher.Stream
	Mac    hash.Hash
	IV     []byte
}

func (s *StreamEncrypt) Read(p []byte) (int, error) {
	n, err := s.Source.Read(p)

	if n > 0 {
		// Overwrite the read byte slice
		s.Stream.XORKeyStream(p[:n], p[:n])
		if err := writeHash(s.Mac, p[:n]); err != nil {
			return n, err
		}

		return n, err
	}

	return 0, io.EOF
}

func main() {
	src := strings.NewReader("some io.Reader stream to be read\n")

	fmt.Println("vim-go")
}
