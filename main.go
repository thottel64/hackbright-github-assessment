package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	sum := add(2, 3)
	fmt.Println("the sum is: ", sum)
	diff := subtract(5, 2)
	fmt.Println("the difference is: ", diff)
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}
