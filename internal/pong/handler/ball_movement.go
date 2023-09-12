package handler

import (
	"log"

	"github.com/vctaragao/pong/internal/pong/entity"
	"golang.org/x/exp/rand"
)

type BallMovementHandler struct {
	logger *log.Logger
	board  *entity.Board
	ball   *entity.Ball
}

func NewBallMovementHandler(board *entity.Board, ball *entity.Ball, logger *log.Logger) BallMovementHandler {
	return BallMovementHandler{board: board, ball: ball, logger: logger}
}

func (h *BallMovementHandler) HandleBallMovement() {
	h.board.ClearCell(h.ball.H, h.ball.W)

	h.moveBall()

	h.board.SetCell(h.ball.H, h.ball.W, entity.BallRune)
}

func (h *BallMovementHandler) moveBall() {
	h.ball.Move()
	h.logger.Printf("Ball moved to: H: %v, W: %v, Dir: %v\n", h.ball.H, h.ball.W, h.ball.Direction)

	if h.hitPlayer() {
		h.logger.Printf("Hit player on: H: %v, W: %v\n", h.ball.H, h.ball.W)
		h.ball.ChangeDirection(entity.Vertical)

		if h.board.IsEdge(h.ball.H) {
			h.logger.Printf("Hit endge on: H: %v, W: %v\n", h.ball.H, h.ball.W)
			h.ball.ChangeDirection(entity.Horizontal)
		}

		h.moveBall()
	} else if h.board.IsGoal(h.ball.W) {
		h.logger.Printf("Hit Goal on: H: %v, W: %v\n", h.ball.H, h.ball.W)
		h.resetBall()
		h.ball.SetDirection(entity.InitialDirection(rand.Intn(4)))
	} else if h.board.IsEdge(h.ball.H) {
		h.logger.Printf("Hit endge on: H: %v, W: %v\n", h.ball.H, h.ball.W)
		h.ball.ChangeDirection(entity.Horizontal)
	}
}

func (h *BallMovementHandler) hitPlayer() bool {
	return h.board.GetCell(h.ball.H, h.ball.W) == entity.PlayerRune
}

func (h *BallMovementHandler) resetBall() {
	h.ball.SetPosition(h.board.GetHeight()/2, h.board.GetWidth()/2)
}
