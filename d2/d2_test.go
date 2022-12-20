package d2

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_translateXYZtoShape_shouldReturnTheCorrespondingRockPaperScissorsShapeForXYZ(t *testing.T) {
	_, err := shape.translateXYZtoShape("")
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
