package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

func main() {
	var a big.Int
	a.SetString(strings.Repeat("69", 1000), 10)
	n := time.Now()
	z := GenerateIllion(&a)
	n2 := time.Since(n)
	fmt.Println(z)
	fmt.Println(n2)
}
