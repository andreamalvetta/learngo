package main

import "fmt"

func main() {
	speed := 100
	force := 2.5

	speed = int(float64(speed) * force)

	fmt.Println(speed)

	var apple int
	var orange int32

	apple = int(orange)

	// isDelicious := bool(orange) // you can't convert an int into a bool

	orange = 65
	color := string(orange) // you can convert int into string -- it will be 'A' which is number 65
	fmt.Println(color)
	// fmt.Println(string(65.0)) you can't convert float to string
	fmt.Println(string([]byte{104, 105})) // It prints 'hi' because 104 is equal to h and 105 is equal to i

	// to prevent unused variable error
	_ = apple
}
