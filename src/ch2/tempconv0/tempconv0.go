// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Const
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF function
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC function
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
