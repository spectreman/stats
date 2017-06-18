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
