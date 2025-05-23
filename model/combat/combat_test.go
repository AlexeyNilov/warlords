package combat

import (
	"fmt"
	"testing"

	"github.com/AlexeyNilov/warlords/model/things"
	"github.com/stretchr/testify/assert"
)

func TestNewCombat(t *testing.T) {
	c := NewCombat(things.NewUnit(1, 1), things.NewUnit(1, 2))
	assert.Equal(t, 1.0, c.Attacker.Strength)
	assert.Equal(t, 2.0, c.Defender.Strength)
}

func TestGetBonus(t *testing.T) {
	assert.Less(t, GetBonus(0), 1.0)
}

func TestFight(t *testing.T) {
	attacker := things.NewUnit(1, 2)
	defender := things.NewUnit(1, 1)
	c := NewCombat(attacker, defender)
	winner := c.Fight(0)
	assert.Equal(t, attacker, winner)

	attackerWins := 0
	for range 10000 {
		winner := c.Fight(10000)
		if winner == attacker {
			attackerWins++
		}
	}
	fmt.Println(attackerWins)
	assert.Greater(t, 5100, attackerWins)
	assert.Less(t, 4900, attackerWins)
}

