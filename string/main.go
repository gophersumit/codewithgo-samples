package main

import (
	"fmt"
)

func main() {
	// declaring strings
	var str1 string = "Hello, World!"
	var str2 = "Hello, World!"
	str3 := "Hello, World!"

	fmt.Println(str1, str2, str3)
	// output: Hello, World! Hello, World! Hello, World!

	// use len() to get length of string
	fmt.Printf("length of str1 is %d \n", len(str1))

	//substring operation
	str4 := str1[7:13]
	fmt.Println(str4)

}
