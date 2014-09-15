package random

import (
	"encoding/binary"
	"errors"
)

const (
	_N        = 312
	_A uint64 = 0xB5026F5AA96619E9
)

type mt64 struct {
	i int
	s [_N]uint64
}

// NewMT64 creates and seeds a 64-bit Mersenne twister PRNG.
func NewMT64(seed int64) Seeder {
	s := mt64{}
	s.Seed(seed)
	return &s
}

func (m *mt64) Uint64() uint64 {
	if m.i >= _N {
		i := 0
		for i < _N/2 {
			x := m.s[i]&0xffffffff80000000 | m.s[i+1]&0x000000007fffffff
			x = x>>1 ^ _A*(x&1)
			m.s[i] = m.s[i+_N/2] ^ x
			i++
		}
		for i < _N-1 {
			x := m.s[i]&0xffffffff80000000 | m.s[i+1]&0x000000007fffffff
			x = x>>1 ^ _A*(x&1)
			m.s[i] = m.s[i-_N/2] ^ x
			i++
		}
		x := m.s[_N-1]&0xffffffff80000000 | m.s[0]&0x000000007fffffff
		x = x>>1 ^ _A*(x&1)
		m.s[_N-1] = m.s[_N/2-1] ^ x
		m.i = 0
	}

	x := m.s[m.i]
	m.i++
	x ^= x >> 29 & 0x5555555555555555
	x ^= x << 17 & 0x71D67FFFEDA60000
	x ^= x << 37 & 0xFFF7EEE000000000
	x ^= x >> 43
	return x
}

func (m *mt64) Seed(seed int64) {
	m.s[0] = uint64(seed)
	for i := 1; i < _N; i++ {
		m.s[i] = 6364136223846793005*(m.s[i-1]^m.s[i-1]>>62) + uint64(i)
	}
	m.i = _N
}

func (m *mt64) MarshalBinary() ([]byte, error) {
	p := [_N*8 + 2]byte{}
	for i, v := range m.s {
		binary.LittleEndian.PutUint64(p[i<<3:], v)
	}
	// Writing m.i last allows us to stay aligned while writing m.s.
	p[len(p)-2] = byte(m.i >> 8)
	p[len(p)-1] = byte(m.i)
	return p[:], nil
}

func (m *mt64) UnmarshalBinary(data []byte) error {
	if len(data) < _N*8+2 {
		return errors.New("not enough data to unmarshal MT64 state")
	}
	for i := range m.s {
		m.s[i] = binary.LittleEndian.Uint64(data[i<<3:])
	}
	m.i = int(data[len(data)-2])<<8 | int(data[len(data)-1])
	return nil
}
