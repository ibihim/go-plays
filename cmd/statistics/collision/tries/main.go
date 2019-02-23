package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	collision := 1.0

	n, _ := strconv.Atoi(os.Args[1])
	p, _ := strconv.ParseFloat(os.Args[2], 64)
	p = 1.0 - p

	var i int
	for i = 0; collision > p; i++ {
		reduce := float64(n-i) / float64(n)
		collision = collision * reduce
	}

	fmt.Printf("A collision will happen in %.1f%% of cases after %d within a space of %d numbers\n", (1.0-p)*100.0, i, n)
}
