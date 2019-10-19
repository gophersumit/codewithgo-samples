package main

import "fmt"

func main() {
	// traditional for loop
	for i := 1; i < 10; i++ {
		fmt.Printf("Hello %d \n", i)
	}

	// for loop as while
	j := 1
	for j < 10 {
		fmt.Printf("Hello %d\n ", j)
		j = j + 1
	}

	// continue and break
	for k := 1; k < 20; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Printf("Hello  %d\n", k)

	}
	for k := 1; k < 20; k++ {
		if k == 13 {
			break
		}
		fmt.Printf("Hello  %d\n", k)

	}

	for {
		fmt.Println("Hello!")
		// break based on some condition to exit infinite loop
		break
	}

	// iterating using range

	data := []string{"a", "b", "c"}
	for index, value := range data {
		fmt.Println(index, value)
	}

	capitals := map[string]string{
		"India":   "Pune",
		"England": "London",
		"U.S.":    "Washington",
	}
	for key, value := range capitals {
		fmt.Println(key, value)
	}

	greetings := "नमस्कार"
	for _, char := range greetings {
		fmt.Printf("%c\n", char)
	}

}
