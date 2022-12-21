package d2

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_translateXYZtoShape_shouldReturnTheCorrespondingRockPaperScissorsShapeForXYZ(t *testing.T) {
	_, err := translateXYZtoShape("")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("shape not found"), err)
	}

	_, err = translateXYZtoShape("W")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("shape not found"), err)
	}

	r, err := translateXYZtoShape("X")
	assert.Nil(t, err)
	assert.Equal(t, r, rock)

	p, err := translateXYZtoShape("Y")
	assert.Nil(t, err)
	assert.Equal(t, p, paper)

	s, err := translateXYZtoShape("Z")
	assert.Nil(t, err)
	assert.Equal(t, s, scissors)
}

func Test_translateABCtoShape_shouldReturnTheCorrespondingRockPaperScissorsShapeForABC(t *testing.T) {
	_, err := translateXYZtoShape("")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("shape not found"), err)
	}

	_, err = translateABCtoShape("W")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("shape not found"), err)
	}

	r, err := translateABCtoShape("A")
	assert.Nil(t, err)
	assert.Equal(t, r, rock)

	p, err := translateABCtoShape("B")
	assert.Nil(t, err)
	assert.Equal(t, p, paper)

	s, err := translateABCtoShape("C")
	assert.Nil(t, err)
	assert.Equal(t, s, scissors)
}

func Test_scenarioScore(t *testing.T) {
	_, err := scenarioScore(shape(""), shape(""))
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	_, err = scenarioScore(rock, shape(""))
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	_, err = scenarioScore(shape(""), rock)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	score, err := scenarioScore(rock, rock)
	assert.Nil(t, err)
	assert.Equal(t, 4, score)

	score, err = scenarioScore(rock, scissors)
	assert.Nil(t, err)
	assert.Equal(t, 3, score)

	score, err = scenarioScore(paper, scissors)
	assert.Nil(t, err)
	assert.Equal(t, 9, score)
}

func Test_SolvePartOne(t *testing.T) {
	solution, err := SolvePartOne("input")
	assert.Nil(t, err)
	assert.Equal(t, 10718, solution)
}

func Test_translateXYZToDesiredOutcome_shouldReturnTheCorrespondingLoseDrawWinOutcomeForXYZ(t *testing.T) {

	_, err := translateXYZToDesiredOutcome("")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("input unknown"), err)
	}

	_, err = translateXYZToDesiredOutcome("W")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("input unknown"), err)
	}

	l, err := translateXYZToDesiredOutcome("X")
	assert.Nil(t, err)
	assert.Equal(t, l, lose)

	d, err := translateXYZToDesiredOutcome("Y")
	assert.Nil(t, err)
	assert.Equal(t, d, draw)

	w, err := translateXYZToDesiredOutcome("Z")
	assert.Nil(t, err)
	assert.Equal(t, w, win)
}

func Test_shapeIMustUse_tellsMeTheShapeIMustUseToObtainTheDesiredOutcomeBasedOnTheOponentsMove(t *testing.T) {

	_, err := findShapeIMustUse(shape(""), outcome(""))
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	_, err = findShapeIMustUse(rock, outcome(""))
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	_, err = findShapeIMustUse(shape(""), lose)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	s, _ := findShapeIMustUse(rock, win)
	assert.Equal(t, paper, s)

	s, _ = findShapeIMustUse(scissors, draw)
	assert.Equal(t, scissors, s)

	s, _ = findShapeIMustUse(paper, lose)
	assert.Equal(t, rock, s)
}

func Test_SolvePartTwo(t *testing.T) {
	solution, err := SolvePartTwo("input")
	assert.Nil(t, err)
	assert.Equal(t, 14652, solution)
}
