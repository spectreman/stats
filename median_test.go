package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	assert := assert.New(t)

	m := &Median{}
	assert.Equal(m.Med(), 0.0)

	m.Add(1)
	assert.Equal(m.Med(), 1.0)

	m.Add(4)
	assert.Equal(m.Med(), 2.5)

	m.Add(3)
	assert.Equal(m.Med(), 3.0)
}

func TestPercentile(t *testing.T) {
	assert := assert.New(t)

	m := &Median{}
	m.Add(1)
	m.Add(2)
	m.Add(3)
	m.Add(4)
	m.Add(5)

	assert.Equal(m.Percentile(0.0), 1.0)
	assert.Equal(m.Percentile(0.1), 1.0)
	assert.Equal(m.Percentile(0.2), 2.0)
	assert.Equal(m.Percentile(0.3), 2.0)
	assert.Equal(m.Percentile(0.4), 3.0)
	assert.Equal(m.Percentile(0.5), 3.0)
	assert.Equal(m.Percentile(0.6), 4.0)
	assert.Equal(m.Percentile(0.7), 4.0)
	assert.Equal(m.Percentile(0.8), 5.0)
	assert.Equal(m.Percentile(0.9), 5.0)
	assert.Equal(m.Percentile(1.0), 5.0)
}
