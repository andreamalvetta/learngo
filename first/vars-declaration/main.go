package main

import "fmt"

// You can't use short declaration at package level
var safe = true

func main() {

	// var name = "Carl"
	// var isScientist = true
	// var age = 62
	// var degree = 5.

	name := "Carl"
	isScientist := true
	age := 62
	degree := 5.

	fmt.Println(name, isScientist, age, degree)
}
