package board

import (
	"fmt"
)

type Board interface {
	Make(Move) error
	IsValidMove(Move) error
	Copy() Board
	String() string
	GetPieceAt(Square) (*Piece, error)
	GetPly() int
	GetStatus() Status
	GetTurn() Color
	GetNumOf(c Color, pt PieceType) int
	GetPieceBitmap(c Color, pt PieceType) BitMap
	IsCheckmate() bool
	IsStalemate() bool

	makeUnsafe(Move)
	toggleTurn()
	setTurn(Color)
	isCheck() bool
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

	// ply
	ply int
}

func (b *board) GetPieceBitmap(color Color, pieceType PieceType) BitMap {
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

func (b *board) GetPly() int {
	return b.ply
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

func (b *board) GetPieceAt(square Square) (*Piece, error) {
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

	switch {
	case piece.GetPieceType() == PAWN && (dstSquare.GetRank() == 1 || dstSquare.GetRank() == 8):
		// promotion moves
		promotionPiece := GetPiece(piece.GetColor(), *m.GetPromotionPieceType())
		b.placePieceAt(promotionPiece, dstSquare)
	case piece.GetPieceType() == PAWN && dstSquare.ToBitMap() == b.enPassentBitMap:
		// en-passent captures
		b.placePieceAt(piece, dstSquare)

		var capturedPawnSquare Square
		if b.GetTurn() == WHITE {
			capturedPawnSquare, _ = dstSquare.Step(SOUTH) // white captured enpassent, so the captured piece is below the destination square
		} else {
			capturedPawnSquare, _ = dstSquare.Step(NORTH) // black captured enpassent, so the captured piece is above the destination square
		}

		b.pickUpPieceAt(capturedPawnSquare)
	case piece.GetPieceType() == KING && srcSquare.DistanceSquaredTo(dstSquare) == 4:
		// castling moves
		directionToRook := srcSquare.DirectionTo(dstSquare)
		rookEndSquare, _ := srcSquare.Step(directionToRook) // rook ends up on square between srcSquare and dstSquare
		_, rookStartSquare := b.getClosestPieceInDirection(srcSquare, directionToRook)

		rook := b.pickUpPieceAt(rookStartSquare)
		b.placePieceAt(rook, rookEndSquare)
		b.placePieceAt(piece, dstSquare)
	default:
		// normal moves
		b.placePieceAt(piece, dstSquare)
	}

	b.updateCastlingRights(piece, srcSquare)
	b.updateEnPassentBitmap(m)
	b.toggleTurn()
	b.incrementPly()
}

func (b *board) updateCastlingRights(p *Piece, s Square) {
	switch p.GetPieceType() {
	case KING:
		// moving the king revokes castling rights
		if p.GetColor() == WHITE {
			b.whiteKingside = false
			b.whiteQueenside = false
		} else {
			b.blackKingside = false
			b.blackQueenSide = false
		}
	case ROOK:
		switch s {
		case H1_SQUARE:
			b.whiteKingside = false // a rook moved from H1, revoke white kingside castling rights
		case A1_SQUARE:
			b.whiteQueenside = false // a rook moved from A1, revoke white queenside castling rights
		case H8_SQUARE:
			b.blackKingside = false // a rook moved from H8, revoke black kingside castling rights
		case A8_SQUARE:
			b.blackQueenSide = false // a rook moved from A8, revoke black queenside castling rights
		}
	default:
		// castling rights are not affected by non-king or non-rook moves
		return
	}
}

func (b *board) updateEnPassentBitmap(m Move) {
	var srcSquare, dstSquare Square
	var enPassentSquare Square
	var piece *Piece

	srcSquare = m.GetSrcSquare()
	dstSquare = m.GetDstSquare()
	piece, _ = b.GetPieceAt(dstSquare)

	if piece.GetPieceType() != PAWN {
		b.enPassentBitMap = 0 // pawn must be moved to set en-passent bitmap
		return
	}

	if srcSquare.DistanceSquaredTo(dstSquare) != 4 {
		b.enPassentBitMap = 0 // pawn must be moved 2 squares up or down to set en-passent bitmap
		return
	}

	enPassentSquare, _ = srcSquare.Step(srcSquare.DirectionTo(dstSquare))
	b.enPassentBitMap = enPassentSquare.ToBitMap()
}

func (b *board) toggleTurn() {
	b.turn = b.turn.Opposite()
}

func (b *board) incrementPly() {
	b.ply += 1
}

func (b *board) setTurn(color Color) {
	b.turn = color
}

func (b *board) isAnyMoveValid() bool {
	var srcSquare Square = 1
	var dstSquare Square = 1
	var move Move

	for i := 0; i < 64; i++ {
		dstSquare = 1

		for j := 0; j < 64; j++ {
			move = NewMove(srcSquare, dstSquare).PromotionPieceType(QUEEN).Build()

			if err := b.IsValidMove(move); err == nil {
				return true
			}

			dstSquare <<= 1
		}

		srcSquare <<= 1
	}

	return false
}

func (b *board) IsCheckmate() bool {
	return b.isCheck() && !b.isAnyMoveValid()
}

func (b *board) IsStalemate() bool {
	return !b.isCheck() && !b.isAnyMoveValid()
}

func (b *board) GetNumOf(c Color, pt PieceType) int {
	switch c {
	case WHITE:
		switch pt {
		case KING:
			return NumSetBits(b.whiteKingBitMap)
		case QUEEN:
			return NumSetBits(b.whiteQueenBitMap)
		case BISHOP:
			return NumSetBits(b.whiteBishopBitMap)
		case KNIGHT:
			return NumSetBits(b.whiteKnightBitMap)
		case ROOK:
			return NumSetBits(b.whiteRookBitMap)
		case PAWN:
			return NumSetBits(b.whitePawnBitMap)
		}
	case BLACK:
		switch pt {
		case KING:
			return NumSetBits(b.blackKingBitMap)
		case QUEEN:
			return NumSetBits(b.blackQueenBitMap)
		case BISHOP:
			return NumSetBits(b.blackBishopBitMap)
		case KNIGHT:
			return NumSetBits(b.blackKnightBitMap)
		case ROOK:
			return NumSetBits(b.blackRookBitMap)
		case PAWN:
			return NumSetBits(b.blackPawnBitMap)
		}
	}

	panic(fmt.Sprintf("Unhandled switch case: %s, %s", c, pt))
}

func (b *board) isInsufficientMaterial() bool {
	numWhiteQueens := NumSetBits(b.whiteQueenBitMap)
	numWhiteBishops := NumSetBits(b.whiteBishopBitMap)
	numWhiteKnights := NumSetBits(b.whiteKnightBitMap)
	numWhiteRooks := NumSetBits(b.whiteRookBitMap)
	numWhitePawns := NumSetBits(b.whitePawnBitMap)

	numBlackQueens := NumSetBits(b.blackQueenBitMap)
	numBlackBishops := NumSetBits(b.blackBishopBitMap)
	numBlackKnights := NumSetBits(b.blackKnightBitMap)
	numBlackRooks := NumSetBits(b.blackRookBitMap)
	numBlackPawns := NumSetBits(b.blackPawnBitMap)

	if numWhiteQueens > 0 || numWhiteRooks > 0 || numWhitePawns > 0 || numBlackQueens > 0 || numBlackRooks > 0 || numBlackPawns > 0 {
		// not insufficient material if there is at least one queen, rook, or pawn
		return false
	}

	// king v king
	isKingVersusKing := numWhiteBishops == 0 && numWhiteKnights == 0 && numBlackBishops == 0 && numBlackKnights == 0
	if isKingVersusKing {
		return true
	}

	// king v king+bishop
	isKingVersusKingAndBishop := numWhiteKnights == 0 && numBlackKnights == 0 && ((numWhiteBishops == 1 && numBlackBishops == 0) || (numWhiteBishops == 0 && numBlackBishops == 1))
	if isKingVersusKingAndBishop {
		return true
	}

	// king v king+knight
	isKingVersusKingAndKnight := numWhiteBishops == 0 && numBlackBishops == 0 && ((numWhiteKnights == 1 && numBlackKnights == 0) || (numWhiteKnights == 0 && numBlackKnights == 1))
	if isKingVersusKingAndKnight {
		return true
	}

	// TODO add king+bishop(s) v king+bishop(s) where all bishops are on same color complex

	return false
}

func (b *board) isFiftyMoveRule() bool {
	// TODO fifty movee rule
	return false
}

func (b *board) isThreefoldRepetition() bool {
	// TODO threefold repetition
	return false
}

func (b *board) GetStatus() Status {
	switch {
	case b.isInsufficientMaterial():
		return INSUFFICIENT_MATERIAL
	case b.IsCheckmate():
		return CHECKMATE
	case b.IsStalemate():
		return STALEMATE
	case b.isFiftyMoveRule():
		return FIFTY_MOVE_RULE
	case b.isThreefoldRepetition():
		return THREEFOLD_REPETITION
	default:
		return UNDETERMINED
	}
}

func (b *board) pickUpPieceAt(s Square) *Piece {
	piece, err := b.GetPieceAt(s)
	if err != nil {
		// no piece on the square, do nothing and return nothing
		return nil
	}

	color := piece.GetColor()
	pieceType := piece.GetPieceType()

	bitmap := b.GetPieceBitmap(color, pieceType)
	bitmap = bitmap &^ (s.ToBitMap()) // removes the bit that square represents from the bitmap, if it exists
	b.setPieceBitmap(color, pieceType, bitmap)

	return piece
}

func (b *board) placePieceAt(p *Piece, s Square) *Piece {
	pieceOriginallyAtSqaure, err := b.GetPieceAt(s)
	if err == nil {
		// update the bitmap of the removed piece
		pieceOriginallyAtSqaureColor := pieceOriginallyAtSqaure.GetColor()
		pieceOriginallyAtSqaurePieceType := pieceOriginallyAtSqaure.GetPieceType()
		bitmap := b.GetPieceBitmap(pieceOriginallyAtSqaureColor, pieceOriginallyAtSqaurePieceType)
		bitmap = bitmap &^ (s.ToBitMap()) // removes the bit that square represents from the bitmap, if it exists
		b.setPieceBitmap(pieceOriginallyAtSqaureColor, pieceOriginallyAtSqaurePieceType, bitmap)
	}

	// update the bitmap of the added piece
	pieceColor := p.GetColor()
	piecePieceType := p.GetPieceType()
	bitmap := b.GetPieceBitmap(pieceColor, piecePieceType)
	bitmap = bitmap | (s.ToBitMap()) // adds the bit that square represents to the bitmap
	b.setPieceBitmap(pieceColor, piecePieceType, bitmap)

	if err != nil {
		return nil
	} else {
		return pieceOriginallyAtSqaure
	}
}

func (b *board) isAnyPieceBetween(srcSquare, dstSquare Square) bool {
	var curr Square = srcSquare
	var dir Direction = srcSquare.DirectionTo(dstSquare)
	var err error

	for {
		curr, _ = curr.Step(dir)

		if curr == dstSquare {
			break
		}

		_, err = b.GetPieceAt(curr)
		if err == nil {
			// found a piece
			return true
		}
	}

	return false
}

func (b *board) IsValidMove(m Move) error {
	var srcSquare, dstSquare Square
	var srcPiece *Piece
	var dstPiece *Piece
	var err error

	srcSquare = m.GetSrcSquare()
	dstSquare = m.GetDstSquare()

	// check that a piece exists on the src square
	srcPiece, err = b.GetPieceAt(srcSquare)
	if err != nil {
		return fmt.Errorf("no piece on src square: %s", srcSquare.GetName())
	}

	// check that the piece being moved is the same color as the side to move
	if srcPiece.GetColor() != b.GetTurn() {
		return fmt.Errorf("%s can't move a piece with color %s", b.GetTurn(), srcPiece.GetColor())
	}

	// check that if there a piece on the destination square, that the captured piece's color is the opposite of the side to move
	dstPiece, err = b.GetPieceAt(dstSquare)
	if err == nil && dstPiece.GetColor() == b.GetTurn() {
		return fmt.Errorf("%s can't move a piece onto another %s piece", b.GetTurn(), dstPiece.GetColor())
	}

	// check that the src piece is moving according to its movement or capture set (e.g. a knight moves in L shape, pawns capture diagonally)
	if srcPiece.IsValidMovement(srcSquare, dstSquare) != nil && srcPiece.IsValidCapture(srcSquare, dstSquare) != nil {
		return fmt.Errorf("move: %s falls outside of %s's movement or capture capabilities", m, srcPiece)
	}

	// if the src piece is not a knight, check that the src piece is not jumping over other pieces
	if srcPiece.GetPieceType() != KNIGHT && b.isAnyPieceBetween(srcSquare, dstSquare) {
		return fmt.Errorf("piece: %s can't move from %s to %s because it would jump over another piece", srcPiece, srcSquare.GetName(), dstSquare.GetName())
	}

	// if the src piece is a pawn, and if the pawn is capturing a piece, check that the destination square either has a piece, or is the en-passent square
	if srcPiece.GetPieceType() == PAWN && srcPiece.IsValidCapture(srcSquare, dstSquare) == nil && dstPiece == nil && dstSquare.ToBitMap() != b.enPassentBitMap {
		return fmt.Errorf("pawn can't capture a piece that does not exist on the destination square, without en-passent")
	}

	// if the src piece is a pawn, and if the pawn is just moving (not a capture), check that the destination square has no piece
	if srcPiece.GetPieceType() == PAWN && srcPiece.IsValidMovement(srcSquare, dstSquare) == nil && dstPiece != nil {
		return fmt.Errorf("pawn can't move vertically onto a piece")
	}

	// if the src piece is a pawn, and if the pawn is entering the 1st or 8th rank, check that promotion piece is valid
	if srcPiece.GetPieceType() == PAWN && (dstSquare.GetRank() == 1 || dstSquare.GetRank() == 8) && (m.GetPromotionPieceType() == nil || !m.GetPromotionPieceType().IsValidPromotionPiece()) {
		return fmt.Errorf("pawn can't enter 1st or 8th rank without promoting to a valid piece")
	}

	// if the src piece is a king, and if the king is trying to castle...
	if srcPiece.GetPieceType() == KING && srcSquare.DistanceSquaredTo(dstSquare) == 4 {
		castlingDirection := srcSquare.DirectionTo(dstSquare)

		// ..check that the king has castling rights on that side of the board
		if !b.isCastlingAllowed(b.GetTurn(), srcSquare, dstSquare) {
			return fmt.Errorf("king is not allowed to castle on that side of the board")
		}

		// ..check that all squares between the rook and king are empty
		piece, pieceSquare := b.getClosestPieceInDirection(srcSquare, castlingDirection)
		if piece == nil {
			return fmt.Errorf("king can't castle in a direction where there is no rook")
		}

		if !pieceSquare.IsOnLeftOrRightEdge() {
			return fmt.Errorf("king can't castle in a direction where squares between the king and the corner are not empty")
		}

		if piece.GetPieceType() != ROOK {
			return fmt.Errorf("king can't castle in a direction where is no rook")
		}

		// ..check that king is not castling 'out of', or 'through' check ('into' check case is captured by the last validation of this function)
		if b.isCheck() {
			return fmt.Errorf("king can't castle out of check")
		}

		bCopy := b.Copy()
		kingStartSquare := srcSquare
		kingEndSquare, _ := srcSquare.Step(castlingDirection)
		bCopy.makeUnsafe(NewMove(kingStartSquare, kingEndSquare).Build())
		if bCopy.isCheck() {
			return fmt.Errorf("king can't castle through check")
		}
	}

	// check that the resulting position does not put the current king in check (includes 'into' check castling case)
	bCopy := b.Copy()
	bCopy.makeUnsafe(m)
	bCopy.setTurn(b.GetTurn())
	if bCopy.isCheck() {
		return fmt.Errorf("%s can't make a move %s that puts king in check", m.String(), b.GetTurn())
	}

	// move passes all checks
	return nil
}

func (b *board) isCastlingAllowed(c Color, srcSquare, dstSquare Square) bool {
	switch c {
	case WHITE:
		if dstSquare.GetFile() > srcSquare.GetFile() {
			return b.whiteKingside
		} else {
			return b.whiteQueenside
		}
	case BLACK:
		if dstSquare.GetFile() > srcSquare.GetFile() {
			return b.blackKingside
		} else {
			return b.blackQueenSide
		}
	}

	panic(fmt.Sprintf("Unhandled switch case: %s", c))
}

func (b *board) isCheck() bool {
	var square Square
	var piece *Piece
	var pieceSquare Square
	var myColor Color = b.GetTurn()
	var oppColor Color = myColor.Opposite()

	if myColor == WHITE {
		square = b.whiteKingBitMap.ToSquare()
	} else {
		square = b.blackKingBitMap.ToSquare()
	}

	// look north
	piece, pieceSquare = b.getClosestPieceInDirection(square, NORTH)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == ROOK:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look northeast
	piece, pieceSquare = b.getClosestPieceInDirection(square, NORTHEAST)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == BISHOP:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		case myColor == WHITE && piece.GetPieceType() == PAWN && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look east
	piece, pieceSquare = b.getClosestPieceInDirection(square, EAST)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == ROOK:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look southeast
	piece, pieceSquare = b.getClosestPieceInDirection(square, SOUTHEAST)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == BISHOP:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		case myColor == BLACK && piece.GetPieceType() == PAWN && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look south
	piece, pieceSquare = b.getClosestPieceInDirection(square, SOUTH)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == ROOK:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look southwest
	piece, pieceSquare = b.getClosestPieceInDirection(square, SOUTHWEST)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == BISHOP:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		case myColor == BLACK && piece.GetPieceType() == PAWN && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look west
	piece, pieceSquare = b.getClosestPieceInDirection(square, WEST)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == ROOK:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look northwest
	piece, pieceSquare = b.getClosestPieceInDirection(square, NORTHWEST)
	if piece != nil && piece.GetColor() == oppColor {
		switch {
		case piece.GetPieceType() == QUEEN:
			return true
		case piece.GetPieceType() == BISHOP:
			return true
		case piece.GetPieceType() == KING && NumSteps(square, pieceSquare) == 1:
			return true
		case myColor == WHITE && piece.GetPieceType() == PAWN && NumSteps(square, pieceSquare) == 1:
			return true
		}
	}

	// look for knight checks
	row := square.GetRow()
	col := square.GetCol()

	piece = b.getKnightAt(row-1, col+2)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	piece = b.getKnightAt(row-1, col-2)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	piece = b.getKnightAt(row+1, col+2)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	piece = b.getKnightAt(row+1, col-2)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	piece = b.getKnightAt(row-2, col+1)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	piece = b.getKnightAt(row-2, col-1)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	piece = b.getKnightAt(row+2, col+1)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	piece = b.getKnightAt(row+2, col-1)
	if piece != nil && piece.GetColor() == oppColor && piece.GetPieceType() == KNIGHT {
		return true
	}

	// no checks found
	return false
}

func (b *board) getKnightAt(row, col int) *Piece {
	var piece *Piece

	if row < 0 || row >= 8 {
		return nil
	}

	if col < 0 || col >= 8 {
		return nil
	}

	piece, _ = b.GetPieceAt(GetSquareFromCoord(row, col))
	return piece
}

func (b *board) getClosestPieceInDirection(s Square, d Direction) (*Piece, Square) {
	var curr Square = s
	var err error
	var piece *Piece

	for {
		curr, err = curr.Step(d)

		if err != nil {
			// did not run into any pieces, move off the board
			return nil, curr
		}

		piece, err = b.GetPieceAt(curr)

		if err == nil {
			// ran into a piece
			return piece, curr
		}
	}
}

func (b *board) Copy() Board {
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

		ply: b.ply,
	}
}

func (b *board) String() string {
	var square Square = 1
	var piece *Piece
	var err error

	ret := ""
	for i := 0; i < 64; i++ {
		if piece, err = b.GetPieceAt(square); err != nil {
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

		ply: 0,
	}
}
