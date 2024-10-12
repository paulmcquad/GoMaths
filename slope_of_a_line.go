package main

import "fmt"

func main() {
	var x1, y1 float64
	var x2, y2 float64
	var ret float64

	fmt.Printf("Type x1, y1: ")
	fmt.Scan(&x1, &y1)
	fmt.Printf("Type x2, y2: ")
	fmt.Scan(&x2, &y2)

	ret = slope(x1, y1, x2, y2)

	fmt.Printf("Scope is : %g\n", ret)
}

// Go program for slope of a line
func slope(x1, y1, x2, y2 float64) float64 {
	if x2-x1 != 0 {
		return (float64)(y2-y1) / (x2 - x1)
	}
	return 0
}
