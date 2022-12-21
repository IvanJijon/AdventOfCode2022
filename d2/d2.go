package d2

import (
	"bufio"
	"os"
	"strings"
)

// Part One To-do list:
// - Create a type shape with constants rock, paper, scissors
// - Create translating functions XYZ or ABC -> rock, paper, scissors
// - Create a type outcome and constants win, draw,lose
// - Create an structure round with oponentMove, myMove and outcome
// - Create a table with all the possible scenarios and outcomes, eg:
// oponent's shape	|  my shape | outcome
// -----------------------------------------
// 		paper		|   rock 	| 	lose
// 		rock		|   paper 	| 	win
// 		scissors	|  scissors	| 	draw
// ...
// - Create a way to read the outcome of a round
//  based on both my oponent's move and mine
// - Create a way to retrieve the score I made in a round
//  based on both my oponent's move and mine
// - Find a way to keep track of the score

func SolvePartOne(inputFile string) (int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	score := 0
	for scanner.Scan() {
		moves := strings.Fields(scanner.Text())
		om, err := translateABCtoShape(moves[0])
		if err != nil {
			return -1, err
		}
		mm, err := translateXYZtoShape(moves[1])
		if err != nil {
			return -1, err
		}
		s, err := scenarioScore(om, mm)
		if err != nil {
			return -1, err
		}
		score += s
	}

	return score, err
}

// Part One To-do list:
// - Create a function that translates XYZ to the desired outcome
// - Create a function that, based on the opponent's move and what outcome I am supposed to get (win, draw or lose)
// it will tell me which shape should I use for the round

func SolvePartTwo(inputFile string) (int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	score := 0
	for scanner.Scan() {
		moves := strings.Fields(scanner.Text())
		om, err := translateABCtoShape(moves[0])
		if err != nil {
			return -1, err
		}
		do, err := translateXYZToDesiredOutcome(moves[1])
		if err != nil {
			return -1, err
		}
		mm, err := findShapeIMustUse(om, do)
		if err != nil {
			return -1, err
		}
		s, err := scenarioScore(om, mm)
		if err != nil {
			return -1, err
		}
		score += s
	}

	return score, err
}
