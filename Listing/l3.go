package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	// nil указатель на *os.PathError = nil
	fmt.Println(err)
	// false
	fmt.Println(err == nil)
}
