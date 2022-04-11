package board

import "fmt"

type Move interface {
	GetSrcSquare() Square
	GetDstSquare() Square
	GetPromotionPieceType() PieceType
	AddPromotionPieceType(PieceType) Move
	String() string
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

func (m *move) String() string {
	ret := "{"
	ret += fmt.Sprintf("%s -> %s", m.srcSquare.GetName(), m.dstSquare.GetName())
	if m.promotionPieceType != nil {
		ret += fmt.Sprintf(" =%s", m.promotionPieceType.String())
	}
	ret += "}"
	return ret
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

func (em *emptyMove) String() string {
	return "{emptyMove}"
}

func GetEmptyMove() Move {
	return &emptyMove{}
}
