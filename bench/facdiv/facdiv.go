package main

import (
	"fmt"
	"math/big"
)

const target = 20000

func main() {
	f := new(big.Int)
	c := new(big.Int)
	one := new(big.Int)

	f.SetInt64(1)
	c.SetInt64(1)
	one.SetInt64(1)

	for i := 0; i < target; i++ {
		f.Mul(f, c)
		c.Add(c, one)
	}

	r := new(big.Int)

	for i := target - 1; i != 0; i-- {
		f.QuoRem(f, c, r)
		c.Sub(c, one)
	}

	fmt.Printf("%x", f)
}
