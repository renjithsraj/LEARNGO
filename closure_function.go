package main

import "fmt"

func Add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := Add(), Add()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

}
