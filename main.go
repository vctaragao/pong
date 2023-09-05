package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	board := Board{}
	board.InitializeBoard()

	player1 := NewPlayer(1, Height/2, 2)
	fmt.Printf("Player 1: %v\n", player1)

	player2 := NewPlayer(2, Height/2, Width-3)
	fmt.Printf("Player 2: %v\n", player2)

	board.AddPlayer(&player1)
	board.AddPlayer(&player2)

	keysEvents, err := keyboard.GetKeys(0)
	chk(err)
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")

	go movementHandler(keysEvents, &board, &player1, 'w', 's')
	go movementHandler(keysEvents, &board, &player2, 'k', 'j')
	go render(&board)

	<-signals
}

func render(board *Board) {
	for {
		screen.Clear()
		screen.MoveTopLeft()

		board.Print()
		time.Sleep(time.Second / time.Duration(FPS))
	}
}

func movementHandler(keysEvent <-chan keyboard.KeyEvent, b *Board, p *Player, up, down rune) {
	log, err := os.Create("log.txt")
	chk(err)
	defer log.Close()

	for {
		var event keyboard.KeyEvent
		select {
		case event = <-keysEvent:
			chk(event.Err)
		case <-time.After(time.Second * 30):
			fmt.Println("Bye Bye due to inactivity")
			os.Exit(0)
		}

		fmt.Fprintf(log, "Key received: %v\n", event)

		if event.Key == keyboard.KeyEsc || event.Key == keyboard.KeyCtrlC {
			os.Exit(0)
		}

		if event.Rune == up {
			b.Up(p)
		} else if event.Rune == down {
			b.Down(p)
		}
	}
}
