package d2

import "errors"

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
