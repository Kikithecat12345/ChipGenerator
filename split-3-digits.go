package main

import (
	"math/big"
	"strconv"
)

func MathCeilInt(n, div int) int {
	if n%div == 0 {
		return n / div
	}
	return n/div + 1
}

func EndSplitN(number string) []string {
	l := MathCeilInt(len(number), 3)
	z := make([]string, l)
	n := len(number) % 3
	if n == 0 {
		n = 3
	}
	for i := 0; i < l; i++ {
		if i == 0 {
			z[0] = number[:n]
		} else {
			z[i] = number[i*3-3+n : i*3+n]
		}
	}
	return z
}

func Split3Digits(number *big.Int) []int {
	n := EndSplitN(number.String())
	l := len(n)
	z := make([]int, l)
	for i := range n {
		// ignoring the error because big.Int.String() only outputs numeric digits
		z[i], _ = strconv.Atoi(n[i])
	}
	return z
}
