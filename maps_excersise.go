package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("Hello world")
	count := make(map[string]int, len(words))
	for _, word := range words {
		count[word]++
	}
	fmt.Println(count)

}
