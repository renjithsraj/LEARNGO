package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["age"] = 27
	fmt.Println(m)
	v, ok := m["age"]
	fmt.Println(v, ok)
}
