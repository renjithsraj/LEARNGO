package main

import "fmt"

type I interface {
	// No method here name as M
	// I think we will get an runtime error
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
