package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/akamensky/argparse"
)

// inspired by this TikTok: https://www.tiktok.com/foryou?_r=1&_t=8RwOVSKaLIU&feed_mode=v1&is_from_webapp=v1&item_id=7086552077370985770#/@mathletters/video/7086552077370985770
const radius = 1
const squareArea = 2 * (radius * 2)
const h = 1

type Result struct {
	numIterations int
	numHit        int
}

func withinCircle(x, y float64) bool {
	yMax := math.Abs(math.Sqrt(math.Pow(radius, 2) - math.Pow(x-h, 2)))
	return yMax >= y
}

func generateCoordinates() (float64, float64) {
	x := rand.Float64() * 2
	y := rand.Float64()
	return x, y
}

func estimatePi(numHit, totalNumTries int) float64 {
	return squareArea * (float64(numHit) / (float64(totalNumTries)))
}

func estimatePiCorrectness(estimatedPi float64) float64 {
	return 1 - math.Abs(math.Pi-estimatedPi)/math.Pi
}

func runMonteCarlo(maxIterations, numThreads int, verbose bool) *Result {

	results := make(chan Result)
	for thread := 0; thread < numThreads; thread++ {
		if verbose {
			fmt.Println("\t", "Starting Go Routine", thread+1)
		}
		go func() {
			numTriesInCircle := 0
			numIterations := maxIterations / numThreads
			for iter := 0; iter < numIterations; iter++ {
				x, y := generateCoordinates()
				if withinCircle(x, y) {
					numTriesInCircle++
				}
			}
			results <- Result{
				numIterations: numIterations,
				numHit:        numTriesInCircle,
			}
		}()
	}

	finalResult := &Result{
		numIterations: 0,
		numHit:        0,
	}
	for resultIter := 0; resultIter < numThreads; resultIter++ {
		result := <-results
		finalResult.numIterations += result.numIterations
		finalResult.numHit += result.numHit
		if verbose {
			fmt.Println("\t", "result from go routine", result)
		}

	}
	if verbose {
		fmt.Println("\t", "final result", finalResult)
	}

	return finalResult
}

func validatePositiveInteger(args []string) error {
	if len(args) < 1 {
		return errors.New("must provide an integer for maximum interations")
	}

	iterations := args[0]
	if n, err := strconv.ParseInt(iterations, 10, 64); err != nil || n < 0 {
		return errors.New("iterations must be a positive integer")
	}

	return nil
}

func main() {

	parser := argparse.NewParser("go-pi-estimate", "Estimates Pi through a Monte Carlo method")
	iterations := parser.Int("i", "iterations", &argparse.Options{
		Required: false,
		Validate: validatePositiveInteger,
		Help:     "the maximum number of iterations to use to estimate Pi per thread",
		Default:  10_000,
	})
	threadCount := parser.Int("t", "threads", &argparse.Options{
		Required: false,
		Validate: validatePositiveInteger,
		Help:     "the number of threads to use",
		Default:  4,
	})
	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "print verbose output during calculation", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	result := runMonteCarlo(*iterations, *threadCount, *verbose)

	piEstimate := estimatePi(result.numHit, result.numIterations)
	piEstimateCorrectness := estimatePiCorrectness(piEstimate)
	fmt.Println("Total Iterations", result.numIterations)
	fmt.Println("Pi Estimate:", piEstimate)
	fmt.Println("Correctness:", piEstimateCorrectness)

}
