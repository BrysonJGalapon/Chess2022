package board

import "fmt"

type Piece struct {
	color     Color
	pieceType PieceType
}

var (
	WHITE_KING   *Piece = &Piece{WHITE, KING}
	WHITE_QUEEN  *Piece = &Piece{WHITE, QUEEN}
	WHITE_BISHOP *Piece = &Piece{WHITE, BISHOP}
	WHITE_KNIGHT *Piece = &Piece{WHITE, KNIGHT}
	WHITE_ROOK   *Piece = &Piece{WHITE, ROOK}
	WHITE_PAWN   *Piece = &Piece{WHITE, PAWN}

	BLACK_KING   *Piece = &Piece{BLACK, KING}
	BLACK_QUEEN  *Piece = &Piece{BLACK, QUEEN}
	BLACK_BISHOP *Piece = &Piece{BLACK, BISHOP}
	BLACK_KNIGHT *Piece = &Piece{BLACK, KNIGHT}
	BLACK_ROOK   *Piece = &Piece{BLACK, ROOK}
	BLACK_PAWN   *Piece = &Piece{BLACK, PAWN}
)

func (p *Piece) GetColor() Color {
	return p.color
}

func (p *Piece) GetPieceType() PieceType {
	return p.pieceType
}

func (p *Piece) String() string {
	switch p.color {
	case WHITE:
		switch p.pieceType {
		case KING:
			return "K"
		case QUEEN:
			return "Q"
		case BISHOP:
			return "B"
		case KNIGHT:
			return "N"
		case ROOK:
			return "R"
		case PAWN:
			return "P"
		}
	case BLACK:
		switch p.pieceType {
		case KING:
			return "k"
		case QUEEN:
			return "q"
		case BISHOP:
			return "b"
		case KNIGHT:
			return "n"
		case ROOK:
			return "r"
		case PAWN:
			return "p"
		}
	}

	panic(fmt.Sprintf("Unhandled switch case: %d, %d", p.color, p.pieceType))
}
