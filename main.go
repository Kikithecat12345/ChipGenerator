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

// == variables ==S
var prefixes = map[int]string{
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
var littlePrefixes = map[int]string{
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

func main() {
	fmt.Println(generateIllion(big.NewInt(471)))
	fmt.Println(generateIllion(big.NewInt(1984)))
	fmt.Println(generateIllion(big.NewInt(981)))
	fmt.Println(generateIllion(big.NewInt(3485793)))
}

// this function takes a number and returns a string with the number in illion form, where the number is the illion in the sequence of illions.
// takes in a big.Int and returns a string.
// examples: 1 -> "million", 10 -> "decillion", 24 -> "quattorvigintillion" etc.
func generateIllion(illn *big.Int) string {
	str := illn.String()
	// pad the start of the string with 0s so that it's divisible by 3.
	for len(str)%3 != 0 {
		str = "0" + str
	}
	strLen := len(str)
	var lastPrefix int
	var illionWord string

	// iterate in reverse order
	for i := strLen - 1; i >= 0; i-- {
		iMod := i % 3
		// are we at the start of a new group of 3 digits, and are they all 0?
		if iMod == 2 && str[i-2:i+1] == "000" {
			// if all the digits are 0, it's "nilli"
			illionWord = "nilli" + illionWord
		} else {
			switch iMod {
			case 0: // hundreds digit
				if str[i] != '0' {
					lastPrefix = int(str[i]-'0') * 100
					illionWord = prefixes[lastPrefix] + illionWord
				}
			case 1: // tens digit
				if str[i] != '0' {
					lastPrefix = int(str[i]-'0') * 10
					illionWord = prefixes[lastPrefix] + illionWord
				}
			case 2: // ones digit
				// if tens and hundreds digits are 0, we use the littlePrefixes map instead of the prefixes map.
				if str[i-1] == '0' && str[i-2] == '0' {
					lastPrefix = int(str[i] - '0')
					illionWord = littlePrefixes[lastPrefix] + illionWord
				} else {
					// we use the prefixes map, but account for english grammar rules.
					if str[i] == '0' {
						continue
					} else if (str[i] == '3' || str[i] == '6') && (lastPrefix == 20 || lastPrefix == 30 || lastPrefix == 40 || lastPrefix == 50 || lastPrefix == 300 || lastPrefix == 400 || lastPrefix == 500) {
						illionWord = prefixes[int(str[i]-'0')] + "s" + illionWord
					} else if str[i] == 6 && (lastPrefix == 80 || lastPrefix == 100 || lastPrefix == 800) {
						illionWord = "sex" + illionWord
					} else if (str[i] == 7 || str[i] == 9) && (lastPrefix == 20 || lastPrefix == 80 || lastPrefix == 800) {
						illionWord = prefixes[int(str[i]-'0')] + "n" + illionWord
					} else if (str[i] == 7 || str[i] == 9) && (lastPrefix == 10 || lastPrefix == 30 || lastPrefix == 40 || lastPrefix == 50 || lastPrefix == 60 || lastPrefix == 70 || lastPrefix == 100 || lastPrefix == 200 || lastPrefix == 300 || lastPrefix == 400 || lastPrefix == 500 || lastPrefix == 600 || lastPrefix == 700) {
						illionWord = prefixes[int(str[i]-'0')] + "m" + illionWord
					} else {
						illionWord = prefixes[int(str[i]-'0')] + illionWord
					}
				}
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
	return illionWord
}
