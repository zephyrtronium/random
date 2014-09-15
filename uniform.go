package random

import "math"

// Uniform1_2 produces random numbers uniformly distributed in the range
// [1, 2). This works by setting the mantissa of the value to random bits and
// the exponent to a particular value, making it faster than Uniform.
type Uniform1_2 struct{}

func (Uniform1_2) Next(src Source) float64 {
	return math.Float64frombits(src.Uint64()&0x000fffffffffffff | 0x3ff0000000000000)
}

// Uniform produces random numbers uniformly distributed in the range
// [Low, High).
type Uniform struct {
	Low, High float64
}

func (u Uniform) Next(src Source) float64 {
	return (Uniform1_2{}.Next(src)-1)*u.High + u.Low
}
