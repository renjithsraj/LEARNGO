package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Data() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)

}

func (v *Vertex) Scale(f float64) {

	v.x = v.x * f
	v.y = v.y * f
}

func main() {

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Data())

}
