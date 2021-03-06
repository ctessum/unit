// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"fmt"
	"math"
	"testing"
)

var formatTests = []struct {
	unit   *Unit
	format string
	expect string
}{
	{New(9.81, Dimensions{MassDim: 1, TimeDim: -2}), "%f", "9.81 kg s^-2"},
	{New(9.81, Dimensions{MassDim: 1, TimeDim: -2}), "%1.f", "10 kg s^-2"},
	{New(9.81, Dimensions{MassDim: 1, TimeDim: -2}), "%.1f", "9.8 kg s^-2"},
	{New(9.81, Dimensions{MassDim: 1, TimeDim: -2, LengthDim: 0}), "%f", "9.81 kg s^-2"},
	{New(6.62606957e-34, Dimensions{MassDim: 2, TimeDim: -1}), "%e", "6.62606957e-34 kg^2 s^-1"},
	{New(6.62606957e-34, Dimensions{MassDim: 2, TimeDim: -1}), "%.3e", "6.626e-34 kg^2 s^-1"},
	{New(6.62606957e-34, Dimensions{MassDim: 2, TimeDim: -1}), "%v", "6.62606957e-34 kg^2 s^-1"},
	{New(6.62606957e-34, Dimensions{MassDim: 2, TimeDim: -1}), "%s", "%!s(*Unit=6.62606957e-34 kg^2 s^-1)"},
	{New(math.E, Dimless), "%v", "2.718281828459045"},
	{New(math.E, Dimless), "%#v", "unit.Dimless(2.718281828459045)"},
	{New(math.E, Kilogram), "%s", "%!s(unit.Dimless=2.718281828459045)"},
	{New(1, Kilogram), "%v", "1 kg"},
	{New(1, Kilogram), "%#v", "unit.Mass(1)"},
	{New(1, Kilogram), "%s", "%!s(unit.Mass=1 kg)"},
	{New(1.61619926e-35, Meter), "%v", "1.61619926e-35 m"},
	{New(1.61619926e-35, Meter), "%#v", "unit.Length(1.61619926e-35)"},
	{New(1.61619926e-35, Meter), "%s", "%!s(unit.Length=1.61619926e-35 m)"},
	{New(15.2, Second), "%v", "15.2 s"},
	{New(15.2, Second), "%#v", "unit.Time(15.2)"},
	{New(15.2, Second), "%s", "%!s(unit.Time=15.2 s)"},
}

func TestFormat(t *testing.T) {
	for _, ts := range formatTests {
		if r := fmt.Sprintf(ts.format, ts.unit); r != ts.expect {
			t.Errorf("Format %q: got: %q expected: %q", ts.format, r, ts.expect)
		}
	}
}

func TestGoStringFormat(t *testing.T) {
	expect1 := `&unit.Unit{dimensions:unit.Dimensions{4:2, 6:-1}, formatted:"", value:6.62606957e-34}`
	expect2 := `&unit.Unit{dimensions:unit.Dimensions{6:-1, 4:2}, formatted:"", value:6.62606957e-34}`
	if r := fmt.Sprintf("%#v", New(6.62606957e-34, Dimensions{MassDim: 2, TimeDim: -1})); r != expect1 && r != expect2 {
		t.Errorf("Format %q: got: %q expected: %q", "%#v", r, expect1)
	}
}

var initializationTests = []struct {
	unit     *Unit
	expValue float64
	expMap   map[Dimension]int
}{
	{New(9.81, Dimensions{MassDim: 1, TimeDim: -2}), 9.81, Dimensions{MassDim: 1, TimeDim: -2}},
	{New(9.81, Dimensions{MassDim: 1, TimeDim: -2, LengthDim: 0, CurrentDim: 0}), 9.81, Dimensions{MassDim: 1, TimeDim: -2}},
}

func TestInitialization(t *testing.T) {
	for _, ts := range initializationTests {
		if ts.expValue != ts.unit.value {
			t.Errorf("Value wrong on initialization: got: %v expected: %v", ts.unit.value, ts.expValue)
		}
		if len(ts.expMap) != len(ts.unit.dimensions) {
			t.Errorf("Map mismatch: got: %#v expected: %#v", ts.unit.dimensions, ts.expMap)
		}
		for key, val := range ts.expMap {
			if ts.unit.dimensions[key] != val {
				t.Errorf("Map mismatch: got: %#v expected: %#v", ts.unit.dimensions, ts.expMap)
			}
		}
	}
}

