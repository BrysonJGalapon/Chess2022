package player

import (
	"galapb/chess2022/pkg/board"
	b "galapb/chess2022/pkg/board"

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
	return getBoardMoveFromNetOutputs(outputs)
}

func getNetInputsFromBoard(b board.Board) []float64 {
	var inputs []float64 = []float64{
		float64(b.GetPieceBitmap(board.WHITE, board.KING)),
		float64(b.GetPieceBitmap(board.WHITE, board.QUEEN)),
		float64(b.GetPieceBitmap(board.WHITE, board.BISHOP)),
		float64(b.GetPieceBitmap(board.WHITE, board.KNIGHT)),
		float64(b.GetPieceBitmap(board.WHITE, board.ROOK)),
		float64(b.GetPieceBitmap(board.WHITE, board.PAWN)),
		float64(b.GetPieceBitmap(board.BLACK, board.KING)),
		float64(b.GetPieceBitmap(board.BLACK, board.QUEEN)),
		float64(b.GetPieceBitmap(board.BLACK, board.BISHOP)),
		float64(b.GetPieceBitmap(board.BLACK, board.KNIGHT)),
		float64(b.GetPieceBitmap(board.BLACK, board.ROOK)),
		float64(b.GetPieceBitmap(board.BLACK, board.PAWN)),

		// TODO castling rights, en-passent, + other inputs
	}

	return inputs
}

func getBoardMoveFromNetOutputs(outputs []float64) board.Move {
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
				dstIndexes = []int{i}
				dstIndexScore = output
			} else if output == dstIndexScore {
				dstIndexes = append(dstIndexes, i)
			}
		}
	}

	var srcIndex int = getRandomInt(srcIndexes)
	var dstIndex int = getRandomInt(dstIndexes)

	var srcSquare b.Square
	var dstSquare b.Square

	srcSquare = b.GetSquareFromIndex(srcIndex)
	dstSquare = b.GetSquareFromIndex(dstIndex)

	return nil
}
