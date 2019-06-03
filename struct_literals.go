package main

import "fmt"

type Employee struct {
	name string
	age  int
}

var (
	v  = Employee{"renjith", 25}
	v1 = Employee{"renjith", 1}
	v2 = Employee{}
	p  = &v
)

func main() {
	fmt.Println(v, v1, v2, p)

}
