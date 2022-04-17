package main

import (
	"galapb/chess2022/pkg/game"
	"galapb/chess2022/pkg/players/minimax_player"
	"galapb/chess2022/pkg/players/player"
	"galapb/chess2022/pkg/time_control"
)

func main() {
	// build the players
	var whitePlayer player.Player = minimax_player.New()
	var blackPlayer player.Player = minimax_player.New()

	// build the game
	var timeControl time_control.TimeControl = time_control.Builder().Minutes(3).Build()
	var g game.Game = game.New(timeControl, whitePlayer, blackPlayer).Build()

	// run the game
	g.Run()
}
