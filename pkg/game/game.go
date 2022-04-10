package game

import "fmt"
import "galapb/chess2022/pkg/players/player"
import "galapb/chess2022/pkg/time_control"
import "galapb/chess2022/pkg/board"

type Game interface {
	GetTimeControl() time_control.TimeControl
	GetWhitePlayer() player.Player
	GetBlackPlayer() player.Player
	GetBoard() board.Board
	IsOver() bool
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

func (g *game) IsOver() bool {
	// TODO
	return false
}

func (g *game) Build() Game {
	return g
}

func (g *game) String() string {
	return fmt.Sprintf("{time control: %s, white player: %s, black player: %s, board: %s}", g.timeControl, g.whitePlayer, g.blackPlayer, g.board)
}

func New(tc time_control.TimeControl, whitePlayer player.Player, blackPlayer player.Player) GameBuilder {
	return &game{tc, whitePlayer, blackPlayer, board.Standard()}
}
