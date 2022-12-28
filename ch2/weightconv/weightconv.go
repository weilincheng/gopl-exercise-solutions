// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package lengthconv performs feet and meters conversions.
package weightconv

import "fmt"

type Pounds float64
type Kilograms float64

func (p Pounds) String() string    { return fmt.Sprintf("%g pounds", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%g kilograms", k) }

//!-
