package board

type Status uint8

const (
	UNDETERMINED Status = iota
	CHECKMATE
	STALEMATE
	INSUFFICIENT_MATERIAL
	FIFTY_MOVE_RULE
	THREEFOLD_REPETITION
)
