package things

type Unit struct {
	Speed float64
	Strength float64
}

func NewUnit(speed float64, strength float64) *Unit {
	return &Unit{Speed: speed, Strength: strength}
}

type Army struct {
	Units []*Unit
	Speed float64
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
