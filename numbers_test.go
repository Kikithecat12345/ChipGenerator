package ChipGenerator

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsMap = map[int]string{
	1:    "million",
	2:    "billion",
	3:    "trillion",
	4:    "quadrillion",
	5:    "quintillion",
	6:    "sextillion",
	7:    "septillion",
	8:    "octillion",
	9:    "nonillion",
	10:   "decillion",
	20:   "vigintillion",
	30:   "trigintillion",
	40:   "quadragintillion",
	50:   "quinquagintillion",
	60:   "sexagintillion",
	70:   "septuagintillion",
	80:   "octogintillion",
	90:   "nonagintillion",
	100:  "centillion",
	200:  "ducentillion",
	300:  "trecentillion",
	400:  "quadringentillion",
	500:  "quingentillion",
	600:  "sescentillion",
	700:  "septingentillion",
	800:  "octingentillion",
	900:  "nongentillion",
	1000: "millinillion",
}

func g(t *testing.T, number, name string) {
	var a big.Int
	a.SetString(number, 10)
	h(t, &a, name)
}

func h(t *testing.T, number *big.Int, name string) {
	a := CustomChip{
		Magnitude:   number,
		StartDigits: 0,
		Colors:      []string{},
	}
	assert.Equal(t, name, a.GenerateIllion())
}

func TestIllionGenerator(t *testing.T) {
	for k, v := range testsMap {
		g(t, strconv.Itoa(k), v)
	}
	g(t, "389457", "novemoctogintatrecentiseptenquinquagintaquadringentillion")
}

func BenchmarkIllionGenerator(b *testing.B) {
	var a big.Int
	a.SetString("69696969696969696969", 10)
	var c CustomChip
	c.Magnitude = &a
	for i := 0; i < b.N; i++ {
		_ = c.GenerateIllion()
	}
}
