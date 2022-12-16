package aoc2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewElf_aNewlyCreatedElfShouldHaveZeroCalories(t *testing.T) {
	e := NewElf()
	assert.Equal(t, 0, e.calories, "A newly created Elf should have zero calories")
}

func Test_incrementCalories_shouldIncrementTheCaloriesAnElfIsCarrying(t *testing.T) {
	e := NewElf()
	assert.Equal(t, 0, e.calories)
	e.incrementCalories(1)
	assert.Equal(t, 1, e.calories)
}

func Test_IsEmptyLine_returnsTrueIfTheLineIsEmpty_FalseIfNot(t *testing.T) {
	emptyLine := ""
	b := IsEmptyLine(emptyLine)
	assert.Equal(t, true, b)

	lineWithContent := "this is not an empty line"
	b = IsEmptyLine(lineWithContent)
	assert.Equal(t, false, b)
}
