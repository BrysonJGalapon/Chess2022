package main

import (
	"fmt"
	"galapb/chess2022/pkg/board"
	"galapb/chess2022/pkg/game"
	"galapb/chess2022/pkg/players/interactive_player"
	"galapb/chess2022/pkg/players/minimax_player"
	"galapb/chess2022/pkg/players/player"
	"galapb/chess2022/pkg/time_control"
	"log"
)

func main() {
	// move channels
	var whitePrompt chan board.Move = make(chan board.Move, 1)
	var whiteResponse chan board.Move = make(chan board.Move, 1)
	var blackPrompt chan board.Move = make(chan board.Move, 1)
	var blackResponse chan board.Move = make(chan board.Move, 1)

	// build the players
	var whitePlayer player.Player = interactive_player.New(whitePrompt, whiteResponse)
	var blackPlayer player.Player = minimax_player.New(blackPrompt, blackResponse)

	// quit channels
	var whiteQuit chan bool = make(chan bool)
	var blackQuit chan bool = make(chan bool)

	// build the game
	var timeControl time_control.TimeControl = time_control.Builder().Minutes(3).Build()
	var g game.Game = game.New(timeControl, whitePlayer, blackPlayer).Build()

	// start the players
	go whitePlayer.Start(g.GetBoard().Copy(), whiteQuit)
	go blackPlayer.Start(g.GetBoard().Copy(), blackQuit)

	var result game.Result
	var reason game.Reason

	// run the game
	var move board.Move = board.GetEmptyMove()
	for {
		log.Printf("Board:\n%s", g.GetBoard().String())

		if result, reason = g.GetResult(); result != game.UNDETERMINED {
			// game is over
			break
		}

		whitePrompt <- move
		move = <-whiteResponse

		if err := g.GetBoard().Make(move); err != nil {
			panic(fmt.Sprintf("invalid move by white: %s", err))
		}

		log.Printf("White made move: %s", move)
		log.Printf("Board:\n%s", g.GetBoard().String())

		if result, reason = g.GetResult(); result != game.UNDETERMINED {
			// game is over
			break
		}

		blackPrompt <- move
		move = <-blackResponse

		if err := g.GetBoard().Make(move); err != nil {
			panic(fmt.Sprintf("invalid move by black: %s", err))
		}

		log.Printf("Black made move: %s", move)
	}

	// print results of the game
	log.Printf("%s due to %s", result, reason)
}
