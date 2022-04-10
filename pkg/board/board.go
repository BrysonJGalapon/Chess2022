package board

import "fmt"

type Board interface {
	Make(Move) error
	Copy() Board
	String() string
}

type board struct {
	// piece bitmaps
	whiteKingBitMap   BitMap
	whiteQueenBitMap  BitMap
	whiteBishopBitMap BitMap
	whiteKnightBitMap BitMap
	whiteRookBitMap   BitMap
	whitePawnBitMap   BitMap
	blackKingBitMap   BitMap
	blackQueenBitMap  BitMap
	blackBishopBitMap BitMap
	blackKnightBitMap BitMap
	blackRookBitMap   BitMap
	blackPawnBitMap   BitMap

	// en passent square
	enPassentBitMap BitMap

	// castling rights
	whiteKingside  bool
	whieQueenside  bool
	blackKingside  bool
	blackQueenSide bool
}

func (b *board) getPieceBitmap(color Color, pieceType PieceType) BitMap {
	switch color {
	case WHITE:
		switch pieceType {
		case KING:
			return b.whiteKingBitMap
		case QUEEN:
			return b.whiteQueenBitMap
		case BISHOP:
			return b.whiteBishopBitMap
		case KNIGHT:
			return b.whiteKnightBitMap
		case ROOK:
			return b.whiteRookBitMap
		case PAWN:
			return b.whitePawnBitMap
		}
	case BLACK:
		switch pieceType {
		case KING:
			return b.blackKingBitMap
		case QUEEN:
			return b.blackQueenBitMap
		case BISHOP:
			return b.blackBishopBitMap
		case KNIGHT:
			return b.blackKnightBitMap
		case ROOK:
			return b.blackRookBitMap
		case PAWN:
			return b.blackPawnBitMap
		}
	}

	panic(fmt.Sprintf("Unhandled switch case: %s, %s", color.String(), pieceType.String()))
}

func (b *board) getPieceAt(square Square) (*Piece, error) {
	var bitmap BitMap = square.ToBitMap()

	switch {
	case b.whiteKingBitMap&bitmap != 0:
		return WHITE_KING, nil
	case b.whiteQueenBitMap&bitmap != 0:
		return WHITE_QUEEN, nil
	case b.whiteBishopBitMap&bitmap != 0:
		return WHITE_BISHOP, nil
	case b.whiteKnightBitMap&bitmap != 0:
		return WHITE_KNIGHT, nil
	case b.whiteRookBitMap&bitmap != 0:
		return WHITE_ROOK, nil
	case b.whitePawnBitMap&bitmap != 0:
		return WHITE_PAWN, nil
	case b.blackKingBitMap&bitmap != 0:
		return BLACK_KING, nil
	case b.blackQueenBitMap&bitmap != 0:
		return BLACK_QUEEN, nil
	case b.blackBishopBitMap&bitmap != 0:
		return BLACK_BISHOP, nil
	case b.blackKnightBitMap&bitmap != 0:
		return BLACK_KNIGHT, nil
	case b.blackRookBitMap&bitmap != 0:
		return BLACK_ROOK, nil
	case b.blackPawnBitMap&bitmap != 0:
		return BLACK_PAWN, nil
	}

	return nil, fmt.Errorf("no piece at square: %d", square)
}

func (b *board) Make(m Move) error {
	// TODO
	return nil
}

func (b *board) Copy() Board {
	// TODO
	return &board{}
}

func (b *board) String() string {
	var square Square = 1
	var piece *Piece
	var err error

	ret := ""
	for i := 0; i < 64; i++ {
		if piece, err = b.getPieceAt(square); err != nil {
			ret += "-"
		} else {
			ret += piece.String()
		}

		if (i%8 == 7) && (i != 63) {
			ret += "\n"
		}

		square <<= 1
	}

	return ret
}

func Standard() Board {
	return &board{
		whiteKingBitMap:   1152921504606846976,
		whiteQueenBitMap:  576460752303423488,
		whiteBishopBitMap: 2594073385365405696,
		whiteKnightBitMap: 4755801206503243776,
		whiteRookBitMap:   9295429630892703744,
		whitePawnBitMap:   71776119061217280,
		blackKingBitMap:   16,
		blackQueenBitMap:  8,
		blackBishopBitMap: 36,
		blackKnightBitMap: 66,
		blackRookBitMap:   129,
		blackPawnBitMap:   65280,

		// en passent square
		enPassentBitMap: 0,

		// castling rights
		whiteKingside:  true,
		whieQueenside:  true,
		blackKingside:  true,
		blackQueenSide: true,
	}
}

func GetBoardFromString(s string) Board {
	// TODO
	return &board{}
}
