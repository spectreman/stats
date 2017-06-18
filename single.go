package stats

import (
	"math"
)

// Single houses single-variable descriptive statistics
type Single struct {
	N   int
	X   float64
	XX  float64
	Min float64
	Max float64

	started bool
}

// Reset ...
func (s *Single) Reset() {
	s.N = 0
	s.X = 0
	s.XX = 0
	s.Min = 0
	s.Max = 0

	s.started = false
}

// Add ...
func (s *Single) Add(x float64) {
	if !s.started {
		s.started = true
		s.Min = x
		s.Max = x
	}

	s.N++
	s.X += x
	s.XX += x * x

	if x < s.Min {
		s.Min = x
	}

	if x > s.Max {
		s.Max = x
	}
}

// Mean ...
func (s *Single) Mean() float64 {
	if s.N == 0 {
		return 0
	}
	return s.X / s.n()
}

// SD ...
func (s *Single) SD() float64 {
	variance := s.Var()
	if variance < 0 {
		return 0
	}
	return math.Sqrt(variance)
}

// Var ...
func (s *Single) Var() float64 {
	if s.N < 2 {
		return 0
	}

	num := s.XX - s.X*s.X/s.n()
	den := s.n() - 1
	return num / den
}

// n for convenience
func (s *Single) n() float64 {
	return float64(s.N)
}

// Div returns the quotient of x and y
// It returns zero if the divisor is zero
func div(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	return x / y
}
