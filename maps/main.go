package main

import (
	"fmt"
)

func main() {
	// declaring maps
	var capitals map[string]string
	fmt.Println(capitals)
	fmt.Println("is map nil ==>", capitals == nil) // returns true
	// since capitals is nil map, any assignments will cause run-time panic
	// capitals["India"] = "Delhi"

	// creating maps
	caps := make(map[string]string)
	// we can also use var keyword as below
	// var caps map[string]string = make(map[string]string)
	fmt.Println(caps)
	fmt.Println("is map nil ==>", caps == nil) // returns false
	caps["India"] = "Delhi"
	caps["England"] = "London"
	caps["U.S."] = "Washington"
	fmt.Println(caps)

	// map literals
	capitalsOfCountries := map[string]string{
		"India":   "Delhi",
		"England": "London",
		"U.S.":    "Washington",
	}
	fmt.Println(capitalsOfCountries)

	// accessing map elements
	fmt.Println(capitalsOfCountries["India"])
	fmt.Println(capitalsOfCountries["England"])

	// non exisiting key
	fmt.Println(capitalsOfCountries["XYZ"])
	// prints empty string

	// zero value for map
	var s map[string]string
	var i map[string]int
	fmt.Println("is s nil =>", s == nil)
	fmt.Println("is i nil =>", i == nil)

	// zero value for map elements
	j := map[string]int{}
	fmt.Println(j["India"])

	// empty map
	var t = make(map[string]string)
	k := map[string]int{}
	fmt.Println("is t nil =>", t == nil)
	fmt.Println("is k nil =>", k == nil)

	// len of nil and empty
	//	var s map[string]string
	//	var t = make(map[string]string)
	fmt.Println(len(s))
	fmt.Println(len(t))

	// iterating over map
	for key, value := range caps {
		fmt.Printf("key =>%v, value => %v\n", key, value)
	}

	// blank identifier
	for _, value := range caps {
		fmt.Printf("value => %v\n", value)
	}
	for key, _ := range caps {
		fmt.Printf("key =>%v\n", key)
	}

}
