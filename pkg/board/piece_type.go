package board

import "fmt"

type PieceType uint8

const (
	KING PieceType = iota
	QUEEN
	KNIGHT
	BISHOP
	ROOK
	PAWN
)

func (pt *PieceType) String() string {
	switch *pt {
	case KING:
		return "KING"
	case QUEEN:
		return "QUEEN"
	case KNIGHT:
		return "KNIGHT"
	case BISHOP:
		return "BISHOP"
	case ROOK:
		return "ROOK"
	case PAWN:
		return "PAWN"
	}

	panic(fmt.Sprintf("Unhandled switch case: %d", *pt))
}
