package main

import "fmt"

type Test interface {
	M()
}

type Data struct {
	name string
}

func (d Data) M() {
	fmt.Println(d.name)
}

func main() {
	var i Test = Data{"Renjith"}
	i.M()
}
