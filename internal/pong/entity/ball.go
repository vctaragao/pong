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
		Direction: Right,
	}
}

func (b *Ball) SetPosition(h, w int) {
	b.H = h
	b.W = w
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

func (b *Ball) ChangeDirection(wall Wall) {
	var newDirection Direction

	switch b.Direction {
	case Up:
		newDirection = Down
	case Down:
		newDirection = Up
	case UpRight:
		if wall == Vertical {
			newDirection = UpLeft
		} else {
			newDirection = DownRight
		}
	case UpLeft:
		if wall == Vertical {
			newDirection = UpRight
		} else {
			newDirection = DownLeft
		}
	case DownLeft:
		if wall == Vertical {
			newDirection = DownRight
		} else {
			newDirection = UpLeft
		}
	case DownRight:
		if wall == Vertical {
			newDirection = DownLeft
		} else {
			newDirection = UpRight
		}
	}

	b.SetDirection(newDirection)
}

func (b *Ball) SetDirection(d Direction) {
	b.Direction = d
}
