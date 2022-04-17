package player

import "galapb/chess2022/pkg/board"

type Player interface {
	Init(prompt chan board.Move, response chan board.Move)
	Start(board board.Board, quit chan bool)
}
