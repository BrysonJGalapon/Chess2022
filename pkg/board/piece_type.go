package board

type PieceType int64

const (
	KING PieceType = iota
	QUEEN
	KNIGHT
	BISHOP
	ROOK
	PAWN
)
