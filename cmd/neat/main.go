package main

import (
	"context"
	"fmt"
	"galapb/chess2022/pkg/players/neat_player/evaluator"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yaricom/goNEAT/v2/experiment"
	"github.com/yaricom/goNEAT/v2/neat"
	"github.com/yaricom/goNEAT/v2/neat/genetics"
)

const NEAT_PARAMS_FILE string = "./pkg/players/neat_player/player/config/params.yml"
const GENOME_FILE string = "./pkg/players/neat_player/player/config/startgenes.yml"
const OUT_DIR string = "./pkg/players/neat_player/player/data"

func main() {
	// Load Neat Options
	params, err := os.Open(NEAT_PARAMS_FILE)
	if err != nil {
		log.Fatal("Failed to open context configuration file: ", err)
	}

	neatOptions, err := neat.LoadYAMLOptions(params)
	if err != nil {
		log.Fatal("Failed to load YAML options: ", err)
	}

	// Load Start Genome
	genomeFile, err := os.Open(GENOME_FILE)
	if err != nil {
		log.Fatal("Failed to open genome file: ", err)
	}

	r, err := genetics.NewGenomeReader(genomeFile, genetics.YAMLGenomeEncoding)
	if err != nil {
		log.Fatal("Failed to create genome reader: ", err)
	}

	startGenome, err := r.Read()
	if err != nil {
		log.Fatal("Failed to read start genome: ", err)
	}

	// Check if output dir exists
	if _, err := os.Stat(OUT_DIR); err == nil {
		// Backup it
		backUpDir := fmt.Sprintf("%s-%s", OUT_DIR, time.Now().Format("2006-01-02T15_04_05"))
		// Clear it
		err = os.Rename(OUT_DIR, backUpDir)
		if err != nil {
			log.Fatal("Failed to do previous results backup: ", err)
		}
	}

	// Create output dir
	err = os.MkdirAll(OUT_DIR, os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create output directory: ", err)
	}

	// Create experiment
	seed := time.Now().Unix()
	rand.Seed(seed)
	expt := experiment.Experiment{
		Id:       0,
		Trials:   make(experiment.Trials, neatOptions.NumRuns),
		RandSeed: seed,
	}

	var generationEvaluator experiment.GenerationEvaluator = evaluator.NewNeatPlayerGenerationEvaluator(OUT_DIR)

	// Run experiment in the separate goroutine
	errChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err = expt.Execute(neat.NewContext(ctx, neatOptions), startGenome, generationEvaluator, nil); err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Register handler to wait for termination signals
	go func(cancel context.CancelFunc) {
		fmt.Println("\nPress Ctrl+C to stop")

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		select {
		case <-signals:
			// signal to stop test fixture
			cancel()
		case err = <-errChan:
			// stop waiting
		}
	}(cancel)

	// Wait for experiment completion
	err = <-errChan
	if err != nil {
		// error during execution
		log.Fatalf("Experiment execution failed: %s", err)
	}

	// Print experiment results statistics
	expt.PrintStatistics()

	expResPath := fmt.Sprintf("%s/%s.dat", OUT_DIR, "results")
	if expResFile, err := os.Create(expResPath); err != nil {
		log.Fatal("Failed to create file for experiment results", err)
	} else if err = expt.Write(expResFile); err != nil {
		log.Fatal("Failed to save experiment results", err)
	}
}
