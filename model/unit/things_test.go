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

func TestArmy(t *testing.T) {
	army := NewArmy()
	assert.Equal(t, 0, army.Speed)
	assert.Equal(t, 0, army.Size)
}

func TestAddUnit(t *testing.T) {
	army := NewArmy()
	army.AddUnit(NewUnit(1, 1))
	assert.Equal(t, 1, army.Speed)
	assert.Equal(t, 1, army.Size)
}

// Army speed equals to speed of the slowest unit
func TestArmySpeed(t *testing.T) {
	army := NewArmy()
	army.AddUnit(NewUnit(1, 1))
	army.AddUnit(NewUnit(2, 1))
	assert.Equal(t, 1, army.Speed)
}
