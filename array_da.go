// Arrays
// The type [n]T is an array of n values of type T.
// The expression
// var a [10]int
// declares a variable a as an array of ten integers.
// An array's length is part of its type, so arrays cannot be resized.
// This seems limiting, but don't worry;
// Go provides a convenient way of working with arrays.

package main

import "fmt"

func main() {
	var data [2]string
	data[0] = "renjith"
	data[1] = "leena"
	fmt.Println(data)
	Evennum := [5]int{2, 4, 6, 8, 10}
	fmt.Println(Evennum)

}
