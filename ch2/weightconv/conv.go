// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package weightconv

// CToF converts a Celsius temperature to Fahrenheit.
func PToK(p Pounds) Kilograms { return Kilograms(p / 0.45359237) }

// FToC converts a Fahrenheit temperature to Celsius.
func KToP(k Kilograms) Pounds { return Pounds(k * 0.45359237) }

//!-
