package random

import "encoding"

// A Source is a source of uniformly-distributed values in the interval
// [0, 1<<64).
type Source interface {
	Uint64() uint64
}

// A Seeder is a source with a state that can be seeded, marshaled, and
// unmarshaled.
type Seeder interface {
	Source
	Seed(seed int64)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
