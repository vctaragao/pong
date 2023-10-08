package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mattn/go-tty"
	"github.com/vctaragao/pong/internal/pong"
	"github.com/vctaragao/pong/internal/pong/entity"
	"github.com/vctaragao/pong/pkg/terminal"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	game := pong.NewGame(render, keyboardHandler, initLog())
	game.Start()

	<-signals
}

func initLog() *log.Logger {
	logFile, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	return log.New(logFile, "", log.LstdFlags)
}

func render(board *entity.Board) {
	terminal.ClearFull()
	terminal.MoveTopLeft()

	board.Print()
}

func keyboardHandler(playerChannels []entity.PlayerChannel, logger *log.Logger) {
	t, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		t.Close()
	}()

	for {
		key, err := t.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		for _, channel := range playerChannels {
			channel <- key
		}
	}
}
