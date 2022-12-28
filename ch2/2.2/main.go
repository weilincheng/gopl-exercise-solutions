// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl_exercise/ch2/lengthconv"
	"gopl_exercise/ch2/tempconv"
	"gopl_exercise/ch2/weightconv"
)

func main() {
	input := os.Args[1:]
	if len(input) == 0 {
		fmt.Print("Input a number: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = append(input, scanner.Text())
	}
	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		feet := lengthconv.Feet(t)
		meters := lengthconv.Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			feet, lengthconv.FToM(feet), meters, lengthconv.MToF(meters))

		pounds := weightconv.Pounds(t)
		kilograms := weightconv.Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n",
			pounds, weightconv.PToK(pounds), kilograms, weightconv.KToP(kilograms))
	}
}

//!-
