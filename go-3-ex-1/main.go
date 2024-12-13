package main

import "fmt"

const (
	Aries       = '\u2648' // Widder 21.03. – 20.04
	Taurus      = '\u2649' // Stier	21.04. – 20.05.
	Gemini      = '\u264a' // Zwillinge	21.05. – 21.06.
	Cancer      = '\u264b' // Krebs	22.06. – 22.07.
	Leo         = '\u264c' // Löwe		23.07. – 23.08.
	Virgo       = '\u264d' // Jungfrau	24.08. – 23.09.
	Libra       = '\u264e' // Waage	24.09. – 23.10.
	Scorpius    = '\u264f' // Skorpion		24.10. – 22.11
	Sagittarius = '\u2650' // Schütze	23.11. – 21.12.
	Capricornus = '\u2651' // Steinbock 	22.12. – 20.01.
	Aquarius    = '\u2652' // Wassermann		21.01. – 19.02.
	Pisces      = '\u2653' // Fische 	20.02. – 20.03
)

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'
	if (p.Month == 3 && p.Day >= 21) || (p.Month == 4 && p.Day <= 20) {
		zodiacSign = Aries
	}
	if (p.Month == 4 && p.Day >= 21) || (p.Month == 5 && p.Day <= 20) {
		zodiacSign = Taurus
	}
	if (p.Month == 5 && p.Day >= 21) || (p.Month == 6 && p.Day <= 21) {
		zodiacSign = Gemini
	}
	if (p.Month == 6 && p.Day >= 22) || (p.Month == 7 && p.Day <= 22) {
		zodiacSign = Cancer
	}
	if (p.Month == 7 && p.Day >= 23) || (p.Month == 8 && p.Day <= 23) {
		zodiacSign = Leo
	}
	if (p.Month == 8 && p.Day >= 24) || (p.Month == 9 && p.Day <= 23) {
		zodiacSign = Virgo
	}
	if (p.Month == 9 && p.Day >= 24) || (p.Month == 10 && p.Day <= 23) {
		zodiacSign = Libra
	}
	if (p.Month == 10 && p.Day >= 24) || (p.Month == 11 && p.Day <= 22) {
		zodiacSign = Scorpius
	}
	if (p.Month == 11 && p.Day >= 23) || (p.Month == 12 && p.Day <= 21) {
		zodiacSign = Sagittarius
	}
	if (p.Month == 12 && p.Day >= 22) || (p.Month == 1 && p.Day <= 20) {
		zodiacSign = Capricornus
	}
	if (p.Month == 1 && p.Day >= 21) || (p.Month == 2 && p.Day <= 19) {
		zodiacSign = Aquarius
	}
	if (p.Month == 2 && p.Day >= 20) || (p.Month == 3 && p.Day <= 20) {
		zodiacSign = Pisces
	}

	// TODO: Assign proper value to zodiacSign using if/else branching.
	// NOTE: The runes are defined above as constants.

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)
}
