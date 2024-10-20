package main

// fmt package provides the function to print anything
import (
	"fmt"
	"plane/plane"
)

func main() {

	// declaring the floating variables using the var keyword for
	// storing the radius of the circle also a variable area to store Area
	var Radius float64
	var radius float64

	fmt.Println("Annulus Calculator.")

	// initializing the radius of a circle
	fmt.Printf("Outer Radius = ")
	fmt.Scan(&Radius)

	// initializing the radius of a circle
	fmt.Printf("Inner Radius = ")
	fmt.Scan(&radius)

	// finding the Area of a Circle
	plane.Annulus(Radius, radius)
}
