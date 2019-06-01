package main

import (
	"fmt"
	"math"
)

func pow(x, y, lim float64) float64 {
	var d float64
	if v := math.Pow(x, y); v < lim {
		d = v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return d
}
func main() {

	fmt.Println(pow(3, 2, 10))

}
