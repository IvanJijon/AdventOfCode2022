package d2

import "errors"

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
}

var scenariosTable = []*round{
	{rock, rock, draw},
	{rock, paper, win},
	{rock, scissors, lose},
	{paper, rock, lose},
	{paper, paper, draw},
	{paper, scissors, win},
	{scissors, rock, win},
	{scissors, paper, lose},
	{scissors, scissors, draw},
}

func findOutcome(om shape, mm shape) (outcome, error) {

	for _, scenario := range scenariosTable {
		if scenario.oponentMove == om && scenario.myMove == mm {
			return scenario.outcome, nil
		}
	}
	return outcome(""), errors.New("scenario not found")
}

// To-do list:
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
//  based on both my oponent's shape and mine
// - Create a function that calculates the score of a round
// - Find a way to keep track of the score
