package main

import (
	"fmt"
)

func main() {

	// declaring arrays
	var numbers [8]float64
	var days [7]string
	var marks [10]int

	// accessing array elements
	days[0] = "Sunday"   // days[0] is the first element
	days[6] = "Saturday" // days[6] is the last element
	// days[7] = "Funday" // this will not work since we do not have 8 elements in array
	fmt.Println(days)
	// output [Sunday      Saturday]

	// zero values
	fmt.Println(numbers, marks) // both numbers and marks elements get their zero values
	//output : [0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0]

	// array literals
	var data [2]string = [2]string{
		"Hello",
		"World!",
	}
	fmt.Println(data)
	// output [Hello World!]

	// ellipsis
	newData := [2]string{"Hello", "World"}
	fmt.Println(newData)
	// output [Hello World!]

	// output [Hello World!]
	// short variable declaration for array literals
	arr := [...]string{"Hello", "World", "From Go"}
	fmt.Println(arr)
	// output [Hello World From Go]

	// iterating over array
	days = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index := 0; index < 7; index++ {
		fmt.Println(days[index])
	}
	// use of len()
	for index := 0; index < len(days); index++ {
		fmt.Println(days[index])
	}
	// use of for...range
	for i, day := range days {
		fmt.Println(i, day)
	}
	// use of blank idetifier
	for _, day := range days {
		fmt.Println(day)
	}

	//type of array
	greetings := [...]string{"Hello", "World"}
	//greetings = [...]string{"Hello", "World", "From Go"}
	//above line is error :  cannot use [3]string literal (type [3]string) as type [2]string in assignment
	fmt.Println(greetings)

	a := [2]string{"Hello", "World"}
	b := [...]string{"Hello", "World"}
	fmt.Println(a == b)
	// output : true
}
