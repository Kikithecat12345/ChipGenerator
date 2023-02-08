package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func g(t *testing.T, number, name string) {
	var a big.Int
	a.SetString(number, 10)
	assert.Equal(t, name, GenerateIllion(&a))
}

func TestIllionGenerator(t *testing.T) {
	g(t, "1", "million")
	g(t, "10", "decillion")
	g(t, "100", "centillion")
	g(t, "1000", "millinillion")
	g(t, "389457", "treoctogintanongentiquattorquinquagintaseptingentillion")
}