func TestOp(t *testing.T) {
	var v1 = New(10, Dimensions{LengthDim: 1, TimeDim: -2})
	var v2 = New(2, Dimensions{LengthDim: 1, TimeDim: -2})

	type tester struct {
		f    func(...*Unit) *Unit
		args []*Unit
		Val  float64
		Dims string
	}

	tests := []tester{
		tester{
			f:    Add,
			args: []*Unit{v1, v2},
			Val:  12.,
			Dims: "m s^-2",
		},
		tester{
			f:    Add,
			args: []*Unit{nil, v1, v2},
			Val:  12.,
			Dims: "m s^-2",
		},
		tester{
			f:    Sub,
			args: []*Unit{v1, v2},
			Val:  8.,
			Dims: "m s^-2",
		},
		tester{
			f:    Mul,
			args: []*Unit{v1, v2},
			Val:  20.,
			Dims: "m^2 s^-4",
		},
		tester{
			f:    Mul,
			args: []*Unit{v1, v2},
			Val:  20.,
			Dims: "m^2 s^-4",
		},
		tester{
			f:    Div,
			args: []*Unit{v1, v2},
			Val:  5.,
			Dims: "",
		},
		tester{
			f:    Max,
			args: []*Unit{v1, v2},
			Val:  10.,
			Dims: "m s^-2",
		},
		tester{
			f:    Max,
			args: []*Unit{nil, v1, v2},
			Val:  10.,
			Dims: "m s^-2",
		},
		tester{
			f:    Min,
			args: []*Unit{v1, v2},
			Val:  2.,
			Dims: "m s^-2",
		},
		tester{
			f:    Min,
			args: []*Unit{nil, v1, v2},
			Val:  2.,
			Dims: "m s^-2",
		},
	}
	for _, test := range tests {
		r := test.f(test.args...)
		if r.Value() != test.Val {
			t.Errorf("Got value %v, expected %v.",
				r.Value(), test.Val)
		}
		if r.Dimensions().String() != test.Dims {
			t.Errorf("Got value %v, expected %v.",
				r.Dimensions().String(), test.Dims)
		}
	}
}

func TestOpInPlace(t *testing.T) {
	var v1 = New(10, Dimensions{LengthDim: 1, TimeDim: -2})
	var v2 = New(2, Dimensions{LengthDim: 1, TimeDim: -2})

	type tester struct {
		f    func(*Unit)
		Val  float64
		Dims string
	}

	tests := []tester{
		tester{
			f: func(v1 *Unit) {
				v1.Add(v2)
			},
			Val:  12.,
			Dims: "m s^-2",
		},
		tester{
			f: func(v1 *Unit) {
				v1.Sub(v2)
			},
			Val:  8.,
			Dims: "m s^-2",
		},
		tester{
			f: func(v1 *Unit) {
				v1.Mul(v2)
			},
			Val:  20.,
			Dims: "m^2 s^-4",
		},
		tester{
			f: func(v1 *Unit) {
				v1.Div(v2)
			},
			Val:  5.,
			Dims: "",
		},
	}
	for _, test := range tests {
		vv1 := v1.Clone()
		test.f(vv1)
		if vv1.Value() != test.Val {
			t.Errorf("Got value %v, expected %v.",
				vv1.Value(), test.Val)
		}
		if vv1.Dimensions().String() != test.Dims {
			t.Errorf("Got value %v, expected %v.",
				vv1.Dimensions().String(), test.Dims)
		}
	}
}

