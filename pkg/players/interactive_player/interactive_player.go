package interactive_player

import (
	"bufio"
	"fmt"
	b "galapb/chess2022/pkg/board"
	"os"
	"strings"
)

type InteractivePlayer struct {
	prompt   chan b.Move
	response chan b.Move
}

func New(prompt chan b.Move, response chan b.Move) *InteractivePlayer {
	return &InteractivePlayer{prompt, response}
}

func (ip *InteractivePlayer) Start(board b.Board, quit chan bool) {
	var err error

	for {
		select {
		case <-quit:
			return
		default:
			move := <-ip.prompt
			if err = board.Make(move); err != nil {
				panic(err)
			}

			response := ip.getMove(board)
			board.Make(response)

			ip.response <- response
		}
	}
}

func (rp *InteractivePlayer) getMove(board b.Board) b.Move {
	// get move from user

	var srcSquare b.Square
	var dstSquare b.Square
	var ok bool

	for {
		fmt.Printf("> Enter your move: ")
		in := bufio.NewReader(os.Stdin)

		moveString, _ := in.ReadString('\n')
		words := strings.Fields(moveString)

		if len(words) != 2 && len(words) != 3 {
			fmt.Printf("\t%s is not a valid move. Try again... \n", moveString)
			continue
		}

		if srcSquare, ok = b.GetSquareFromStringNotExistsOkay(words[0]); !ok {
			fmt.Printf("\t%s is not a valid square. Try again... \n", words[0])
			continue
		}

		if dstSquare, ok = b.GetSquareFromStringNotExistsOkay(words[1]); !ok {
			fmt.Printf("\t%s is not a valid square. Try again... \n", words[1])
			continue
		}

		move := b.NewMove(srcSquare, dstSquare).Build()

		if err := board.IsValidMove(move); err != nil {
			fmt.Printf("\t%s is not a valid move. Error: %s. Try again... \n", move, err)
			continue
		}

		return move
	}
}
