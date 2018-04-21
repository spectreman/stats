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

// Percentile returns the largest value less than or equal to the highest
// (1-p) portion of the data..
func (m *Median) Percentile(p float64) float64 {

	if len(m.values) == 0 {
		return 0.0
	}

	n := len(m.values)
	index := int(p * float64(n))

	if index < 0 {
		return m.values[0]
	}

	if index >= n {
		return m.values[n-1]
	}

	return m.values[index]
}
