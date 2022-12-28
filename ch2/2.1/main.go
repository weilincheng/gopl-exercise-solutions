package main

import (
	"fmt"

	"gopl_exercise/ch2/tempconv"
)

func main() {
	c := tempconv.KToC(tempconv.ZeroK)
	fmt.Println(c.String())

	k := tempconv.CToK(tempconv.AbsoluteZeroC)
	fmt.Println(k.String())
}
