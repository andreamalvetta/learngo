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
}
