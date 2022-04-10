package board

type Move interface {
	GetSrcSquare() Square
	GetDstSquare() Square
	GetPromotionPieceType() PieceType
	IsEmpty() bool
}

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

func (em *emptyMove) IsEmpty() bool {
	return true
}

func GetEmptyMove() Move {
	return &emptyMove{}
}
