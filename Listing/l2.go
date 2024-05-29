package main

import "fmt"

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	// 2 тк дефер вызывается в конце
	fmt.Println(test())
	// 1
	fmt.Println(anotherTest())
}
