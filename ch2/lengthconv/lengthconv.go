// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package lengthconv performs feet and meters conversions.
package lengthconv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }
func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }

//!-
