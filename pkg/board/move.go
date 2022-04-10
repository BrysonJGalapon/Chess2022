package board

import "fmt"

type Move interface {
	GetSrcSquare() Square
	GetDstSquare() Square
	GetPromotionPieceType() PieceType
	AddPromotionPieceType(PieceType) Move
	IsEmpty() bool
}

type move struct {
	srcSquare          Square
	dstSquare          Square
	promotionPieceType *PieceType
}

func (m *move) GetSrcSquare() Square {
	return m.srcSquare
}

func (m *move) GetDstSquare() Square {
	return m.dstSquare
}

func (m *move) GetPromotionPieceType() PieceType {
	if m.promotionPieceType == nil {
		panic("no promotion piece type in this move")
	}

	return *m.promotionPieceType
}

func (m *move) AddPromotionPieceType(promotionPieceType PieceType) Move {
	return NewMove(m.srcSquare, m.dstSquare).PromotionPieceType(promotionPieceType).Build()
}

func (m *move) IsEmpty() bool {
	return false
}

type MoveBuilder interface {
	PromotionPieceType(PieceType) MoveBuilder
	Build() Move
}

type moveBuilder struct {
	srcSquare          Square
	dstSquare          Square
	promotionPieceType *PieceType
}

func (mv *moveBuilder) PromotionPieceType(promotionPieceType PieceType) MoveBuilder {
	mv.promotionPieceType = &promotionPieceType
	return mv
}

func (mv *moveBuilder) Build() Move {
	return &move{mv.srcSquare, mv.dstSquare, mv.promotionPieceType}
}

func NewMove(srcSquare, dstSquare Square) MoveBuilder {
	return &moveBuilder{srcSquare, dstSquare, nil}
}

// empty move
type emptyMove struct{}

func (em *emptyMove) GetSrcSquare() Square {
	panic("can't call GetSrcSquare on empty move")
}

func (em *emptyMove) GetDstSquare() Square {
	panic("can't call GetDstSquare on empty move")
}

func (em *emptyMove) GetPromotionPieceType() PieceType {
	panic("can't call GetPromotionPieceType on empty move")
}

func (em *emptyMove) AddPromotionPieceType(promotionPieceType PieceType) Move {
	panic("can't call GetPromotionPieceType on empty move")
}

func (em *emptyMove) IsEmpty() bool {
	return true
}

func GetEmptyMove() Move {
	return &emptyMove{}
}

// castling move
type castlingMove struct {
	color Color
	side  Side
}

func (cm *castlingMove) GetSrcSquare() Square {
	switch cm.color {
	case WHITE:
		return GetSquareFromString("E1")
	case BLACK:
		return GetSquareFromString("E8")
	}

	panic(fmt.Sprintf("unhandled switch case: %s", cm))
}

func (cm *castlingMove) GetDstSquare() Square {
	switch cm.color {
	case WHITE:
		switch cm.side {
		case KINGSIDE:
			return GetSquareFromString("F1")
		case QUEENSIDE:
			return GetSquareFromString("C1")
		}
	case BLACK:
		switch cm.side {
		case KINGSIDE:
			return GetSquareFromString("F8")
		case QUEENSIDE:
			return GetSquareFromString("C8")
		}
	}

	panic(fmt.Sprintf("unhandled switch case: %s", cm))
}

func (cm *castlingMove) GetPromotionPieceType() PieceType {
	panic("can't call GetPromotionPieceType on castling move")
}

func (cm *castlingMove) IsEmpty() bool {
	return false
}

func (cm *castlingMove) AddPromotionPieceType(promotionPieceType PieceType) Move {
	panic("can't call GetPromotionPieceType on castling move")
}

func (cm *castlingMove) String() string {
	return fmt.Sprintf("{color: %d, side: %d}", cm.color, cm.side)
}

func GetCastlingMove(color Color, side Side) Move {
	return &castlingMove{color, side}
}
