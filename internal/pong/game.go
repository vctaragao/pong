package pong

import (
	"github.com/vctaragao/pong/internal/pong/entity"
	"github.com/vctaragao/pong/internal/pong/handler"
	"golang.org/x/exp/slices"
)

type BoardRenderFunc func(b *entity.Board)
type KeyboardHandlerFunc func(playerChannels []entity.PlayerChannel)

type Game struct {
	entity.Board
	handler.MovementHandler
	handler.BallMovementHandler
	Players         []*entity.Player
	Ball            entity.Ball
	render          BoardRenderFunc
	keyboardHandler KeyboardHandlerFunc
}

func NewGame(render BoardRenderFunc, keyBoardHandler KeyboardHandlerFunc) Game {
	board := entity.NewBoard(Height, Width)
	movementHandler := handler.NewMovementHandler(&board)

	ball := entity.NewBall(Height/2, Width/2)
	ballMovementHandler := handler.NewBallMovementHandler(&board, &ball)

	g := Game{
		Ball:                ball,
		Board:               board,
		render:              render,
		MovementHandler:     movementHandler,
		keyboardHandler:     keyBoardHandler,
		BallMovementHandler: ballMovementHandler,
	}

	player1 := entity.NewPlayer(1, Height/2, 2, 'w', 's')
	player2 := entity.NewPlayer(2, Height/2, Width-3, 'k', 'j')

	g.AddPlayer(&player1)
	g.AddPlayer(&player2)

	return g
}

func (g *Game) AddPlayer(p *entity.Player) {
	if !slices.Contains(g.Players, p) {
		g.Players = append(g.Players, p)
		g.SetCell(p.H, p.W, entity.PlayerRune)
	}
}

func (g *Game) Start() {
	playerChannels := g.StartMomevementPlayerHandlers(&g.Board, g.Players)
	go g.keyboardHandler(playerChannels)
	go g.loop()
}

func (g *Game) loop() {
	g.HandleBallMovement()
	g.render(&g.Board)
}
