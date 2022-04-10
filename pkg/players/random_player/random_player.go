package random_player

import "time"
import "math/rand"
import b "galapb/chess2022/pkg/board"

func init() {
	rand.Seed(33423432)
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

			rp.response <- response
		}
	}
}

func (rp *RandomPlayer) getMove(board b.Board) b.Move {
	time.Sleep(3 * time.Second)
	return b.GetEmptyMove()
}
