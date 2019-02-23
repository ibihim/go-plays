package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var collision float64
	collision = 1.0

	n, _ := strconv.Atoi(os.Args[1])
	p, _ := strconv.Atoi(os.Args[2])

	for i := 0; i < p; i++ {
		reduce := float64(n-i) / float64(n)
		collision = collision * reduce
	}

	fmt.Printf("%5f\n", collision)
}
