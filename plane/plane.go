package plane

// fmt package provides the function to print anything
import (
	"fmt"
	"math"
)

func Annulus(Radius float64, radius float64) float64 {
	// returning the area by applying the formula
	var Area float64

	Area = math.Pi * (math.Pow(Radius, 2) - math.Pow(radius, 2))

	fmt.Println("Annulus Area: ", Area)
	return 0
}
