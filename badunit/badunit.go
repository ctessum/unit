package badunit

import "github.com/ctessum/unit"

// HorsePower creates a new unit from an amount of horsepower hp.
func HorsePower(hp float64) *unit.Unit {
	return unit.New(hp*745.699872, unit.Watt)
}

// Ton creates a new unit from a number of short tons t.
func Ton(t float64) *unit.Unit {
	return unit.New(t*907.185, unit.Kilogram)
}
