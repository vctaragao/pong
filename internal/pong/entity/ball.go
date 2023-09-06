package entity

const BallRune rune = '*'

type Ball struct {
	Direction
	Position
}

func NewBall(h, w int) Ball {
	return Ball{
		Position: Position{
			H: h,
			W: w,
		},
	}
}

func (b *Ball) Move() {
	switch b.Direction {
	case Up:
		b.MoveUp()
	case Down:
		b.MoveDown()
	case Left:
		b.MoveLeft()
	case Right:
		b.MoveRight()
	case UpRight:
		b.MoveUpRight()
	case UpLeft:
		b.MoveUpLeft()
	case DownLeft:
		b.MoveDownLeft()
	case DownRight:
		b.MoveDownRight()
	}
}

func (b *Ball) MoveUp() {
	b.H--
}

func (b *Ball) MoveUpLeft() {
	b.H--
	b.W--
}
func (b *Ball) MoveUpRight() {
	b.H--
	b.W++
}

func (b *Ball) MoveDown() {
	b.H++
}

func (b *Ball) MoveDownLeft() {
	b.H++
	b.W--
}
func (b *Ball) MoveDownRight() {
	b.H++
	b.W++
}

func (b *Ball) MoveLeft() {
	b.W--
}

func (b *Ball) MoveRight() {
	b.W++
}
