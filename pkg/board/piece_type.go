package board

import "fmt"

type PieceType uint8

const (
	KING PieceType = iota + 1
	QUEEN
	KNIGHT
	BISHOP
	ROOK
	PAWN
)

var PROMOTION_PIECE_TYPES = [4]PieceType{QUEEN, KNIGHT, BISHOP, ROOK}

func (pt PieceType) IsValidPromotionPiece() bool {
	switch pt {
	case QUEEN:
		return true
	case KNIGHT:
		return true
	case BISHOP:
		return true
	case ROOK:
		return true
	case KING:
		return false
	case PAWN:
		return false
	}

	panic(fmt.Sprintf("Unhandled switch case: %s", pt))
}

func (pt PieceType) String() string {
	switch pt {
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

	panic(fmt.Sprintf("Unhandled switch case: %d", pt))
}

func NewPieceTypeFromString(s string) (PieceType, error) {
	switch s {
	case "Q":
		return QUEEN, nil
	case "N":
		return KNIGHT, nil
	case "B":
		return BISHOP, nil
	case "R":
		return ROOK, nil
	case "K":
		return KING, nil
	case "P":
		return PAWN, nil
	}

	return 0, fmt.Errorf("no piece type associated to: %s", s)
}
