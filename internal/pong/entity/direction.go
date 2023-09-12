package entity

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right

	UpLeft
	UpRight

	DownLeft
	DownRight
)

func InitialDirection(i int) Direction {
	return []Direction{UpLeft, UpRight, DownRight, DownLeft}[i]
}

func (d Direction) String() string {
	return [...]string{
		"Up",
		"Down",
		"Left",
		"Right",
		"UpLeft",
		"UpRight",
		"DownLeft",
		"DownRight",
	}[d]
}
