package evaluator

import (
	"galapb/chess2022/pkg/game"
	neat_player "galapb/chess2022/pkg/players/neat_player/player"
	"galapb/chess2022/pkg/players/player"
	"galapb/chess2022/pkg/players/random_player"
	"galapb/chess2022/pkg/time_control"
	"log"

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
	for i, org := range pop.Organisms {
		log.Printf("Evaluating Organism %d out of %d", i, len(pop.Organisms))

		tc := time_control.Builder().Minutes(3).Build()

		var whitePlayer player.Player
		var blackPlayer player.Player
		var result game.Result
		var g game.Game
		var score float64 = 0

		// play a game as white
		whitePlayer = neat_player.New(org)
		blackPlayer = random_player.New()
		g = game.New(tc, whitePlayer, blackPlayer).Verbose(true).PlyLimit(1000).Build()
		result, _ = g.Run()
		switch result {
		case game.BLACK_WINS:
			log.Println("Organism lost as white")
			score += 0.0
		case game.WHITE_WINS:
			log.Println("Organism won as white")
			score += 0.5
		case game.GAME_DRAWN:
			log.Println("Organism drew as white")
			score += 0.25
		}

		// ... and as black
		whitePlayer = random_player.New()
		blackPlayer = neat_player.New(org)
		g = game.New(tc, whitePlayer, blackPlayer).Verbose(true).PlyLimit(1000).Build()
		result, _ = g.Run()
		switch result {
		case game.BLACK_WINS:
			log.Println("Organism won as black")
			score += 0.5
		case game.WHITE_WINS:
			log.Println("Organism lost as black")
			score += 0.0
		case game.GAME_DRAWN:
			log.Println("Organism drew as black")
			score += 0.25
		}

		org.Fitness = score
	}

	return nil
}
