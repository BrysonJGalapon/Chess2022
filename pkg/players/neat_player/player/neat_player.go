package player

import (
	b "galapb/chess2022/pkg/board"
	"math/rand"

	"github.com/yaricom/goNEAT/v2/neat/genetics"
)

type NeatPlayer struct {
	prompt   chan b.Move
	response chan b.Move
	org      *genetics.Organism
}

func New(org *genetics.Organism) *NeatPlayer {
	return &NeatPlayer{nil, nil, org}
}

func (np *NeatPlayer) Init(prompt chan b.Move, response chan b.Move) {
	np.prompt = prompt
	np.response = response
}

func (np *NeatPlayer) Start(board b.Board, quit chan bool) {
	var err error

	for {
		select {
		case <-quit:
			return
		default:
			move := <-np.prompt
			if err = board.Make(move); err != nil {
				panic(err)
			}

			response := np.getMove(board)
			board.Make(response)

			np.response <- response
		}
	}
}

func (np *NeatPlayer) getMove(board b.Board) b.Move {
	var inputs []float64
	var outputs []float64
	net := np.org.Phenotype // Neural Network (NN)

	// Send inputs to NN
	inputs = getNetInputsFromBoard(board)
	net.LoadSensors(inputs)

	// Run the NN
	net.Activate()

	// Get output from NN
	outputs = net.ReadOutputs()

	// Translate NN output to a board move
	move, isRandomMove := getBoardMoveFromNetOutputs(outputs, board)

	// Penalize the fitness of the organism for making random moves
	if isRandomMove {
		np.org.Fitness -= 0.001
	}

	return move
}

func getNetInputsFromBoard(board b.Board) []float64 {
	var inputs []float64 = []float64{
		float64(board.GetPieceBitmap(b.WHITE, b.KING)),
		float64(board.GetPieceBitmap(b.WHITE, b.QUEEN)),
		float64(board.GetPieceBitmap(b.WHITE, b.BISHOP)),
		float64(board.GetPieceBitmap(b.WHITE, b.KNIGHT)),
		float64(board.GetPieceBitmap(b.WHITE, b.ROOK)),
		float64(board.GetPieceBitmap(b.WHITE, b.PAWN)),
		float64(board.GetPieceBitmap(b.BLACK, b.KING)),
		float64(board.GetPieceBitmap(b.BLACK, b.QUEEN)),
		float64(board.GetPieceBitmap(b.BLACK, b.BISHOP)),
		float64(board.GetPieceBitmap(b.BLACK, b.KNIGHT)),
		float64(board.GetPieceBitmap(b.BLACK, b.ROOK)),
		float64(board.GetPieceBitmap(b.BLACK, b.PAWN)),

		// TODO castling rights, en-passent, + other inputs
	}

	return inputs
}

func getBoardMoveFromNetOutputs(outputs []float64, board b.Board) (b.Move, bool) {
	srcIndexes, dstIndexes := getSrcAndDstIndexesFromOutputs(outputs)
	return getBoardMoveFromIndexes(srcIndexes, dstIndexes, board)
}

func getBoardMoveFromIndexes(srcIndexes, dstIndexes []int, board b.Board) (b.Move, bool) {
	var srcSquare b.Square
	var dstSquare b.Square
	var move b.Move = nil
	var err error = nil

	Shuffle(srcIndexes)
	Shuffle(dstIndexes)

	for _, srcIndex := range srcIndexes {
		for _, dstIndex := range dstIndexes {
			srcSquare = b.GetSquareFromIndex(srcIndex)
			dstSquare = b.GetSquareFromIndex(dstIndex)

			move = b.NewMove(srcSquare, dstSquare).Build()

			if err = board.IsValidMove(move); err == nil {
				return move, false
			}
		}
	}

	return getRandomMove(board), true
}

func getSrcAndDstIndexesFromOutputs(outputs []float64) ([]int, []int) {
	var srcIndexes []int = make([]int, 0)
	var srcIndexScore float64 = 0

	var dstIndexes []int = make([]int, 0)
	var dstIndexScore float64 = 0

	for i, output := range outputs {
		if i < 64 {
			if i == 0 || output > srcIndexScore {
				srcIndexes = []int{i}
				srcIndexScore = output
			} else if output == srcIndexScore {
				srcIndexes = append(srcIndexes, i)
			}
		} else {
			if i == 0 || output > dstIndexScore {
				dstIndexes = []int{i - 64}
				dstIndexScore = output
			} else if output == dstIndexScore {
				dstIndexes = append(dstIndexes, i-64)
			}
		}
	}

	return srcIndexes, dstIndexes
}

func getRandomMove(board b.Board) b.Move {
	srcSquare := b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
	dstSquare := b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
	move := b.NewMove(srcSquare, dstSquare).Build()

	for board.IsValidMove(move) != nil {
		srcSquare = b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
		dstSquare = b.GetSquareFromCoord(rand.Intn(8), rand.Intn(8))
		move = b.NewMove(srcSquare, dstSquare).Build()

		// add a promotion piece if promoting a pawn
		if piece, _ := board.GetPieceAt(srcSquare); piece != nil && piece.GetPieceType() == b.PAWN && (dstSquare.GetRank() == 1 || dstSquare.GetRank() == 8) {
			move = move.AddPromotionPieceType(getRandomPromotionPieceType())
		}
	}

	return move
}

func getRandomPromotionPieceType() b.PieceType {
	i := rand.Intn(len(b.PROMOTION_PIECE_TYPES))
	return b.PROMOTION_PIECE_TYPES[i]
}
