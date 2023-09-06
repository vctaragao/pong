package entity

const (
	PlayerRune rune = '|'

	Uknown = iota
	MovUp
	MovDown
)

type (
	PlayerChannel chan rune

	Position struct {
		H int
		W int
	}

	Movement struct {
		Up   rune
		Down rune
	}

	Player struct {
		ID int
		Position
		Movement
		PlayerChannel
	}
)

func NewPlayer(id, h, w int, up, down rune) Player {
	return Player{
		Position: Position{
			H: h,
			W: w,
		},
		Movement: Movement{
			Up:   up,
			Down: down,
		},
		PlayerChannel: make(PlayerChannel, 1),
	}

}

func (p *Player) Up() {
	p.H--
}

func (p *Player) Down() {
	p.H++
}

func (p *Player) HandleChannel() int {
	key := <-p.PlayerChannel

	switch key {
	case p.Movement.Up:
		return MovUp
	case p.Movement.Down:
		return MovDown
	default:
		return Uknown
	}
}
