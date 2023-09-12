package entity

type Wall int

const (
	Vertical Wall = iota
	Horizontal
)

func (w Wall) String() string {
	return [...]string{"Vertical", "Horizontal"}[w]
}
