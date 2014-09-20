package random

import (
	"math"
	"math/big"
)

// An RNG adapts a Source to produce random numbers.
type RNG struct {
	Source
}

// Uint32 produces a uniformly distributed random number in the range
// [0, 1<<32) as a uint32.
func (r RNG) Uint32() uint32 {
	return uint32(r.Uint64())
}

// Int63 produces a uniformly distributed random number in the range
// [0, 1<<63) as an int64.
func (r RNG) Int63() int64 {
	return int64(r.Uint64() >> 1)
}

// Int31 produces a uniformly distributed random number in the range
// [0, 1<<31) as an int32.
func (r RNG) Int31() int32 {
	return int32(r.Uint64() >> 33)
}

// Uint64n produces a uniformly distributed random number in the range [0, max)
// as a uint64.
func (r RNG) Uint64n(max uint64) uint64 {
	bad := 0xffffffffffffffff - 0xffffffffffffffff%max
	x := r.Uint64()
	for x > bad {
		x = r.Uint64()
	}
	return x % max
}

// Uintn produces a uniformly distributed random number in the range [0, max)
// as a uint.
func (r RNG) Uintn(max uint) uint {
	return uint(r.Uint64n(uint64(max)))
}

// Int63n produces a uniformly distributed random number in the range [0, max)
// as an int64. It panics if max <= 0.
func (r RNG) Int63n(max int64) int64 {
	if max <= 0 {
		panic("maximum zero or below")
	}
	return int64(r.Uint64n(uint64(max)))
}

// Int31n produces a uniformly distributed random number in the range [0, max)
// as an int32. It panics if max <= 0.
func (r RNG) Int31n(max int32) int32 {
	if max <= 0 {
		panic("maximum zero or below")
	}
	return int32(r.Uint64n(uint64(max)))
}

// Intn produces a uniformly distributed random number in the range [0, max) as
// an int. It panics if max <= 0.
func (r RNG) Intn(max int) int {
	if max <= 0 {
		panic("maximum zero or below")
	}
	return int(r.Uint64n(uint64(max)))
}

// Big produces a uniformly distributed random positive number up to nbits bits
// in length. It panics if nbits <= 0.
func (r RNG) Big(nbits int) *big.Int {
	if nbits <= 0 {
		panic("maximum zero or below")
	}
	var p []big.Word
	if ^big.Word(0) != 0xffffffff {
		// 64-bit
		n := nbits >> 6
		if nbits&63 != 0 {
			n++
		}
		p = make([]big.Word, n)
		for i := range p {
			p[i] = big.Word(r.Uint64())
		}
		p[0] &= 0xffffffffffffffff >> big.Word(64-nbits&63)
	} else {
		// 32-bit
		n := nbits >> 5
		if nbits&31 != 0 {
			n++
		}
		p = make([]big.Word, n)
		for i := 1; i < len(p); i += 2 {
			// Use each value for two words.
			x := r.Uint64()
			p[i-1] = big.Word(x)
			p[i] = big.Word(x >> 32)
		}
		p[n-1] = big.Word(r.Uint64())
		p[0] &= 0xffffffff >> big.Word(32-nbits&31)
	}
	return new(big.Int).SetBits(p)
}

// Bign produces a uniformly distributed random number in the range [0, max).
// It panics if max <= 0.
func (r RNG) Bign(max *big.Int) *big.Int {
	if max.Sign() <= 0 {
		panic("maximum zero or below")
	}
	nbits := max.BitLen() + 1
	m := new(big.Int)
	m.Sub(m.SetBit(m, nbits, 1), big.NewInt(1))
	k := new(big.Int)
	k.Rem(m, max)
	k.Sub(m, k)
	x := r.Big(nbits)
	for x.Cmp(k) > 0 {
		x = r.Big(nbits)
	}
	return x.Rem(x, max)
}

// Float64 produces a uniformly distributed random number in the range [0, 1).
func (r RNG) Float64() float64 {
	// Converting a full uint64 to float64 introduces rounding error. We avoid
	// this by only taking 53 bits, which is the same as the number of bits in
	// the float64 mantissa including the implicit normalization bit.
	return float64(r.Uint64()&(1<<53-1)) / (1 << 53)
}

// Float32 produces a uniformly distributed random number in the range [0, 1).
func (r RNG) Float32() float32 {
	return float32(r.Uint64()&(1<<24-1)) / (1 << 24)
}

// NormFloat64 produces a normally distributed random number with mean of 0 and
// standard deviation of 1.
func (r RNG) NormFloat64() float64 {
	return genNormal(r)
}

// ExpFloat64 produces an exponentially distributed random number with rate
// parameter 1.
func (r RNG) ExpFloat64() float64 {
	return genExpo(r)
}

// Perm produces a permutation of the integers in the range [0, n).
func (r RNG) Perm(n int) []int {
	a := make([]int, n)
	for i := range a {
		j := r.Intn(i + 1)
		a[i] = a[j]
		a[j] = i
	}
	return a
}
