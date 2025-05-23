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

func GetWiningProbability(unit1Strength, unit2Strength, minWinChance float64) float64 {
	// Compute total strength
	totalStrength := unit1Strength + unit2Strength

	// Compute raw probabilities based on strength
	unit1Probability := unit1Strength / totalStrength
	unit2Probability := unit2Strength / totalStrength

	// Adjust probabilities to ensure minimum win chance
	if unit1Probability < minWinChance {
		unit1Probability = minWinChance
	}
	if unit2Probability < minWinChance {
		unit1Probability = 1 - minWinChance
	}

	return unit1Probability
}

func (c *Combat) Fight(minWinChance float64) *things.Unit {
	attackerProbability := GetWiningProbability(c.Attacker.Strength, c.Defender.Strength, minWinChance)
	randomNumber := rand.Float64()
	if randomNumber <= attackerProbability {
		return c.Attacker
	}
	return c.Defender
}
