package plane

// fmt package provides the function to print anything
import (
	"fmt"
	"math"
)

// Plane - Menu
func Menu() {
	var Radius float64
	var radius float64
	var menu int64

	fmt.Println("2D Shapes:")
	fmt.Println("1 - Area of a Circle")
	fmt.Println("2 - Area of an Annulus")
	fmt.Println("3 - Circumference of a Circle")

	fmt.Scan(&menu)

	switch menu {
	case 1:
		AreaOfCircle(radius)
		break
	case 2:
		Annulus(Radius, radius)
		break
	case 3:
		CircumferenceOfCircle(radius)
		break
	}
}

// finding the Area of an Annulus
func Annulus(Radius float64, radius float64) float64 {

	fmt.Println("Annulus Calculator.")

	// initializing the radius of the outer circle
	fmt.Printf("Outer Radius = ")
	fmt.Scan(&Radius)

	// initializing the radius of the inner circle
	fmt.Printf("Inner Radius = ")
	fmt.Scan(&radius)

	var Area float64 = math.Pi * (math.Pow(Radius, 2) - math.Pow(radius, 2))

	fmt.Println("Annulus Area: ", Area)
	return 0
}

// Find the Area of a Circle
func AreaOfCircle(radius float64) float64 {

	fmt.Println("Program to find the Area of a Circle.")

	// initializing the radius of a circle
	fmt.Printf("Radius = ")
	fmt.Scan(&radius)

	// returning the area by applying the formula
	var Area float64 = math.Pi * (radius * radius)
	fmt.Println("The Area of the circle =", Area, "cm^2")

	return 0
}

// Find the Circumference of a Circle
func CircumferenceOfCircle(radius float64) float64 {

	fmt.Println("Program to find the Circumference of a Circle.")

	// initializing the radius of a circle
	fmt.Printf("Radius = ")
	fmt.Scan(&radius)

	var Circumference float64 = 2 * math.Pi * radius

	fmt.Println("The Circumference of the circle =", Circumference, "cm^2")
	return 0
}
