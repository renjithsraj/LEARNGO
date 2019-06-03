package main

import "fmt"

func main() {
	data := [5]int{2, 4, 6, 8, 10}

	var data1 []int = data[0:4]
	fmt.Println(data1)
}
