package main

type Position struct {
	H int
	W int
}

type Player struct {
	ID int
	Position
}

func NewPlayer(id, h, w int) Player {
	return Player{
		Position: Position{
			H: h,
			W: w,
		},
	}
}

func (p *Player) Up() {
	p.H--
}

func (p *Player) Down() {
	p.H++
}
