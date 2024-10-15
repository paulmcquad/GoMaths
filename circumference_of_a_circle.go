package main

// fmt package provides the function to print anything
import (
	"fmt"
	"math"
)

func main() {

	// declaring the floating variables using the var keyword for
	// storing the radius of the circle also a variable area to store Area
	var radius float64

	fmt.Println("Program to find the Circumference of a Circle.")

	// initializing the radius of a circle
	fmt.Printf("Radius = ")
	fmt.Scan(&radius)

	// finding the Area of a Circle
	CircumferenceOfCircle(radius)
}

func CircumferenceOfCircle(radius float64) float64 {
	// returning the area by applying the formula
	var Circumference float64

	Circumference = 2 * math.Pi * radius

	//Area = math.Pi * (radius * radius)
	fmt.Println("The Circumference of the circle =", Circumference, "cm^2")
	return 0
}
