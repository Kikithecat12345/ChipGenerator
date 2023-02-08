package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func g(t *testing.T, number, name string) {
	var a big.Int
	a.SetString(number, 10)
	h(t, &a, name)
}

func h(t *testing.T, number *big.Int, name string) {
	assert.Equal(t, name, GenerateIllion(number))
}

func TestIllionGenerator(t *testing.T) {
	g(t, "1", "million")
	g(t, "10", "decillion")
	g(t, "100", "centillion")
	g(t, "120", "viginticentillion")
	g(t, "200", "ducentillion")
	g(t, "300", "trecentillion")
	g(t, "400", "quadringentillion")
	g(t, "500", "quingentillion")
	g(t, "600", "sescentillion")
	g(t, "700", "septingentillion")
	g(t, "800", "octingentillion")
	g(t, "900", "nongentillion")
	g(t, "1000", "millinillion")
	g(t, "389457", "novemoctogintatrecentiseptenquinquagintaquadringentillion")
}

func BenchmarkIllionGenerator(b *testing.B) {
	var a big.Int
	a.SetString("69696969696969696969", 10)
	for i := 0; i < b.N; i++ {
		_ = GenerateIllion(&a)
	}
}
