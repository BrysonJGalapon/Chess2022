package game

import (
	"fmt"
	"galapb/chess2022/pkg/board"
	"galapb/chess2022/pkg/players/player"
	"galapb/chess2022/pkg/time_control"
)

type Game interface {
	GetTimeControl() time_control.TimeControl
	GetWhitePlayer() player.Player
	GetBlackPlayer() player.Player
	GetBoard() board.Board
	GetResult() (Result, Reason)
}

type GameBuilder interface {
	Board(board.Board) GameBuilder
	Build() Game
}

type game struct {
	timeControl time_control.TimeControl
	whitePlayer player.Player
	blackPlayer player.Player
	board       board.Board
}

func (g *game) GetTimeControl() time_control.TimeControl {
	return g.timeControl
}

func (g *game) GetWhitePlayer() player.Player {
	return g.whitePlayer
}

func (g *game) GetBlackPlayer() player.Player {
	return g.blackPlayer
}

func (g *game) GetBoard() board.Board {
	return g.board
}

func (g *game) Board(board board.Board) GameBuilder {
	g.board = board
	return g
}

func (g *game) Build() Game {
	return g
}

func (g *game) String() string {
	return fmt.Sprintf("{time control: %s, white player: %s, black player: %s, board: %s}", g.timeControl, g.whitePlayer, g.blackPlayer, g.board)
}

func (g *game) GetResult() (Result, Reason) {
	b := g.GetBoard()

	switch b.GetStatus() {
	case board.CHECKMATE:
		if b.GetTurn() == board.BLACK {
			return WHITE_WINS, CHECKMATE // black is checkmated; white wins
		} else {
			return BLACK_WINS, CHECKMATE // white is checkmated; white wins
		}
	case board.STALEMATE:
		return GAME_DRAWN, STALEMATE
	case board.INSUFFICIENT_MATERIAL:
		return GAME_DRAWN, INSUFFICIENT_MATERIAL
	case board.FIFTY_MOVE_RULE:
		return GAME_DRAWN, FIFTY_MOVE_RULE
	case board.THREEFOLD_REPETITION:
		return GAME_DRAWN, THREEFOLD_REPETITION
	}

	// TODO check loss on time
	// TODO check mutual agreement
	// TODO check resignation

	// no result determined
	return UNDETERMINED, 0
}

func New(tc time_control.TimeControl, whitePlayer player.Player, blackPlayer player.Player) GameBuilder {
	return &game{tc, whitePlayer, blackPlayer, board.Standard()}
}
