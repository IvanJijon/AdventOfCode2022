package d2

import "errors"

type outcome string

const (
	win  outcome = "win"
	draw outcome = "draw"
	lose outcome = "lose"
)

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
