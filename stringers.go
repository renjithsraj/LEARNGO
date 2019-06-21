package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)

}

func main() {
	a := Person{"renjith", 28}
	b := Person{"leena", 24}
	fmt.Println(a, b)

}
