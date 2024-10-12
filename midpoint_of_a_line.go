package main

import "fmt"

func main() {
	var x1, y1 float64
	var x2, y2 float64

	fmt.Printf("Type x1, y1: ")
	fmt.Scan(&x1, &y1)
	fmt.Printf("Type x2, y2: ")
	fmt.Scan(&x2, &y2)

	midpoint(x1, y1, x2, y2)
}

// Go program for midpoint of a line
func midpoint(x1, y1, x2, y2 float64) float64 {

	fmt.Printf("Midpoint: %g, %g\n", ((x1 + x2) / 2), ((y1 + y2) / 2))
	return 0
}
