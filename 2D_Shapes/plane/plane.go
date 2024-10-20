package plane

// fmt package provides the function to print anything
import (
	"fmt"
	"math"
)

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
