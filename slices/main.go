package main

import (
	"fmt"
)

func printSlice(s []string) {
	fmt.Printf("size: %v, capacity:%v, value:%v\n",
		len(s), cap(s), s)
}
func main() {
	// declaring slice
	var names []string
	printSlice(names)
	//output size: 0, capacity:0,value:[]

	// creating slice with length
	names = make([]string, 5)
	printSlice(names)

	names = make([]string, 5, 20)
	printSlice(names)
	// output : size: 5, capacity:20, value:[    ]

	// accessing slice elements
	greetings := make([]string, 2)
	greetings[0] = "Hello"
	greetings[1] = "World"
	printSlice(greetings)
	// size: 2, capacity:2, value:[Hello World]

	// zero values
	numbers := make([]float64, 8)
	months := make([]string, 12)
	fmt.Println("numbers =>", numbers)
	fmt.Println("months =>", months)

	// slice literals
	data := []string{
		"Hello",
		"World!",
	}
	printSlice(data)
	// size: 2, capacity:2, value:[Hello World!]

	// slice from array
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	weekdays := days[0:5]
	// if first element of slice is same as first element of array, start index can be skipped
	weekdays = days[:5]
	printSlice(weekdays)

	weekend := days[5:7]
	// if last element of slice is same as last element of array, end index can be skipped
	weekend = days[5:]
	printSlice(weekend)

	alldays := days[0:7]
	// if slice has all array elements, both start index and end index can be skipped
	alldays = days[:]
	printSlice(alldays)

	// empty and nil slices
	sl := make([]string, 3)[3:]
	var e []string
	printSlice(sl)
	printSlice(e)

	fmt.Println("Is s nil => ", sl == nil)
	fmt.Println("Is e nil => ", e == nil)

	// iterating over slices
	for _, day := range alldays {
		fmt.Println(day)
	}

	//append operation
	s := make([]string, 0)
	// appending elements
	s = append(s, "a")
	printSlice(s)
	s = append(s, "e")
	printSlice(s)
	s = append(s, "i")
	printSlice(s)
	s = append(s, "o")
	s = append(s, "u")
	printSlice(s)
}

// comparing slices
func testEquality(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
