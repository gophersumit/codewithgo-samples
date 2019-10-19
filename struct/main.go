package main

import "fmt"

func main() {

	// declaring and creating struct

	type person struct {
		name string
		age  int
	}
	var p1 person
	var p2 person
	fmt.Println(p1)
	fmt.Println(p2)

	// accessing fields
	p1.name = "Joey"
	p1.age = 30
	fmt.Println(p1)

	// zero values
	type show struct {
		name        string
		price       float64
		isAvailable bool
		rating      int
	}
	var s show
	fmt.Printf("%#v\n", s)

	// struct literal
	var p3 person = person{
		name: "Chandler",
		age:  32,
	}
	fmt.Printf("%#v\n", p3)

	// comparing structs
	var p4 person = person{
		name: "Chandler",
		age:  32,
	}
	var p5 person = person{
		name: "Joey",
		age:  30,
	}
	fmt.Println("are p4 and p5 equal? ==>", p4 == p5)

	// struct as another struct field
	type human struct {
		name string
		age  int
	}

	type superHuman struct {
		human human
		power string
	}

	superman := superHuman{
		human: human{
			name: "Clark Kent",
			age:  30,
		},
		power: "Flying",
	}

	fmt.Printf("%#v\n", superman)
	fmt.Println(superman.human.name)

	// embedding structs
	type superHero struct {
		livesSaved int
		human
	}
	batman := superHero{}
	ironman := superHero{
		livesSaved: 100000,
		human: human{
			name: "Tony",
			age:  40,
		},
	}

	batman.name = "Bruce"
	batman.age = 50
	batman.livesSaved = 100

	fmt.Printf("%#v\n", batman)
	fmt.Printf("%#v\n", ironman)
}
