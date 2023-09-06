package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-tty"
	"github.com/vctaragao/pong/internal/pong"
	"github.com/vctaragao/pong/internal/pong/entity"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	logFile, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	l := log.New(logFile, "", log.LstdFlags)

	fmt.Fprintf(logFile, "Dest: %v\n", l.Writer())
	l.Printf("test")

	game := pong.NewGame(render, keyboardHandler, l)
	game.Start()

	<-signals
}

func render(board *entity.Board) {
	screen.Clear()
	screen.MoveTopLeft()

	board.Print()
}

func keyboardHandler(playerChannels []entity.PlayerChannel, logger *log.Logger) {
	t, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		logger.Println("Deffer")
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
