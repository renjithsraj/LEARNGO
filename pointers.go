package main

import "fmt"

func main() {

	i, j := 42, 1200
	fmt.Println("Initial value is", i, j)
	p := &i
	fmt.Printf("The pointers value of %v is %v\n", i, p)
	*p = 25
	fmt.Printf("Initial value is %v\n", i)
	p = &j
	fmt.Printf("Value is %v and %v\n", j, p)
	*p = *p / 37
	fmt.Println(j)

}
