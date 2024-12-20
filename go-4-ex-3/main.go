package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - (4 * a * c)
}

func computeQuadraticFormula(a, b, d float64) (float64, float64, bool) {
	if d == 0 {
		fmt.Print("Es gibt eine lösung")
		x := ((0 - b) + math.Sqrt(d)) / 2 * a
		return x, x, false

	}
	if d < 0 {
		fmt.Print("Es gibt keine lösung")
		return 0.0, 0.0, true
	}
	if d > 0 {
		fmt.Print("Es gibt mehrere lösungen")
		x1 := ((0 - b) + math.Sqrt(d)) / (2 * a)
		x2 := ((0 - b) - math.Sqrt(d)) / (2 * a)
		return x1, x2, true
	}
	return 0.0, 0.0, false
}
func main() {
	// TODO: call the function computeQuadraticFormula
	a := 1.0
	b := 2.0
	c := 1.0
	d := computeDiscriminant(a, b, c)
	fmt.Printf("Diskriminante: %.2f\n", d)

	x1, x2, hasSolution := computeQuadraticFormula(a, b, d)
	if hasSolution {
		fmt.Printf("Lösung(en): x1 = %.2f, x2 = %.2f\n", x1, x2)
	}
}
