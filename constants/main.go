package main

import (
	"fmt"
)

// const block
const (
	pi         = 3.14
	daysInWeek = 7
)

//iota
const (
	Monday int = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
