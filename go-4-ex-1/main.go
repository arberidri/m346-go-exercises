package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(E, M float64) float64 {
	grade := 1.0 + 5.0*(E/M)
	if grade > 6.0 {
		return 6.0
	}
	if grade < 1.0 {
		return 1.0
	}
	return grade
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0))
	fmt.Println(computeGrade(0, 28.0))
	fmt.Println(computeGrade(28.0, 28.0))
	fmt.Println(computeGrade(-28.0, 28.0))
	// TODO: call the function computeGrade
}
