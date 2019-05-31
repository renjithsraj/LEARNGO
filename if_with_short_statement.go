package main

import (
	"fmt"
	"math"
)

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); x < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(2, 2, 10))

}
