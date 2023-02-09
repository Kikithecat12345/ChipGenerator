package ChipGenerator

/*
This code (and it's accompanying Lua script) calculates, generates, and imports custom poker chips of arbitrary denominations and colors into Tabletop Simulator.
The code is written in Go, except for the import function which is written in Lua, because the API for TTS uses Lua exclusively.
The code is written in a way that allows for easy modification of the chip denominations, colors, and other parameters via a config file.
To support arbirary chip denominations, the code uses math/big to handle arbitrary precision numbers, and then converts them to strings for the generation of the chip images.
*/

import (
	"math/big"
	"strings"
)

// == variables ==
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
	400: "quadringenti",
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

// GenerateIllion takes a number and returns a string with the number in illion form, where the number is the illion in the sequence of illions.
// takes in a big.Int and returns a string.
// examples: 1 -> "million", 10 -> "decillion", 24 -> "quattorvigintillion" etc.
func GenerateIllion(illn *big.Int) string {
	str := illn.String()

	// ignore an empty string
	if len(str) == 0 {
		return ""
	}

	// pad the start of the string with 0s so that it's divisible by 3.
	str = padToMultipleOf3(str)
	strLen := len(str)
	var lastPrefix int
	var illionWord string

	// iterate in reverse order
	// i = hundreds digit for a set of 3
	for i := strLen - 3; i >= 0; i -= 3 {
		// are all the digits in this group 0
		if str[i:i+3] == "000" {
			illionWord = "nilli" + illionWord
			continue
		}

		// if tens and hundreds digits are 0, we use the littlePrefixes map instead of the prefixes map.
		if str[i:i+2] == "00" {
			lastPrefix = int(str[i+2] - '0')
			illionWord = littlePrefixes[lastPrefix] + illionWord
			continue
		}

		// hundreds digit
		if str[i] != '0' {
			lastPrefix = int(str[i]-'0') * 100
			illionWord = prefixes[lastPrefix] + illionWord
		}

		// tens digit
		if str[i+1] != '0' {
			lastPrefix = int(str[i+1]-'0') * 10
			illionWord = prefixes[lastPrefix] + illionWord
		}

		// we use the prefixes map, but account for english grammar rules.
		illionWord = onesDigitPrefix(int(str[i+2]-'0'), lastPrefix) + illionWord
	}

	// add the "illion" suffix.
	// however, if it ends in "illi" we only add "on", for example "milli" -> "million
	// and if it only ends in a vowel we remove it, then add "illion". for example "quadraginta" -> "quadragintillion"
	if strings.HasSuffix(illionWord, "illi") {
		illionWord += "on"
	} else {
		lastIndex := len(illionWord) - 1
		if isVowel(illionWord[lastIndex]) {
			illionWord = illionWord[:lastIndex]
		}
		illionWord += "illion"
	}
	return illionWord
}

func onesDigitPrefix(digit int, lastPrefix int) string {
	switch {
	case digit == 0:
		return ""
	case (digit == 3 || digit == 6) && (lastPrefix == 20 || lastPrefix == 30 || lastPrefix == 40 || lastPrefix == 50 || lastPrefix == 300 || lastPrefix == 400 || lastPrefix == 500):
		return prefixes[digit] + "s"
	case digit == 6 && (lastPrefix == 80 || lastPrefix == 100 || lastPrefix == 800):
		return "sex"
	case (digit == 7 || digit == 9) && (lastPrefix == 20 || lastPrefix == 80 || lastPrefix == 800):
		return prefixes[digit] + "m"
	case (digit == 7 || digit == 9) && (lastPrefix == 10 || lastPrefix == 30 || lastPrefix == 40 || lastPrefix == 50 || lastPrefix == 60 || lastPrefix == 70 || lastPrefix == 100 || lastPrefix == 200 || lastPrefix == 300 || lastPrefix == 400 || lastPrefix == 500 || lastPrefix == 600 || lastPrefix == 700):
		return prefixes[digit] + "n"
	default:
		return prefixes[digit]
	}
}

func padToMultipleOf3(a string) string {
	switch len(a) % 3 {
	case 1:
		return "00" + a
	case 2:
		return "0" + a
	default:
		return a
	}
}

func isVowel(a uint8) bool {
	switch a {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
