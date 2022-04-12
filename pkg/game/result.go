package game

import "fmt"

type Result uint8

const (
	WHITE_WINS Result = iota
	BLACK_WINS
	GAME_DRAWN
	UNDETERMINED
)

func (r Result) String() string {
	switch r {
	case WHITE_WINS:
		return "white wins"
	case BLACK_WINS:
		return "black wins"
	case GAME_DRAWN:
		return "game drawn"
	}

	panic(fmt.Sprintf("Unhandled switch case: %d", r))
}
