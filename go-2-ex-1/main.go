package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type Birthday struct {
	DayOB   byte
	MonthOB byte
	YearOB  int16
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	DateOfBirth      Birthday
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			FirstName: "Arber",
			LastName:  "idrizaj",
		},
		DateOfBirth: Birthday{
			DayOB:   15,
			MonthOB: 2,
			YearOB:  2007,
		},
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       '=', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
