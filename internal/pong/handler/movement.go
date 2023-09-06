package handler

import (
	"github.com/vctaragao/pong/internal/pong/entity"
	"log"
)

type MovementHandler struct {
	board  *entity.Board
	logger *log.Logger
}

func NewMovementHandler(board *entity.Board, logger *log.Logger) MovementHandler {
	return MovementHandler{board: board, logger: logger}
}

func (h *MovementHandler) StartMomevementPlayerHandlers(b *entity.Board, players []*entity.Player) []entity.PlayerChannel {
	playersChannels := make([]entity.PlayerChannel, 0, len(players))

	for _, player := range players {
		go h.handlerPlayerMovement(b, player)
		playersChannels = append(playersChannels, player.PlayerChannel)
	}

	return playersChannels
}

func (h *MovementHandler) handlerPlayerMovement(b *entity.Board, p *entity.Player) {
	for {
		switch p.HandleChannel() {
		case entity.MovUp:
			h.MoveUp(p)
		case entity.MovDown:
			h.MoveDown(p)
		default:
		}
	}
}

func (h *MovementHandler) MoveUp(player *entity.Player) {
	if player.H-1 == 0 {
		return
	}

	h.board.ClearCell(player.H, player.W)
	player.Up()
	h.board.SetCell(player.H, player.W, entity.PlayerRune)
}

func (h *MovementHandler) MoveDown(player *entity.Player) {
	if player.H+1 == h.board.GetHeight()-1 {
		return
	}

	h.board.ClearCell(player.H, player.W)
	player.Down()
	h.board.SetCell(player.H, player.W, entity.PlayerRune)
}
