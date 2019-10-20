package main

import (
	"fmt"
)

func main() {

	// simple switch statement
	i := 2
	switch i {
	case 1:
		fmt.Println("i is one")
	case 2:
		fmt.Println("i is two")
	case 3:
		fmt.Println("i is three")
	case 4:
		fmt.Println("i is four")
	}

	dayOfWeek := 6
	switch {
	case dayOfWeek == 6 || dayOfWeek == 7:
		fmt.Println("It's weekend!")
	case dayOfWeek < 6 && dayOfWeek > 0:
		fmt.Println("It's weekday!")
	default:
		fmt.Println("You are not on Earth!")
	}
}
