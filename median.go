package stats

// Median ...
type Median struct {
	values []float64
}

// Add ...
func (m *Median) Add(d float64) {

	// Find Position To Insert (Maintain Increasing Sorted Order)
	for i, x := range m.values {
		if d < x {
			m.values = append(m.values, 0)
			copy(m.values[i+1:], m.values[i:])
			m.values[i] = d
			return
		}
	}

	// Insert On End
	m.values = append(m.values, d)
}

// Med returns the current median
func (m *Median) Med() float64 {

	if len(m.values) == 0 {
		return 0.0
	}

	n := len(m.values)

	// Odd Size
	if n%2 == 1 {
		return m.values[n/2]
	}

	// Even Size
	return 0.5 * (m.values[(n/2)-1] + m.values[(n/2)])
}
