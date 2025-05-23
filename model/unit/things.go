package things

type Unit struct {
	Speed int
	Strength int
}

func NewUnit(speed int, strength int) *Unit {
	return &Unit{Speed: speed, Strength: strength}
}

type Army struct {
	Units []*Unit
	Speed int
	Size int
}

func NewArmy() *Army {
	return &Army{Speed: 0, Size: 0}
}

func (a *Army) AddUnit(unit *Unit) {
	a.Units = append(a.Units, unit)
	if a.Speed == 0 || unit.Speed < a.Speed {
		a.Speed = unit.Speed
	}
	a.Size++
}
