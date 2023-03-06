package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	sum := add(2, 3)
	fmt.Println(sum)
}

func add(x int, y int) int {
	return x + y
}
