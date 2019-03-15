package main

import (
	"fmt"
)

func main() {
	// var (
	// 	speed int
	// 	now   time.Time
	// )

	// // For readability do not use this assignment for more than 3 variables
	// speed, now = 100, time.Now()

	// fmt.Println(speed, now)

	var (
		speed     = 100
		prevSpeed = 50
	)

	speed, prevSpeed = prevSpeed, speed

	fmt.Println(speed, prevSpeed)
}
