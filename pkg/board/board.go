package board

import "fmt"

type Board interface {
	Make(Move) error
	IsValidMove(Move) error
	Copy() Board
	String() string

	makeUnsafe(Move)
	isValidPosition() error
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
	whiteQueenside bool
	blackKingside  bool
	blackQueenSide bool

	// turn
	turn Color
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

func (b *board) setPieceBitmap(color Color, pieceType PieceType, bitmap BitMap) {
	switch color {
	case WHITE:
		switch pieceType {
		case KING:
			b.whiteKingBitMap = bitmap
			return
		case QUEEN:
			b.whiteQueenBitMap = bitmap
			return
		case BISHOP:
			b.whiteBishopBitMap = bitmap
			return
		case KNIGHT:
			b.whiteKnightBitMap = bitmap
			return
		case ROOK:
			b.whiteRookBitMap = bitmap
			return
		case PAWN:
			b.whitePawnBitMap = bitmap
			return
		}
	case BLACK:
		switch pieceType {
		case KING:
			b.blackKingBitMap = bitmap
			return
		case QUEEN:
			b.blackQueenBitMap = bitmap
			return
		case BISHOP:
			b.blackBishopBitMap = bitmap
			return
		case KNIGHT:
			b.blackKnightBitMap = bitmap
			return
		case ROOK:
			b.blackRookBitMap = bitmap
			return
		case PAWN:
			b.blackPawnBitMap = bitmap
			return
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

func (b *board) GetTurn() Color {
	return b.turn
}

func (b *board) Make(m Move) error {
	// Do nothing on empty move
	if m.IsEmpty() {
		return nil
	}

	if err := b.IsValidMove(m); err != nil {
		return err
	}

	b.makeUnsafe(m)
	return nil
}

func (b *board) makeUnsafe(m Move) {
	// Do nothing on empty move
	if m.IsEmpty() {
		return
	}

	var srcSquare, dstSquare Square
	var piece *Piece

	srcSquare = m.GetSrcSquare()
	dstSquare = m.GetDstSquare()

	piece = b.pickUpPieceAt(srcSquare)
	if piece != nil {
		b.placePieceAt(piece, dstSquare)
	}

	// toggle the turn
	b.turn = b.turn.Opposite()
}

func (b *board) pickUpPieceAt(s Square) *Piece {
	piece, err := b.getPieceAt(s)
	if err != nil {
		// no piece on the square, do nothing and return nothing
		return nil
	}

	color := piece.GetColor()
	pieceType := piece.GetPieceType()

	bitmap := b.getPieceBitmap(color, pieceType)
	bitmap = bitmap &^ (s.ToBitMap()) // removes the bit that square represents from the bitmap, if it exists
	b.setPieceBitmap(color, pieceType, bitmap)

	return piece
}

func (b *board) placePieceAt(p *Piece, s Square) *Piece {
	pieceOriginallyAtSqaure, err := b.getPieceAt(s)
	if err == nil {
		// update the bitmap of the removed piece
		pieceOriginallyAtSqaureColor := pieceOriginallyAtSqaure.GetColor()
		pieceOriginallyAtSqaurePieceType := pieceOriginallyAtSqaure.GetPieceType()
		bitmap := b.getPieceBitmap(pieceOriginallyAtSqaureColor, pieceOriginallyAtSqaurePieceType)
		bitmap = bitmap &^ (s.ToBitMap()) // removes the bit that square represents from the bitmap, if it exists
		b.setPieceBitmap(pieceOriginallyAtSqaureColor, pieceOriginallyAtSqaurePieceType, bitmap)
	}

	// update the bitmap of the added piece
	pieceColor := p.GetColor()
	piecePieceType := p.GetPieceType()
	bitmap := b.getPieceBitmap(pieceColor, piecePieceType)
	bitmap = bitmap | (s.ToBitMap()) // adds the bit that square represents to the bitmap
	b.setPieceBitmap(pieceColor, piecePieceType, bitmap)

	if err != nil {
		return nil
	} else {
		return pieceOriginallyAtSqaure
	}
}

func (b *board) IsValidMove(m Move) error {
	bCopy := b.Copy()
	bCopy.makeUnsafe(m)
	return bCopy.isValidPosition()
}

func (b *board) isValidPosition() error {
	// TODO
	return nil
}

func (b *board) Copy() Board {
	// TODO
	return &board{
		whiteKingBitMap:   b.whiteKingBitMap,
		whiteQueenBitMap:  b.whiteQueenBitMap,
		whiteBishopBitMap: b.whiteBishopBitMap,
		whiteKnightBitMap: b.whiteKnightBitMap,
		whiteRookBitMap:   b.whiteRookBitMap,
		whitePawnBitMap:   b.whitePawnBitMap,
		blackKingBitMap:   b.blackKingBitMap,
		blackQueenBitMap:  b.blackQueenBitMap,
		blackBishopBitMap: b.blackBishopBitMap,
		blackKnightBitMap: b.blackKnightBitMap,
		blackRookBitMap:   b.blackRookBitMap,
		blackPawnBitMap:   b.blackPawnBitMap,

		enPassentBitMap: b.enPassentBitMap,

		whiteKingside:  b.whiteKingside,
		whiteQueenside: b.whiteQueenside,
		blackKingside:  b.blackKingside,
		blackQueenSide: b.blackQueenSide,

		turn: b.turn,
	}
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

		enPassentBitMap: 0,

		whiteKingside:  true,
		whiteQueenside: true,
		blackKingside:  true,
		blackQueenSide: true,

		turn: WHITE,
	}
}

func GetBoardFromString(s string) Board {
	// TODO
	return &board{}
}
