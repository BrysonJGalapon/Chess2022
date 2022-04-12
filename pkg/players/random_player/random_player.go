package random_player

import (
	"math/rand"

	b "galapb/chess2022/pkg/board"
)

func init() {
	rand.Seed(3629)
}

type RandomPlayer struct {
	prompt   chan b.Move
	response chan b.Move
}

func New(prompt chan b.Move, response chan b.Move) *RandomPlayer {
	return &RandomPlayer{prompt, response}
}

func (rp *RandomPlayer) Start(board b.Board, quit chan bool) {
	var err error

	for {
		select {
		case <-quit:
			return
		default:
			move := <-rp.prompt
			if err = board.Make(move); err != nil {
				panic(err)
			}

			response := rp.getMove(board)
			board.Make(response)

			rp.response <- response
		}
	}
}

func (rp *RandomPlayer) getMove(board b.Board) b.Move {
	srcSquare := b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
	dstSquare := b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
	move := b.NewMove(srcSquare, dstSquare).Build()

	for board.IsValidMove(move) != nil {
		srcSquare = b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
		dstSquare = b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
		move = b.NewMove(srcSquare, dstSquare).Build()

		// add a promotion piece if promoting a pawn
		if piece, _ := board.GetPieceAt(srcSquare); piece != nil && piece.GetPieceType() == b.PAWN && (dstSquare.GetRank() == 1 || dstSquare.GetRank() == 8) {
			move = move.AddPromotionPieceType(rp.getRandomPromotionPieceType())
		}
	}

	return move
}

func (rp *RandomPlayer) getRandomPromotionPieceType() b.PieceType {
	i := rand.Intn(len(b.PROMOTION_PIECE_TYPES))
	return b.PROMOTION_PIECE_TYPES[i]
}
