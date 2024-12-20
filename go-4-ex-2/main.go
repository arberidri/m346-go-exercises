package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypothenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}
func main() {
	// TODO: call the function computeHypotenuse
	a := 4.0
	b := 3.0

	hypotenuse := computeHypothenuse(a, b)
	fmt.Printf("Die Hypotenuse des Dreiecks mit a=%.2f und b=%.2f ist c=%.2f\n", a, b, hypotenuse)

}
