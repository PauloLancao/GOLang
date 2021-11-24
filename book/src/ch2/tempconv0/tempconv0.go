// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Kelvin type
type Kelvin float64

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

// CToK function
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC function
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
