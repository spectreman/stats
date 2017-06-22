package stats

import (
	"math/rand"
)

const epsilon = 1e-10

type obs struct {
	value int
	count int
}

// Distribution ...
type Distribution struct {
	data []*obs
	n    int
	r    *rand.Rand
}

// Seed ...
func (d *Distribution) Seed(seed int64) {
	d.r = rand.New(rand.NewSource(seed))
}

// Add ...
func (d *Distribution) Add(observation int) {

	// Increment Total-Count
	d.n++

	// Find Position To Insert
	for i, x := range d.data {

		// Value Exists: Increment Count
		if x.value == observation {
			x.count++
			return
		}

		// New Value: Insert
		if x.value > observation {
			d.data = append(d.data, &obs{})
			copy(d.data[i+1:], d.data[i:])
			d.data[i] = &obs{value: observation, count: 1}
			return
		}
	}

	// New Max
	d.data = append(d.data, &obs{value: observation, count: 1})
}

// Percentile ...
func (d *Distribution) Percentile(p float64) int {
	switch {
	case p > 1.0:
		p = 1.0
	case p < 0.0:
		p = 0.0
	}

	// Compute Implied-Rank
	rank := p * float64(d.n)

	// Get Value Corresponding To Rank
	return d.rank(rank)
}

// Sample returns a randomly sampled value from the distribution
func (d *Distribution) Sample() int {

	// Ensure Random Source Is Set
	if d.r == nil {
		d.Seed(0)
	}

	// Select A Random Percentile
	p := d.r.Float64()
	return d.Percentile(p)
}

// Sample ...
func (d *Distribution) rank(rank float64) int {

	// Empty
	if d.n == 0 {
		return 0
	}

	// Find Observation
	sum := 0
	for _, x := range d.data {
		sum += x.count
		if float64(sum) >= rank+1+epsilon {
			return x.value
		}
	}

	// At Max
	return d.data[len(d.data)-1].value
}
