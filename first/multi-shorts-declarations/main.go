package main

import "fmt"

func main() {

	name, lastname := "Nikola", "Tesla"
	fmt.Println(name, lastname)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	degree, ratio, heat := 10.55, 30.5, 20.
	fmt.Println(degree, ratio, heat)

	nFiles, valid, msg := 10, true, "hello"
	fmt.Println(nFiles, valid, msg)

	// --- TIPS ---
	// If you know the value of variable use the short declaration
	// var width, height = 100, 50 // DON'T!
	width, height := 100, 50

	fmt.Println(width, height)

	// --- TIPS ---
	// Use short declaration to keep you code concise
}
