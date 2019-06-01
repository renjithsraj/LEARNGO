package main

import (
	"fmt"
	"runtime"
)

func main() {

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other than linux")

	}

}
