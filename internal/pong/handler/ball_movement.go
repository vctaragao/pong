package handler

import "github.com/vctaragao/pong/internal/pong/entity"

type BallMovementHandler struct {
	board *entity.Board
	ball  *entity.Ball
}

func NewBallMovementHandler(board *entity.Board, ball *entity.Ball) BallMovementHandler {
	return BallMovementHandler{board: board, ball: ball}
}

func (h *BallMovementHandler) HandleBallMovement() {
	// check colision
	h.ball.Move()
	h.board.SetCell(h.ball.H, h.ball.W, entity.BallRune)
}
