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
	direction := h.ball.Direction

}
