package handler

import (
	"github.com/vctaragao/pong/internal/pong/entity"
	"golang.org/x/exp/rand"
)

type BallMovementHandler struct {
	board *entity.Board
	ball  *entity.Ball
}

func NewBallMovementHandler(board *entity.Board, ball *entity.Ball) BallMovementHandler {
	return BallMovementHandler{board: board, ball: ball}
}

func (h *BallMovementHandler) HandleBallMovement() {
	h.board.ClearCell(h.ball.H, h.ball.W)

	h.ball.Move()

	if h.board.IsGoal(h.ball.W) {
		h.resetBall()
		h.ball.SetDirection(entity.InitialDirection(rand.Intn(4)))
	} else if h.board.IsEdge(h.ball.H) {
		h.ball.ChangeDirection()
	}

	h.board.SetCell(h.ball.H, h.ball.W, entity.BallRune)
}

func (h *BallMovementHandler) resetBall() {
	h.ball.SetPosition(h.board.GetHeight()/2, h.board.GetWidth()/2)
}
