package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"strconv"
	"testing"
)

func TestMathCeilInt(t *testing.T) {
	assert.Equal(t, 0, MathCeilInt(0, 3))
	assert.Equal(t, 1, MathCeilInt(1, 3))
	assert.Equal(t, 1, MathCeilInt(2, 3))
	assert.Equal(t, 1, MathCeilInt(3, 3))
	assert.Equal(t, 2, MathCeilInt(4, 3))
	assert.Equal(t, 2, MathCeilInt(5, 3))
	assert.Equal(t, 2, MathCeilInt(6, 3))
}

func TestEndSplitN(t *testing.T) {
	assert.Equal(t, []string{}, EndSplitN(""))
	assert.Equal(t, []string{"1", "234", "567"}, EndSplitN("1234567"))
	assert.Equal(t, []string{"12", "345", "678"}, EndSplitN("12345678"))
	assert.Equal(t, []string{"123", "456", "789"}, EndSplitN("123456789"))
	assert.Equal(t, []string{"1", "234", "567", "890"}, EndSplitN("1234567890"))
}

func TestSplit3Digits(t *testing.T) {
	assertSplit3Digits(t, []int{0}, "")
	assertSplit3Digits(t, []int{1, 234, 567}, "1234567")
	assertSplit3Digits(t, []int{12, 345, 678}, "12345678")
	assertSplit3Digits(t, []int{123, 456, 789}, "123456789")
	assertSplit3Digits(t, []int{1, 234, 567, 890}, "1234567890")
}

func assertSplit3Digits(t *testing.T, expected []int, input string) {
	var a big.Int
	a.SetString(input, 10)
	assert.Equal(t, expected, Split3Digits(&a))
}

func BenchmarkSplit3Digits(b *testing.B) {
	var a big.Int
	a.SetString("1234567890123456789012345678901234567890", 10)
	for i := 0; i < b.N; i++ {
		_ = Split3Digits(&a)
	}
}

func BenchmarkKikiSplit3(b *testing.B) {
	var a big.Int
	a.SetString("1234567890123456789012345678901234567890", 10)
	for i := 0; i < b.N; i++ {
		_ = aaa(&a)
	}
}

func aaa(illion *big.Int) []int {
	illionStr := illion.String()
	// now we split the string into groups of 3 digits.
	arrTemp := make([]string, len(illionStr)/3+1) // we add 1 to the length of the array to account for the case where the number is not divisible by 3. Potentially overallocating memory here, but it's not a big deal.
	for i := len(illionStr); i > 0; i -= 3 {
		if i-3 > 0 {
			arrTemp[i/3] = illionStr[i-3 : i]
		} else {
			arrTemp[i/3] = illionStr[:i]
		}
	}
	// now we convert the array of strings to an array of ints.
	arr := make([]int, len(arrTemp))
	for i, v := range arrTemp {
		arr[i], _ = strconv.Atoi(v)
	}
	return arr
}
