package main

import "fmt"

type Field [Height][Width]rune

type Board struct {
	Field Field
	p1    *Player
	p2    *Player
}

func (b *Board) InitializeBoard() {
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			b.Field[i][j] = ' '
		}
	}

	for i := 0; i < Height; i++ {
		b.Field[i][0] = '|'
		b.Field[i][Width-1] = '|'
	}

	for i := 0; i < Width; i++ {
		b.Field[0][i] = '-'
		b.Field[Height-1][i] = '-'
	}

}

func (b *Board) Print() {
	b.setPlayersOnField()

	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			fmt.Printf("%v", string(b.Field[i][j]))
		}
		fmt.Println()
	}
}

func (b *Board) AddPlayer(p *Player) {
	if b.p1 == nil {
		fmt.Printf("Setting P1: %v\n", p)
		b.p1 = p
		return
	}

	if b.p2 == nil {
		fmt.Printf("Setting P2: %v\n", p)
		b.p2 = p
	}

}

func (b *Board) Up(p *Player) {
	if p.H-1 == 0 {
		return
	}

	b.Field[p.H][p.W] = ' '
	p.Up()
}

func (b *Board) Down(p *Player) {
	if p.H+1 == Height-1 {
		return
	}

	b.Field[p.H][p.W] = ' '
	p.Down()
}

func (b *Board) setPlayersOnField() {
	if b.p1 != nil {
		fmt.Printf("Setting P1: H: %v, W: %v\n", b.p1.H, b.p1.W)
		b.Field[b.p1.H][b.p1.W] = PlayerRune
	}

	if b.p2 != nil {
		fmt.Printf("Setting P2: H: %v, W: %v\n", b.p2.H, b.p2.W)
		b.Field[b.p2.H][b.p2.W] = PlayerRune
	}
}
