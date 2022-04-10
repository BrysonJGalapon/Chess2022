package player

import "galapb/chess2022/pkg/board"

type Player interface {
	Start(board board.Board, quit chan bool)
}
