package game

import "fmt"

type Reason uint8

const (
	RESIGNATION Reason = iota + 1
	MUTUAL_AGREEMENT
	TIME
	CHECKMATE
	STALEMATE
	INSUFFICIENT_MATERIAL
	FIFTY_MOVE_RULE
	THREEFOLD_REPETITION
	PLY_LIMIT_REACHED
)

func (r Reason) String() string {
	switch r {
	case RESIGNATION:
		return "resignation"
	case MUTUAL_AGREEMENT:
		return "mutual agreement"
	case TIME:
		return "time"
	case CHECKMATE:
		return "checkmate"
	case STALEMATE:
		return "stalemate"
	case INSUFFICIENT_MATERIAL:
		return "insufficient material"
	case FIFTY_MOVE_RULE:
		return "fifty move rule"
	case THREEFOLD_REPETITION:
		return "three-fold repetition"
	case PLY_LIMIT_REACHED:
		return "ply limit reached"
	}

	panic(fmt.Sprintf("Unhandled switch case: %d", r))
}
