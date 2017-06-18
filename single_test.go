package stats

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingle(t *testing.T) {
	assert := assert.New(t)

	s := &Single{}
	assert.Equal(s.N, 0)
	assert.Equal(s.Mean(), 0.0)
	assert.Equal(s.SD(), 0.0)
	assert.Equal(s.Var(), 0.0)

	s.Add(1)
	assert.Equal(s.N, 1)
	assert.Equal(s.Mean(), 1.0)
	assert.Equal(s.SD(), 0.0)
	assert.Equal(s.Var(), 0.0)

	s.Add(5)
	assert.Equal(s.N, 2)
	assert.Equal(s.Mean(), 3.0)
	assert.Equal(s.SD(), math.Sqrt(8.0))
	assert.Equal(s.Var(), 8.0)

	s.Add(6)
	assert.Equal(s.N, 3)
	assert.Equal(s.Mean(), 4.0)
	assert.Equal(s.SD(), math.Sqrt(7.0))
	assert.Equal(s.Var(), 7.0)
}
