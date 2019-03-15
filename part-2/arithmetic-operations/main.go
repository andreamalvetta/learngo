package main

import "fmt"

func main() {
	fmt.Println("sum: ", 3+2)    // sum: int
	fmt.Println("sum: ", 2+3.5)  // sum: float64
	fmt.Println("dif: ", 3-1)    // difference: int
	fmt.Println("dif: ", 3-0.5)  // difference: float64
	fmt.Println("prod: ", 4*5)   // product: int
	fmt.Println("prod: ", 5*2.5) // difference: float64
	fmt.Println("quot: ", 8/2)   // quotient: int
	fmt.Println("quot: ", 8/1.5) // quotient: float64

	fmt.Println("rem: ", 8%3) // remainder only for ints

	f := 8.0
	fmt.Println("rem: ", int(f)%3)

	var ratio float64 = 3 / 2 // go converts the quotient of 2 ints into float64 under the hood
	fmt.Println(ratio)        // it prints 1, not 1.5

	// ratio = float64(int(3) / int(2)) // this expression is the same as above one
	// ratio = float64(3) / 2 // to fix that is necessary that one of the value has to be a float64
	ratio = 3.0 / 2    // another fix
	fmt.Println(ratio) // it prints 1.5
}
