package terminal

import "fmt"

type position struct {
	X, Y int
}

type view struct {
	Height int
	Width  int
	position
}

func NewView(h, w, x, y int) view {
	return view{
		Height: h,
		Width:  w,
		position: position{
			X: x,
			Y: y,
		},
	}
}

func (v *view) WriteAt(h, w int, c rune) {
	v.MoveCursor(h, w)
	fmt.Print(c)
}

func (v *view) MoveCursor(h, w int) {
	fmt.Printf("\033[%d;%dH", v.Y+h, v.X+w)
}

func (v *view) MoveTopLeft() {
	fmt.Printf("\033[%d;%dH", v.Y, v.X)
}

func (v *view) Clear() {
	fmt.Print("\033[2J")
}
