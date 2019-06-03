package main

import "fmt"

func main() {
	even_number := []int{2, 4, 6, 8, 10}
	boolean_val := []bool{true, false, true, false, true}
	fmt.Println(even_number)
	fmt.Println(boolean_val)
	data := []struct {
		x int
		y string
	}{
		{1, "iron-man"},
		{2, "captain-america"},
		{3, "thor"},
	}
	fmt.Println(data)

}
