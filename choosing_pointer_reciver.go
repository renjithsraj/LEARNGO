package main

import (
	"fmt"
	"math"
)

type Demo struct {
	X, Y float64
}

func (v *Demo) Test(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Demo) Abs() float64 {
	return math.Sqrt(v.X*v.Y + v.Y*v.Y)
}

func main() {
	v := &Demo{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Test(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

}
