package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("When is saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 0:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("To far away")
	}

}
