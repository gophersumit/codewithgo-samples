package main

import "fmt"

// normal flow with defer
func sayHello() {

	defer fmt.Println("Performing cleanup")
	fmt.Println("Hello, There!")

}

func doPanic() {
	defer fmt.Println("Performing cleanup")
	panic("Oops")
	fmt.Println("Hello, There")
}
func main() {
	doPanic()
}
