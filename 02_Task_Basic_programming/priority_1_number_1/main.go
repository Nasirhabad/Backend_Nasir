package main

import "fmt"

func main() {
	// Constants
	const Phi = 3.14

	// Given dimensions
	radius := 5.0
	height := 10.0

	// Calculate volume of the tube
	volume := Phi * radius * radius * height

	// Output the result
	fmt.Printf("The volume of the tube with radius %.1f units and height %.1f units is: %.2f cubic units\n", radius, height, volume)
}