var dimensionEqualityTests = []struct {
	name        string
	a           *Unit
	b           *Unit
	shouldMatch bool
}{
	{"same_empty", New(1.0, Dimensions{}), New(1.0, Dimensions{}), true},
	{"same_one", New(1.0, Dimensions{TimeDim: 1}), New(1.0, Dimensions{TimeDim: 1}), true},
	{"same_mult", New(1.0, Dimensions{TimeDim: 1, LengthDim: -2}), New(1.0, Dimensions{TimeDim: 1, LengthDim: -2}), true},
	{"diff_one_empty", New(1.0, Dimensions{}), New(1.0, Dimensions{TimeDim: 1, LengthDim: -2}), false},
	{"diff_same_dim", New(1.0, Dimensions{TimeDim: 1}), New(1.0, Dimensions{TimeDim: 2}), false},
	{"diff_same_pow", New(1.0, Dimensions{LengthDim: 1}), New(1.0, Dimensions{TimeDim: 1}), false},
	{"diff_numdim", New(1.0, Dimensions{TimeDim: 1, LengthDim: 2}), New(1.0, Dimensions{TimeDim: 2}), false},
	{"diff_one_same_dim", New(1.0, Dimensions{LengthDim: 1, TimeDim: 1}), New(1.0, Dimensions{LengthDim: 1, TimeDim: 2}), false},
}

func TestDimensionEquality(t *testing.T) {
	for _, ts := range dimensionEqualityTests {
		if DimensionsMatch(ts.a, ts.b) != ts.shouldMatch {
			t.Errorf("Dimension comparison incorrect for case %s. got: %v, expected: %v", ts.name, !ts.shouldMatch, ts.shouldMatch)
		}
	}
}

type UnitStructer interface {
	UnitStruct() *UnitStruct
}

type UnitStruct struct {
	current     int
	length      int
	luminosity  int
	mass        int
	temperature int
	time        int
	chemamt     int // For mol
	value       float64
}

// Check if the dimensions of two units are the same
func DimensionsMatchStruct(aU, bU UnitStructer) bool {
	a := aU.UnitStruct()
	b := bU.UnitStruct()
	if a.length != b.length {
		return false
	}
	if a.time != b.time {
		return false
	}
	if a.mass != b.mass {
		return false
	}
	if a.current != b.current {
		return false
	}
	if a.temperature != b.temperature {
		return false
	}
	if a.luminosity != b.luminosity {
		return false
	}
	if a.chemamt != b.chemamt {
		return false
	}
	return true
}

func (u *UnitStruct) UnitStruct() *UnitStruct {
	return u
}

func (u *UnitStruct) Add(aU UnitStructer) *UnitStruct {
	a := aU.UnitStruct()
	if !DimensionsMatchStruct(a, u) {
		panic("dimension mismatch")
	}
	u.value += a.value
	return u
}

func (u *UnitStruct) Mul(aU UnitStructer) *UnitStruct {
	a := aU.UnitStruct()
	u.length += a.length
	u.time += a.time
	u.mass += a.mass
	u.current += a.current
	u.temperature += a.temperature
	u.luminosity += a.luminosity
	u.chemamt += a.chemamt
	u.value *= a.value
	return u
}

var u3 *UnitStruct

func BenchmarkAddStruct(b *testing.B) {
	u1 := &UnitStruct{current: 1, chemamt: 5, value: 10}
	u2 := &UnitStruct{current: 1, chemamt: 5, value: 100}
	for i := 0; i < b.N; i++ {
		u2.Add(u1)
	}
}

func BenchmarkMulStruct(b *testing.B) {
	u1 := &UnitStruct{current: 1, chemamt: 5, value: 10}
	u2 := &UnitStruct{mass: 1, time: 1, value: 100}
	for i := 0; i < b.N; i++ {
		u2.Mul(u1)
	}
}

type UnitMapper interface {
	UnitMap() *UnitMap
}

type dimensionMap int

const (
	LengthB dimensionMap = iota
	TimeB
	MassB
	CurrentB
	TemperatureB
	LuminosityB
	ChemAmtB
)

type UnitMap struct {
	dimension map[dimensionMap]int
	value     float64
}

// Check if the dimensions of two units are the same
func DimensionsMatchMap(aU, bU UnitMapper) bool {
	a := aU.UnitMap()
	b := bU.UnitMap()
	if len(a.dimension) != len(b.dimension) {
		panic("Unequal dimension")
	}
	for key, dimA := range a.dimension {
		dimB, ok := b.dimension[key]
		if !ok || dimA != dimB {
			panic("Unequal dimension")
		}
	}
	return true
}

