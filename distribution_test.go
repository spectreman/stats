package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistribution(t *testing.T) {
	assert := assert.New(t)

	d := &Distribution{}
	d.Add(1)
	d.Add(1)
	d.Add(3)
	d.Add(5)

	assert.Equal(d.Percentile(0.76), 5)
	return

	// Out-Of-Bounds
	assert.Equal(d.Percentile(-0.1), 1)
	assert.Equal(d.Percentile(2.00), 5)

	assert.Equal(d.Percentile(0.25), 1)
	assert.Equal(d.Percentile(0.50), 1)
	assert.Equal(d.Percentile(0.51), 3)
	assert.Equal(d.Percentile(0.75), 3)
	assert.Equal(d.Percentile(0.76), 5)
	assert.Equal(d.Percentile(1.00), 5)
}
