package combat

import (
	"testing"

	"github.com/AlexeyNilov/warlords/model/things"
	"github.com/stretchr/testify/assert"
)

func TestNewCombat(t *testing.T) {
	c := NewCombat(things.NewUnit(1, 1), things.NewUnit(1, 2))
	assert.Equal(t, 1.0, c.Attacker.Strength)
	assert.Equal(t, 2.0, c.Defender.Strength)
}

func TestGetWiningProbability(t *testing.T) {
	assert.InDelta(t, 0.5, GetWiningProbability(1, 1, 0), 0.01)
	assert.InDelta(t, 0.33, GetWiningProbability(1, 2, 0), 0.01)
	assert.InDelta(t, 0.5, GetWiningProbability(1, 10, 0.5), 0.01)
}

func TestFight(t *testing.T) {
	attacker := things.NewUnit(1, 2)
	defender := things.NewUnit(1, 1)
	c := NewCombat(attacker, defender)

	attackerWins := 0
	for range 10000 {
		winner := c.Fight(0)
		if winner == attacker {
			attackerWins++
		}
	}
	assert.InDelta(t, 0.67, float64(attackerWins)/10000, 0.01)
}
