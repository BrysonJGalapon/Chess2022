package main

import (
	"fmt"
	"galapb/chess2022/pkg/board"
	"galapb/chess2022/pkg/game"
	"galapb/chess2022/pkg/players/player"
	"galapb/chess2022/pkg/players/random_player"
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
	var whitePlayer player.Player = random_player.New(whitePrompt, whiteResponse)
	var blackPlayer player.Player = random_player.New(blackPrompt, blackResponse)

	// quit channels
	var whiteQuit chan bool = make(chan bool)
	var blackQuit chan bool = make(chan bool)

	// build the game
	var timeControl time_control.TimeControl = time_control.Builder().Minutes(3).Build()
	var game game.Game = game.New(timeControl, whitePlayer, blackPlayer).Build()

	// start the players
	go whitePlayer.Start(game.GetBoard().Copy(), whiteQuit)
	go blackPlayer.Start(game.GetBoard().Copy(), blackQuit)

	// run the game
	var move board.Move = board.GetEmptyMove()
	for !game.IsOver() && game.GetBoard().GetPly() < 200 {
		whitePrompt <- move
		move = <-whiteResponse

		if err := game.GetBoard().Make(move); err != nil {
			panic(fmt.Sprintf("invalid move by white: %s", err))
		}

		log.Printf("White made move: %s", move)
		log.Printf("Board:\n%s", game.GetBoard().String())

		blackPrompt <- move
		move = <-blackResponse

		if err := game.GetBoard().Make(move); err != nil {
			panic(fmt.Sprintf("invalid move by black: %s", err))
		}

		log.Printf("Black made move: %s", move)
		log.Printf("Board:\n%s", game.GetBoard().String())
	}
}