func (u *UnitMap) UnitMap() *UnitMap {
	return u
}

func (u *UnitMap) Add(aU UnitMapper) *UnitMap {
	a := aU.UnitMap()
	if !DimensionsMatchMap(a, u) {
		panic("dimension mismatch")
	}
	u.value += a.value
	return u
}

func (u *UnitMap) Mul(aU UnitMapper) *UnitMap {
	a := aU.UnitMap()
	for key, val := range a.dimension {
		u.dimension[key] += val
	}
	u.value *= a.value
	return u
}

func BenchmarkAddFloat(b *testing.B) {
	a := 0.0
	c := 10.0
	for i := 0; i < b.N; i++ {
		a += c
	}
}

func BenchmarkMulFloat(b *testing.B) {
	a := 0.0
	c := 10.0
	for i := 0; i < b.N; i++ {
		a *= c
	}
}

func BenchmarkAddMapSmall(b *testing.B) {
	u1 := &UnitMap{value: 10}
	u1.dimension = make(map[dimensionMap]int)
	u1.dimension[CurrentB] = 1
	u1.dimension[ChemAmtB] = 5

	u2 := &UnitMap{value: 10}
	u2.dimension = make(map[dimensionMap]int)
	u2.dimension[CurrentB] = 1
	u2.dimension[ChemAmtB] = 5
	for i := 0; i < b.N; i++ {
		u2.Add(u1)
	}
}

func BenchmarkMulMapSmallDiff(b *testing.B) {
	u1 := &UnitMap{value: 10}
	u1.dimension = make(map[dimensionMap]int)
	u1.dimension[LengthB] = 1

	u2 := &UnitMap{value: 10}
	u2.dimension = make(map[dimensionMap]int)
	u2.dimension[MassB] = 1
	for i := 0; i < b.N; i++ {
		u2.Mul(u1)
	}
}

func BenchmarkMulMapSmallSame(b *testing.B) {
	u1 := &UnitMap{value: 10}
	u1.dimension = make(map[dimensionMap]int)
	u1.dimension[LengthB] = 1

	u2 := &UnitMap{value: 10}
	u2.dimension = make(map[dimensionMap]int)
	u2.dimension[LengthB] = 2
	for i := 0; i < b.N; i++ {
		u2.Mul(u1)
	}
}

func BenchmarkMulMapLargeDiff(b *testing.B) {
	u1 := &UnitMap{value: 10}
	u1.dimension = make(map[dimensionMap]int)
	u1.dimension[LengthB] = 1
	u1.dimension[MassB] = 1
	u1.dimension[ChemAmtB] = 1
	u1.dimension[TemperatureB] = 1
	u1.dimension[LuminosityB] = 1
	u1.dimension[TimeB] = 1
	u1.dimension[CurrentB] = 1

	u2 := &UnitMap{value: 10}
	u2.dimension = make(map[dimensionMap]int)
	u2.dimension[MassB] = 1
	for i := 0; i < b.N; i++ {
		u2.Mul(u1)
	}
}

func BenchmarkMulMapLargeSame(b *testing.B) {
	u1 := &UnitMap{value: 10}
	u1.dimension = make(map[dimensionMap]int)
	u1.dimension[LengthB] = 2
	u1.dimension[MassB] = 2
	u1.dimension[ChemAmtB] = 2
	u1.dimension[TemperatureB] = 2
	u1.dimension[LuminosityB] = 2
	u1.dimension[TimeB] = 2
	u1.dimension[CurrentB] = 2

	u2 := &UnitMap{value: 10}
	u2.dimension = make(map[dimensionMap]int)
	u2.dimension[LengthB] = 3
	u2.dimension[MassB] = 3
	u2.dimension[ChemAmtB] = 3
	u2.dimension[TemperatureB] = 3
	u2.dimension[LuminosityB] = 3
	u2.dimension[TimeB] = 3
	u2.dimension[CurrentB] = 3
	for i := 0; i < b.N; i++ {
		u2.Mul(u1)
	}
}
