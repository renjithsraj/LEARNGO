// The basic for loop has three components separated by semicolons:
// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 500000; i++ {
		sum += i
	}

	fmt.Println(sum)
}
