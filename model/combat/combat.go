package combat

import (
	"math/rand/v2"

	"github.com/AlexeyNilov/warlords/model/things"
)

type Combat struct {
	Attacker *things.Unit
	Defender *things.Unit
}

func NewCombat(attacker *things.Unit, defender *things.Unit) *Combat {
	return &Combat{
		Attacker: attacker,
		Defender: defender,
	}
}

func GetBonus(randomness float64) float64 {
	return rand.Float64() * randomness
}

func (c *Combat) Fight(randomness float64) *things.Unit {
	attackerStrength := float64(c.Attacker.Strength) + GetBonus(randomness)
	defenderStrength := float64(c.Defender.Strength) + GetBonus(randomness)
	if attackerStrength > defenderStrength {
		return c.Attacker
	} else {
		return c.Defender
	}
}
