package random

// A Distribution represents a pdf under which random numbers can be generated.
type Distribution interface {
	Next(Source) float64
}
