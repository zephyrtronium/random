package random

import (
	"math"
	"testing"
)

var _ = math.E

const (
	two64           = 1 << 64
	two53           = 1 << 53
	two1022 float64 = 4.494232837155789769323262976972561834044942447355766431835752028943316895137524e307
	two991  float64 = 2.092790248410678361227392673945316036252743772862370327038574977285841896728391e298

	mantissa      = 0x000fffffffffffff
	expo1To2      = 0x3ff0000000000000
	zeroToTwoM991 = 0x01ffffffffffffff
)

type Gen func(uint64) float64

func NonIDDivision(x uint64) float64 {
	return float64(x) / two64
}

func IDDivision(x uint64) float64 {
	return float64(x&(two53-1)) / two53
}

func SetExponent1(x uint64) float64 {
	return math.Float64frombits(x&mantissa|expo1To2) - 1
}

func SetExponent2(x uint64) float64 {
	return math.Float64frombits(x&mantissa) * two1022
}

func RandExponent1(x uint64) float64 {
	x &= zeroToTwoM991
	e := x >> 52
	x &^= 1<<(32-e-1) - 1
	return math.Float64frombits(x) * two991
}

var corpus = []uint64{0, 1, 2, 1<<32 - 1, 1 << 32, 1<<32 + 1, 1<<53 - 1, 1 << 53, 1<<53 + 1, 1<<64 - 3, 1<<64 - 2, 1<<64 - 1}

func bench(b *testing.B, g Gen) {
	for i := 0; i < b.N; i++ {
		for _, v := range corpus {
			_ = g(v)
		}
	}
}

func BenchmarkNonIDDivision(b *testing.B) {
	bench(b, NonIDDivision)
}

func BenchmarkIDDivision(b *testing.B) {
	bench(b, IDDivision)
}

func BenchmarkSetExponent1(b *testing.B) {
	bench(b, SetExponent1)
}

func BenchmarkSetExponent2(b *testing.B) {
	bench(b, SetExponent2)
}

func BenchmarkRandExponent1(b *testing.B) {
	bench(b, RandExponent1)
}
