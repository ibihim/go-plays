package main

import "fmt"

type NotFoundError struct {
	ID string
}

func (_ NotFoundError) HelloWorld() {
	fmt.Println("hello world")
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("thing with ID %q was not found", err.ID)
}

type Thing struct{}

func FindThing(id string) (*Thing, error) {
	return nil, NotFoundError{id}
}

func main() {
	_, err := FindThing("123")
	if err != nil {
		fmt.Println("err != nil")

		if t, ok := err.(NotFoundError); !ok {
			fmt.Println("!ok")
		} else {
			fmt.Println("ok")
			t.HelloWorld()
		}
	} else {
		fmt.Println("err == nil")
	}
}
