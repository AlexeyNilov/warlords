package things

type Unit struct {
	Speed int
	Strength int
}

func NewUnit(speed int, strength int) *Unit {
	return &Unit{Speed: speed, Strength: strength}
}