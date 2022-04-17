package evaluator

import (
	"galapb/chess2022/pkg/game"
	"galapb/chess2022/pkg/players/minimax_player"
	neat_player "galapb/chess2022/pkg/players/neat_player/player"
	"galapb/chess2022/pkg/players/player"
	"galapb/chess2022/pkg/time_control"

	"github.com/yaricom/goNEAT/v2/experiment"
	"github.com/yaricom/goNEAT/v2/neat"
	"github.com/yaricom/goNEAT/v2/neat/genetics"
)

type neatPlayerEvaluator struct {
	// The output path to store execution results
	OutputPath string
}

func NewNeatPlayerGenerationEvaluator(outputPath string) experiment.GenerationEvaluator {
	return &neatPlayerEvaluator{OutputPath: outputPath}
}

// This method evaluates one epoch for given population and prints results into output directory if any
func (ne *neatPlayerEvaluator) GenerationEvaluate(pop *genetics.Population, epoch *experiment.Generation, context *neat.Options) (err error) {
	for _, org := range pop.Organisms {
		tc := time_control.Builder().Minutes(3).Build()

		var whitePlayer player.Player
		var blackPlayer player.Player
		var result game.Result
		var g game.Game
		var score float64 = 0

		// play a game as white
		whitePlayer = neat_player.New(org)
		blackPlayer = minimax_player.New()
		g = game.New(tc, whitePlayer, blackPlayer).Build()
		result, _ = g.Run()
		switch result {
		case game.BLACK_WINS:
			score += 0.0
		case game.WHITE_WINS:
			score += 1.0
		case game.GAME_DRAWN:
			score += 0.5
		}

		// ... and as black
		whitePlayer = minimax_player.New()
		blackPlayer = neat_player.New(org)
		g = game.New(tc, whitePlayer, blackPlayer).Build()
		result, _ = g.Run()
		switch result {
		case game.BLACK_WINS:
			score += 1.0
		case game.WHITE_WINS:
			score += 0.0
		case game.GAME_DRAWN:
			score += 0.5
		}

		org.Fitness = score
	}

	return nil
}
