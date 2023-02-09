package main

import (
	"fmt"
	"github.com/Kikithecat12345/ChipGenerator"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

var bi1 = big.NewInt(1)
var bi10 = big.NewInt(10)

func main() {
	n := big.NewInt(1)
	for {
		fmt.Printf("Running test for %s digits\n", n)
		run(genNumber(n))
		n.Mul(n, bi10)
	}
}

func run(n string) {
	var a big.Int
	a.SetString(n, 10)
	t := time.Now()
	_ = ChipGenerator.GenerateIllion(&a)
	n2 := time.Since(t)
	fmt.Println(n2)
}

func genNumber(n *big.Int) string {
	end := new(big.Int).Set(n)
	b := ""
	for i := big.NewInt(0); i.Cmp(end) < 0; i.Add(i, bi1) {
		b += strconv.Itoa(rand.Intn(10))
	}
	return b
}
