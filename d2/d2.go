package d2

import (
	"bufio"
	"errors"
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

type shape string

const (
	rock     shape = "rock"
	paper    shape = "paper"
	scissors shape = "scissors"
)

// there is some duplication in those two translating functions
// but I'll intentionally leave them as they are since it's more explicit
func translateXYZtoShape(s string) (shape, error) {
	if s == "X" {
		return rock, nil
	}

	if s == "Y" {
		return paper, nil
	}

	if s == "Z" {
		return scissors, nil
	}

	return "", errors.New("shape not found")
}

func translateABCtoShape(s string) (shape, error) {
	if s == "A" {
		return rock, nil
	}

	if s == "B" {
		return paper, nil
	}

	if s == "C" {
		return scissors, nil
	}

	return "", errors.New("shape not found")
}

type outcome string

const (
	win  outcome = "win"
	draw outcome = "draw"
	lose outcome = "lose"
)

type round struct {
	oponentMove shape
	myMove      shape
	outcome     outcome
	score       int
}

var scenariosTable = []*round{
	{rock, rock, draw, 4},
	{rock, paper, win, 8},
	{rock, scissors, lose, 3},
	{paper, rock, lose, 1},
	{paper, paper, draw, 5},
	{paper, scissors, win, 9},
	{scissors, rock, win, 7},
	{scissors, paper, lose, 2},
	{scissors, scissors, draw, 6},
}

func scenarioScore(om shape, mm shape) (int, error) {
	for _, scenario := range scenariosTable {
		if scenario.oponentMove == om && scenario.myMove == mm {
			return scenario.score, nil
		}
	}
	return -1, errors.New("scenario not found")
}

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

func translateXYZToDesiredOutcome(s string) (outcome, error) {
	if s == "X" {
		return lose, nil
	}

	if s == "Y" {
		return draw, nil
	}

	if s == "Z" {
		return win, nil
	}

	return "", errors.New("input unknown")
}

func findShapeIMustUse(om shape, desiredOutcome outcome) (shape, error) {
	for _, scenario := range scenariosTable {
		if scenario.oponentMove == om && scenario.outcome == desiredOutcome {
			return scenario.myMove, nil
		}
	}
	return shape(""), errors.New("scenario not found")
}

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
