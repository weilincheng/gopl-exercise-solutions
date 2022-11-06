package main

import (
	"fmt"

	"exercises2.1/tempconv"
)

func main() {
	c := tempconv.KToC(tempconv.ZeroK)
	fmt.Println(c.String())

	k := tempconv.CToK(tempconv.AbsoluteZeroC)
	fmt.Println(k.String())
}
