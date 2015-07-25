package unit

import "github.com/gonum/unit"

var (
	Unitless = unit.Dimensions{}
	Joule    = unit.Dimensions{
		unit.MassDim:   1,
		unit.LengthDim: 2,
		unit.TimeDim:   -2,
	}
	Meter = unit.Dimensions{
		unit.LengthDim: 1,
	}
	// Meter2 is a square meter
	Meter2 = unit.Dimensions{
		unit.LengthDim: 2,
	}
	// Meter3 is a cubic meter
	Meter3 = unit.Dimensions{
		unit.LengthDim: 3,
	}
	// KilogramPerMeter3 is density.
	KilogramPerMeter3 = unit.Dimensions{
		unit.MassDim:   1,
		unit.LengthDim: -3,
	}
	// Pascal is a unit of pressure [kg m-1 s-2]
	Pascal = unit.Dimensions{
		unit.MassDim:   1,
		unit.LengthDim: -1,
		unit.TimeDim:   -2,
	}
	Kilogram = unit.Dimensions{
		unit.MassDim: 1,
	}

	Watt = unit.Dimensions{
		unit.MassDim:   1,
		unit.LengthDim: 2,
		unit.TimeDim:   -3,
	}
	Herz = unit.Dimensions{
		unit.TimeDim: -1,
	}
	Second = unit.Dimensions{
		unit.TimeDim: 1,
	}
)
