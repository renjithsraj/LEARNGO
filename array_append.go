package main

import "fmt"

func main() {
	var data []int

	fmt.Println("length", data)
	data = append(data, 1, 3, 4)
	fmt.Printf("data %v length %d\n", data, len(data))
	data = append(data, 1, 2, 3)
	fmt.Println(data)

}
