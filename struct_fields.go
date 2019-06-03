package main

import "fmt"

type Vertex struct {
	age  int
	name string
}

func main() {
	v := Vertex{28, "renjith"}
	// Update age here
	fmt.Println("struct is", v)
	v.age = 29
	fmt.Println("updated struct is", v)

}
