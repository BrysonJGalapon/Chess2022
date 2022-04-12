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

func (p *Piece) IsValidMovement(srcSquare, dstSquare Square) error {
	if srcSquare == dstSquare {
		return fmt.Errorf("can't move a piece onto itself")
	}

	switch p.color {
	case WHITE:
		switch p.pieceType {
		case KING:
			dst := srcSquare.DistanceSquaredTo(dstSquare)
			if dst <= 2 {
				// king can move exactly 1 square in any direction
				return nil
			}

			if srcSquare == GetSquareFromString("E1") && dst == 4 && srcSquare.GetRow() == dstSquare.GetRow() {
				// if king is on E1 square, king can move left or right exactly 2 squares (simulates castling)
				return nil
			}
		case QUEEN:
			startRow := srcSquare.GetRow()
			startCol := srcSquare.GetCol()

			endRow := dstSquare.GetRow()
			endCol := dstSquare.GetCol()

			if Abs(endRow-startRow) == Abs(endCol-startCol) {
				// queens can move diagonally, any number of squares
				return nil
			}

			if Xor(endRow == startRow, endCol == startCol) {
				// queens can move horizontally or vertically, any number of squares
				return nil
			}
		case BISHOP:
			startRow := srcSquare.GetRow()
			startCol := srcSquare.GetCol()

			endRow := dstSquare.GetRow()
			endCol := dstSquare.GetCol()

			if Abs(endRow-startRow) == Abs(endCol-startCol) {
				// bishops can move diagonally, any number of squares
				return nil
			}
		case KNIGHT:
			dst := srcSquare.DistanceSquaredTo(dstSquare)
			if dst == 5 {
				// knight can move in L-shape
				return nil
			}
		case ROOK:
			startRow := srcSquare.GetRow()
			startCol := srcSquare.GetCol()

			endRow := dstSquare.GetRow()
			endCol := dstSquare.GetCol()

			if Xor(endRow == startRow, endCol == startCol) {
				// bishops can move horizontally or vertically, any number of squares
				return nil
			}
		case PAWN:
			startRank := srcSquare.GetRank()
			startFile := srcSquare.GetFile()

			endRank := dstSquare.GetRank()
			endFile := dstSquare.GetFile()

			if startFile == endFile && endRank == startRank+1 {
				// white pawns can move up 1 square
				return nil
			}

			if startFile == endFile && startRank == 2 && endRank == startRank+2 {
				// if pawn is on 2nd rank, pawn can move up 2 squares
				return nil
			}
		}
	case BLACK:
		switch p.pieceType {
		case KING:
			dst := srcSquare.DistanceSquaredTo(dstSquare)
			if dst <= 2 {
				// king can move exactly 1 square in any direction
				return nil
			}

			if srcSquare == GetSquareFromString("E8") && dst == 4 && srcSquare.GetRow() == dstSquare.GetRow() {
				// if king is on E8 square, king can move left or right exactly 2 squares (simulates castling)
				return nil
			}
		case QUEEN:
			startRow := srcSquare.GetRow()
			startCol := srcSquare.GetCol()

			endRow := dstSquare.GetRow()
			endCol := dstSquare.GetCol()

			if Abs(endRow-startRow) == Abs(endCol-startCol) {
				// queens can move diagonally, any number of squares
				return nil
			}

			if Xor(endRow == startRow, endCol == startCol) {
				// queens can move horizontally or vertically, any number of squares
				return nil
			}
		case BISHOP:
			startRow := srcSquare.GetRow()
			startCol := srcSquare.GetCol()

			endRow := dstSquare.GetRow()
			endCol := dstSquare.GetCol()

			if Abs(endRow-startRow) == Abs(endCol-startCol) {
				// bishops can move diagonally, any number of squares
				return nil
			}
		case KNIGHT:
			dst := srcSquare.DistanceSquaredTo(dstSquare)
			if dst == 5 {
				// knight can move in L-shape
				return nil
			}
		case ROOK:
			startRow := srcSquare.GetRow()
			startCol := srcSquare.GetCol()

			endRow := dstSquare.GetRow()
			endCol := dstSquare.GetCol()

			if Xor(endRow == startRow, endCol == startCol) {
				// bishops can move horizontally or vertically, any number of squares
				return nil
			}
		case PAWN:
			startRank := srcSquare.GetRank()
			startFile := srcSquare.GetFile()

			endRank := dstSquare.GetRank()
			endFile := dstSquare.GetFile()

			if startFile == endFile && endRank == startRank-1 {
				// black pawns can move down 1 square
				return nil
			}

			if startFile == endFile && startRank == 7 && endRank == startRank-2 {
				// if pawn is on 7th rank, pawn can move down 2 squares
				return nil
			}
		}
	}

	return fmt.Errorf("%s can't move from %s to %s", (*p).String(), srcSquare.GetName(), dstSquare.GetName())
}

func (p *Piece) IsValidCapture(srcSquare, dstSquare Square) error {
	switch p.color {
	case WHITE:
		switch p.pieceType {
		case KING:
			dst := srcSquare.DistanceSquaredTo(dstSquare)
			if dst <= 2 {
				// king can capture 1 square in any direction
				return nil
			}
		case QUEEN:
		case BISHOP:
		case KNIGHT:
		case ROOK:
			// valid captures for [queen, bishop, knight, rook] are exact same as movement
			return p.IsValidMovement(srcSquare, dstSquare)
		case PAWN:
			startRank := srcSquare.GetRank()
			startFile := srcSquare.GetFile()

			endRank := dstSquare.GetRank()
			endFile := dstSquare.GetFile()

			if Abs(endFile-startFile) == 1 && endRank == startRank+1 {
				// white pawns can capture upward diagonally
				return nil
			}
		}
	case BLACK:
		switch p.pieceType {
		case KING:
			dst := srcSquare.DistanceSquaredTo(dstSquare)
			if dst <= 2 {
				// king can capture 1 square in any direction
				return nil
			}
		case QUEEN:
		case BISHOP:
		case KNIGHT:
		case ROOK:
			// valid captures for [queen, bishop, knight, rook] are exact same as movement
			return p.IsValidMovement(srcSquare, dstSquare)
		case PAWN:
			startRank := srcSquare.GetRank()
			startFile := srcSquare.GetFile()

			endRank := dstSquare.GetRank()
			endFile := dstSquare.GetFile()

			if Abs(endFile-startFile) == 1 && endRank == startRank-1 {
				// black pawns can capture downward diagonally
				return nil
			}
		}
	}

	return fmt.Errorf("%s can't capture from %s to %s", (*p).String(), srcSquare.GetName(), dstSquare.GetName())
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

func GetPiece(c Color, pt PieceType) *Piece {
	switch c {
	case WHITE:
		switch pt {
		case KING:
			return WHITE_KING
		case QUEEN:
			return WHITE_QUEEN
		case BISHOP:
			return WHITE_BISHOP
		case KNIGHT:
			return WHITE_KNIGHT
		case ROOK:
			return WHITE_ROOK
		case PAWN:
			return WHITE_PAWN
		}
	case BLACK:
		switch pt {
		case KING:
			return BLACK_KING
		case QUEEN:
			return BLACK_QUEEN
		case BISHOP:
			return BLACK_BISHOP
		case KNIGHT:
			return BLACK_KNIGHT
		case ROOK:
			return BLACK_ROOK
		case PAWN:
			return BLACK_PAWN
		}
	}

	panic(fmt.Sprintf("Unhandled switch case: %d, %d", c, pt))
}
