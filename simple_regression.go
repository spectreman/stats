package stats

import (
	"math"
)

// SimpleRegr ...
type SimpleRegr struct {
	N  int
	X  float64
	Y  float64
	XX float64
	YY float64
	XY float64
}

// Reset ...
func (s *SimpleRegr) Reset() {
	s.N = 0
	s.X = 0.0
	s.Y = 0.0
	s.XX = 0.0
	s.YY = 0.0
	s.XY = 0.0
}

// UpdateXY ...
func (s *SimpleRegr) UpdateXY(x, y float64) {
	s.N++
	s.X += x
	s.Y += y
	s.XX += x * x
	s.YY += y * y
	s.XY += x * y
}

// Beta ...
func (s *SimpleRegr) Beta() float64 {
	if s.N < 2 {
		return 0.0
	}

	num := s.XY - s.X*s.Y/float64(s.N)
	den := s.XX - s.X*s.X/float64(s.N)

	return div(num, den)
}

// Alpha ...
func (s *SimpleRegr) Alpha() float64 {
	if s.N < 1 {
		return 0.0
	}

	n := float64(s.N)
	return s.Y/n - s.Beta()*s.X/n
}

// BetaT ...
func (s *SimpleRegr) BetaT() float64 {
	if s.N < 3 {
		return 0.0
	}

	b := s.Beta()
	a := s.Alpha()
	e := 0.0
	n := float64(s.N)

	e += s.YY
	e += a * a * n
	e += b * b * s.XX
	e -= 2 * a * s.Y
	e -= 2 * b * s.XY
	e += 2 * a * b * s.X

	num := e / (n - 2.0)
	den := s.XX - s.X*s.X/n
	sv := math.Sqrt(div(num, den))

	return div(b, sv)
}

// Est ...
func (s *SimpleRegr) Est(x float64) float64 {
	return s.Alpha() + s.Beta()*x
}
