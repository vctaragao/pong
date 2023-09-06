package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-tty"
	"github.com/vctaragao/pong/internal/pong"
	"github.com/vctaragao/pong/internal/pong/entity"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	game := pong.NewGame(render, keyboardHandler)
	game.Start()

	<-signals
}

func render(board *entity.Board) {
	for {
		screen.Clear()
		screen.MoveTopLeft()

		board.Print()
		time.Sleep(time.Second / time.Duration(pong.FPS))
	}
}

func keyboardHandler(playerChannels []entity.PlayerChannel) {
	t, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	for {
		key, err := t.ReadRune()
		if err != nil {
			panic(err)
		}

		for _, channel := range playerChannels {
			channel <- key
		}
	}
}
