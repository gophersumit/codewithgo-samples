package main

import "fmt"

type user struct {
	name string
}

func sayHello(u user) {
	fmt.Printf("Hello, %s from function \n", u.name)
}

func (u user) sayHello() {
	fmt.Printf("Hello, %s from method\n", u.name)
}

func (u *user) updateUser() {
	u.name = "Bob"
}

func main() {
	u := &user{
		name: "Gopher",
	}
	// function call
	//	sayHello(u)
	// method call
	u.sayHello()
	u.updateUser()
	fmt.Println(u.name)
}
