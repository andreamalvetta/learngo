package main

import "fmt"

func main() {
	name := "Nikola"
	name, age := "Marie", 66
	fmt.Println(name, age)

	// name = "Albert"
	// birth := 1879
	name, birth := "Albert", 1879
	fmt.Println(name, birth)

	// --- TIPS ---
	// Use short declaration for redeclaration
	width, height := 100, 50

	// DON'T USE -- it's verbose
	// width = 50
	// color := "red"

	width, color := 50, "red"

	fmt.Println(width, height, color)
}
