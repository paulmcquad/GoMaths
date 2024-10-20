package plane

// fmt package provides the function to print anything
import (
	"fmt"
	"math"
)

// Plane - Menu
func Menu() {
	var x1, y1 float64
	var x2, y2 float64
	var Radius float64
	var radius float64
	var menu int64

	fmt.Println("2D Shapes:")
	fmt.Println("1 - Area of a Circle")
	fmt.Println("2 - Area of an Annulus")
	fmt.Println("3 - Circumference of a Circle")
	fmt.Println("4 - Midpoint of a line")
	fmt.Println("5 - Slope of a line")
	fmt.Println("6 - Distance")
	fmt.Println("Input No:")

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
	case 4:
		midpoint(x1, y1, x2, y2)
		break
	case 5:
		slope(x1, y1, x2, y2)
		break
	case 6:
		distance(x1, y1, x2, y2)
		break

	}
}

// Find the Area of an Annulus
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

// Go program for midpoint of a line
func midpoint(x1, y1, x2, y2 float64) float64 {

	fmt.Printf("Type x1, y1: ")
	fmt.Scan(&x1, &y1)
	fmt.Printf("Type x2, y2: ")
	fmt.Scan(&x2, &y2)

	fmt.Printf("Midpoint: %g, %g\n", ((x1 + x2) / 2), ((y1 + y2) / 2))
	return 0
}

// Find the slope of a line
func slope(x1, y1, x2, y2 float64) float64 {

	fmt.Printf("Type x1, y1: ")
	fmt.Scan(&x1, &y1)
	fmt.Printf("Type x2, y2: ")
	fmt.Scan(&x2, &y2)

	if x2-x1 != 0 {
		//return (float64)(y2-y1) / (x2 - x1)
		fmt.Printf("Scope is : %g\n", ((y2 - y1) / (x2 - x1)))
	}
	return 0
}

// Find the distance
func distance(x1, y1, x2, y2 float64) float64 {

	fmt.Printf("Type x1, y1: ")
	fmt.Scan(&x1, &y1)
	fmt.Printf("Type x2, y2: ")
	fmt.Scan(&x2, &y2)

	if x2-x1 != 0 {
		fmt.Printf(
			"Distance is : %g\n",
			math.Sqrt((math.Pow(y2-y1, 2) + math.Pow(x2-x1, 2))))
	}
	return 0
}
