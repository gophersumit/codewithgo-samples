package main

import "fmt"

type location struct {
	street  string
	city    string
	pincode int
}

func main() {
	// address of variable
	a := 10
	b := &a
	fmt.Printf("Value of a is %d \n", a)
	fmt.Printf("Address of a is %v\n", &a)
	fmt.Printf("Value at pointer location is %v\n", *b)

	// Pointer type
	var intP *int
	var floatP *float64
	fmt.Println(intP, floatP)

	myLocation := &location{
		street:  "Xyz",
		city:    "Pune",
		pincode: 444101,
	}

	fmt.Printf("I live on %s street, %s city with %d code\n",
		myLocation.street, myLocation.city, myLocation.pincode)

	// sharing data
	updateLocation(myLocation)

	fmt.Printf("I live on %s street, %s city with %d code\n",
		myLocation.street, myLocation.city, myLocation.pincode)

}

func updateLocation(loc *location) {

	loc.city = "Mumbai"

}
