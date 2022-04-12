package minimax_player

import (
	b "galapb/chess2022/pkg/board"
	"math/rand"
)

type MiniMaxPlayer struct {
	prompt   chan b.Move
	response chan b.Move
	maxDepth int
}

var ALL_POSSIBLE_MOVES []b.Move = make([]b.Move, 0)

func init() {
	var srcSquare b.Square = 1
	var dstSquare b.Square = 1
	var move b.Move

	for i := 0; i < 64; i++ {
		dstSquare = 1

		for j := 0; j < 64; j++ {
			move = b.NewMove(srcSquare, dstSquare).Build()
			ALL_POSSIBLE_MOVES = append(ALL_POSSIBLE_MOVES, move)

			dstSquare <<= 1
		}

		srcSquare <<= 1
	}
}

func New(prompt chan b.Move, response chan b.Move) *MiniMaxPlayer {
	return &MiniMaxPlayer{prompt, response, 2}
}

func (mp *MiniMaxPlayer) Start(board b.Board, quit chan bool) {
	var err error

	for {
		select {
		case <-quit:
			return
		default:
			move := <-mp.prompt
			if err = board.Make(move); err != nil {
				panic(err)
			}

			response := mp.getMove(board)
			board.Make(response)

			mp.response <- response
		}
	}
}

func (rp *MiniMaxPlayer) getMove(board b.Board) b.Move {
	if board.GetTurn() == b.WHITE {
		move, _ := rp.max(board, rp.maxDepth)
		return move
	} else {
		move, _ := rp.min(board, rp.maxDepth)
		return move
	}
}

func (rp *MiniMaxPlayer) heuristic(board b.Board) float64 {
	switch board.GetStatus() {
	case b.CHECKMATE:
		if board.GetTurn() == b.BLACK {
			return 1000 // black is checkmated
		} else {
			return -1000 // white is checkmated
		}
	case b.STALEMATE:
		return 0
	case b.INSUFFICIENT_MATERIAL:
		return 0
	case b.FIFTY_MOVE_RULE:
		return 0
	case b.THREEFOLD_REPETITION:
		return 0
	}

	var h float64 = 0.0

	h += 9.0 * float64(board.GetNumOf(b.WHITE, b.QUEEN))
	h += 3.2 * float64(board.GetNumOf(b.WHITE, b.BISHOP))
	h += 3.1 * float64(board.GetNumOf(b.WHITE, b.KNIGHT))
	h += 5.0 * float64(board.GetNumOf(b.WHITE, b.ROOK))
	h += 1.0 * float64(board.GetNumOf(b.WHITE, b.PAWN))

	h -= 9.0 * float64(board.GetNumOf(b.BLACK, b.QUEEN))
	h -= 3.2 * float64(board.GetNumOf(b.BLACK, b.BISHOP))
	h -= 3.1 * float64(board.GetNumOf(b.BLACK, b.KNIGHT))
	h -= 5.0 * float64(board.GetNumOf(b.BLACK, b.ROOK))
	h -= 1.0 * float64(board.GetNumOf(b.BLACK, b.PAWN))

	return h
}

/**
Returns a move that minimizes the heuristic
*/
func (rp *MiniMaxPlayer) min(board b.Board, depth int) (b.Move, float64) {
	var bCopy b.Board
	var err error
	var h float64
	var p *b.Piece

	var bestH float64 = 1000
	var bestMoves []b.Move = make([]b.Move, 0)

	for _, move := range ALL_POSSIBLE_MOVES {
		bCopy = board.Copy()

		if p, err = bCopy.GetPieceAt(move.GetSrcSquare()); err != nil {
			continue // ignore moves which do not move a piece
		}

		if p.GetPieceType() == b.PAWN && (move.GetDstSquare().GetRank() == 1 || move.GetDstSquare().GetRank() == 8) {
			// always promote to a queen
			move = move.AddPromotionPieceType(b.QUEEN)
		}

		if err = bCopy.Make(move); err != nil {
			continue // ignore invalid moves
		}

		if depth == 1 {
			h = rp.heuristic(bCopy)
		} else {
			_, h = rp.max(bCopy, depth-1)
		}

		if h < bestH {
			bestH = h
			bestMoves = []b.Move{move}
		} else if h == bestH {
			bestMoves = append(bestMoves, move)
		}
	}

	return GetRandomMove(bestMoves), bestH
}

/**
Returns a move that maximizes the heuristic
*/
func (rp *MiniMaxPlayer) max(board b.Board, depth int) (b.Move, float64) {
	var bCopy b.Board
	var err error
	var h float64

	var bestH float64 = -1000
	var bestMoves []b.Move = make([]b.Move, 0)

	for _, move := range ALL_POSSIBLE_MOVES {
		bCopy = board.Copy()
		if err = bCopy.Make(move); err != nil {
			continue // ignore invalid moves
		}

		if depth == 1 {
			h = rp.heuristic(bCopy)
		} else {
			_, h = rp.min(bCopy, depth-1)
		}

		if h > bestH {
			bestH = h
			bestMoves = []b.Move{move}
		} else if h == bestH {
			bestMoves = append(bestMoves, move)
		}
	}

	return GetRandomMove(bestMoves), bestH
}

func GetRandomMove(moves []b.Move) b.Move {
	i := rand.Intn(len(moves))
	return moves[i]
}
