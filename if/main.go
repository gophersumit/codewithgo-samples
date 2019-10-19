package main

import (
	"fmt"
)

func main() {
	flag := true
	count := 1
	if flag {
		fmt.Println("Flag is present")
	}

	// invalid if statements
	// if "hello" {
	// 	fmt.Println("Invalid if")
	// }
	// if count {
	// 	fmt.Println("Invalid if")
	// }

	// multiple branches
	marks := 70
	if marks > 80 {
		fmt.Println("Great!")
	} else if marks > 60 {
		fmt.Println("Good!")
	} else if marks > 40 {
		fmt.Println("You can do better!")
	}
}
