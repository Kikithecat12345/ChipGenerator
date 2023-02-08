package main

/*
This code (and it's accompanying Lua script) calculates, generates, and imports custom poker chips of arbitrary denominations and colors into Tabletop Simulator.
The code is written in Go, except for the import function which is written in Lua, because the API for TTS uses Lua exclusively.
The code is written in a way that allows for easy modification of the chip denominations, colors, and other parameters via a config file.
To support arbirary chip denominations, the code uses math/big to handle arbitrary precision numbers, and then converts them to strings for the generation of the chip images.
*/

import (
	"fmt"
	"math/big"
)

var (
	// == variables ==S
	prefixes = map[int]string{
		0:   "nilli",
		1:   "un",
		2:   "duo",
		3:   "tre",
		4:   "quattuor",
		5:   "quinqua",
		6:   "se",
		7:   "septe",
		8:   "octo",
		9:   "nove",
		10:  "deci",
		20:  "viginti",
		30:  "triginta",
		40:  "quadraginta",
		50:  "quinquaginta",
		60:  "sexaginta",
		70:  "septuaginta",
		80:  "octoginta",
		90:  "nonaginta",
		100: "centi",
		200: "ducenti",
		300: "trecenti",
		400: "quadragcenti",
		500: "quingenti",
		600: "sescenti",
		700: "septingenti",
		800: "octingenti",
		900: "nongenti",
	}
	littlePrefixes = map[int]string{
		1:  "milli", // 1
		2:  "billi",
		3:  "trilli",
		4:  "quadrilli",
		5:  "quintilli",
		6:  "sextilli",
		7:  "septilli",
		8:  "octilli",
		9:  "nonilli",
		10: "decilli", // 10
	}
)

// GenerateIllion takes a number and returns a string with the number in illion form, where the number is the illion in the sequence of illions.
// takes in a big.Int and returns a string.
// examples: 1 -> "million", 10 -> "decillion", 24 -> "quattorvigintillion" etc.
func GenerateIllion(illion *big.Int) string {
	// we need to convert the big.Int to an array of ints, each item in the array is 3 digits of the number. if the number is not divisible by 3, we prioritize the least significant digits.
	// for example, 23498237491 -> [23, 498, 237, 491]
	digits := Split3Digits(illion)

	illionWord := ""
	// now we iterate through the array and generate the illion word using the prefixes map, and the littlePrefixes map.
	for _, v := range digits {
		if v == 0 { // is it 0?
			illionWord += prefixes[0]
			continue
		}
		// if it's less than 10, we have to use the littlePrefixes map.
		if v < 10 {
			illionWord += littlePrefixes[v]
			continue
		}
		// now we need to treat the 3 digits seperately, starting with the most significant digit, the hundreds digit.
		var lastIllion int

		if v/100 != 0 {
			lastIllion = v / 100 * 100
			illionWord += prefixes[lastIllion]

		}
		if v%100 != 0 {
			lastIllion = v % 100 / 10 * 10
			fmt.Println("10s", lastIllion, prefixes[lastIllion])
			illionWord += prefixes[lastIllion]
		}
		/*
			the ones case is special, because english is weird.
			basically, we have to modify the ones digit to account for english grammar.
			"septe" and "nove" might need to get an M or N added to the end of them, depending on the context.
			"se" gets an S or X sometimes, and "tre" gets an S as well.
			all of this information is at https://en.wikipedia.org/wiki/Names_of_large_numbers#Extensions_of_the_standard_dictionary_numbers.
		*/
		onesDigit := v % 10
		if onesDigit != 0 {
			if (onesDigit == 3 || onesDigit == 6) && (lastIllion == 20 || lastIllion == 30 || lastIllion == 40 || lastIllion == 50 || lastIllion == 300 || lastIllion == 400 || lastIllion == 500) {
				// tres / ses case
				illionWord += prefixes[onesDigit] + "s"
			} else if onesDigit == 6 && (lastIllion == 80 || lastIllion == 100 || lastIllion == 800) {
				// sex case
				illionWord += "sex"
			} else if (onesDigit == 7 || onesDigit == 9) && (lastIllion == 20 || lastIllion == 80 || lastIllion == 800) {
				// septen / noven case
				illionWord += prefixes[onesDigit] + "n"
			} else if (onesDigit == 7 || onesDigit == 9) && (lastIllion == 10 || lastIllion == 30 || lastIllion == 40 || lastIllion == 50 || lastIllion == 60 || lastIllion == 70 || lastIllion == 100 || lastIllion == 200 || lastIllion == 300 || lastIllion == 400 || lastIllion == 500 || lastIllion == 600 || lastIllion == 700) {
				// septem / novem case
				illionWord += prefixes[onesDigit] + "m"
			} else {
				illionWord += prefixes[onesDigit]
			}
		}
	}
	// add the "illion" suffix.
	// however, if it ends in "illi" we only add "on", for example "milli" -> "million
	// and if it only ends in a vowel we remove it, then add "illion". for example "quadraginta" -> "quadragintillion"
	if illionWord[len(illionWord)-4:] == "illi" {
		illionWord += "on"
	} else if illionWord[len(illionWord)-1:] == "a" || illionWord[len(illionWord)-1:] == "e" || illionWord[len(illionWord)-1:] == "i" || illionWord[len(illionWord)-1:] == "o" || illionWord[len(illionWord)-1:] == "u" {
		illionWord = illionWord[:len(illionWord)-1] + "illion"
	} else {
		illionWord += "illion"
	}
	return illionWord // @MrMelon54 for the love of god please optimize this function, i don't know what i'm doing
}
