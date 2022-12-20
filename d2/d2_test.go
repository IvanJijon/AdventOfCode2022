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

func Test_findOutcome(t *testing.T) {
	_, err := findOutcome(shape(""), shape(""))
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	_, err = findOutcome(rock, shape(""))
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	_, err = findOutcome(shape(""), rock)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("scenario not found"), err)
	}

	o, err := findOutcome(rock, rock)
	assert.Nil(t, err)
	assert.Equal(t, draw, o)

	o, err = findOutcome(rock, scissors)
	assert.Nil(t, err)
	assert.Equal(t, lose, o)

	o, err = findOutcome(paper, scissors)
	assert.Nil(t, err)
	assert.Equal(t, win, o)
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
