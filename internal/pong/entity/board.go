package entity

import "fmt"

type Field [][]rune

type Board struct {
	Field  Field
	height int
	width  int
}

func NewBoard(height, width int) Board {
	field := make([][]rune, height)
	for i := 0; i < height; i++ {
		field[i] = make([]rune, width)
	}

	b := Board{
		height: height,
		width:  width,
		Field:  field,
	}

	b.InitializeBoard()

	return b
}

func (b *Board) InitializeBoard() {
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			b.Field[i][j] = ' '
		}
	}

	for i := 0; i < b.height; i++ {
		b.Field[i][0] = '|'
		b.Field[i][b.width-1] = '|'
	}

	for i := 0; i < b.width; i++ {
		b.Field[0][i] = '-'
		b.Field[b.height-1][i] = '-'
	}

}

func (b *Board) Print() {
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			fmt.Printf("%v", string(b.Field[i][j]))
		}
		fmt.Println()
	}
}

func (b *Board) ClearCell(h, w int) {
	b.Field[h][w] = ' '
}

func (b *Board) SetCell(h, w int, cell rune) {
	b.Field[h][w] = cell
}

func (b *Board) GetHeight() int {
	return b.height
}

func (b *Board) GetWidth() int {
	return b.width
}
