package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("the new function is", add(3, 5))
}
