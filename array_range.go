package main

import "fmt"

func main() {
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, v := range data {
		sum += v
	}
	fmt.Println(sum)

}
