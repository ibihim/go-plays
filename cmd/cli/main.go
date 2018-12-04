package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "World", "A name to say hello to.")

	var spanish bool
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")

	flag.Parse()

	if spanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
