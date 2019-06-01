package main

import (
	"fmt"
	"time"
)

func main() {

	tm := time.Now()

	switch {
	case tm.Hour() < 12:
		fmt.Println("good Morning")
	case tm.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}

}
