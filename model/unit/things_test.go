package things

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestUnit(t *testing.T) {
	unit := NewUnit(1, 1)
	assert.Equal(t, 1, unit.Speed)
	assert.Equal(t, 1, unit.Strength)
}
