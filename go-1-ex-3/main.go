package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	eyes := rand.Intn(6) + 1
	when := time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	fmt.Fprintln(createFile("eyes.txt"), "the dice shows", eyes, "eyes")
	fmt.Fprintln(createFile("dice.log"), "the dice was rolled at", when)
}

func createFile(name string) *os.File {
	file, _ := os.Create(name)
	return file
}
