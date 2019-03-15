package main

import "fmt"

// --- TIPS ---
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

	// --- TIPS ---
	// If you don't know the value of variable don't use the short declaration
	// score := 0 // DON'T!
	// Use normal declararion instead
	var score int // already score = 0

	fmt.Println(score)

	// --- TIPS ---
	// Use normal declaration for grouping variables together to inhance readability
	var (
		// related
		video string

		// closely related
		duration int
		current  int
	)
	fmt.Println(video, duration, current)
}
