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

// To-do list:
// - Create a type round with oponentShape, myShape and outcome
// - Create a type outcome and constants win, draw,lose
// - Create a table with all the possible scenarios and outcomes, eg:
// paper, rock, win
// rock, paper, lose
