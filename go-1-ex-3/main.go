package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	eyesFile, _ := os.Create("eyes.txt")
	defer eyesFile.Close()
	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")

	diceLogFile, _ := os.Create("dice.log")
	defer diceLogFile.Close()
	fmt.Fprintln(diceLogFile, "the dice was rolled at", when)
}
