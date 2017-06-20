package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleRegr(t *testing.T) {
	assert := assert.New(t)

	s := &SimpleRegr{}
	s.UpdateXY(1.0, 1.0)
	s.UpdateXY(3.0, 4.0)
	s.UpdateXY(8.0, 9.0)
	s.UpdateXY(9.0, 5.0)

	assert.Equal(s.Beta(), 0.6759776536312849)
	assert.Equal(s.BetaT(), 1.8233150434037104)
	assert.Equal(s.Alpha(), 1.2011173184357542)
	assert.Equal(s.Est(5.0), 4.581005586592179)
}
