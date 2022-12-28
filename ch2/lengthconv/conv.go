// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package lengthconv

// CToF converts a Celsius temperature to Fahrenheit.
func FToM(f Feet) Meters { return Meters(f * 0.3048) }

// FToC converts a Fahrenheit temperature to Celsius.
func MToF(m Meters) Feet { return Feet(m / 0.3048) }

//!-
