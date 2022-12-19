package utils

import (
	aoc2022 "aoc2022/d1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveElementFromSlice(t *testing.T) {

	// Each elf must have different calories for this to work
	e1 := aoc2022.NewElfWithCalories(1)
	e2 := aoc2022.NewElfWithCalories(2)
	e3 := aoc2022.NewElfWithCalories(3)
	e4 := aoc2022.NewElfWithCalories(4)
	e5 := aoc2022.NewElfWithCalories(5)

	var slice []*aoc2022.Elf
	slice = append(slice, e1, e2, e3, e4)

	assert.Equal(t, []*aoc2022.Elf{e1, e2, e3, e4}, slice)    // order matters
	assert.NotEqual(t, []*aoc2022.Elf{e1, e2, e3, e5}, slice) // e5 in stead of e4
	assert.ElementsMatch(t, []*aoc2022.Elf{e1, e2, e3, e4}, slice)

	// we remove the first element
	slice = RemoveElementFromSlice(e1, slice)
	assert.Equal(t, []*aoc2022.Elf{e2, e3, e4}, slice)

	// we remove an element in the middle
	slice = RemoveElementFromSlice(e3, slice)
	assert.Equal(t, []*aoc2022.Elf{e2, e4}, slice)

	// we remove the last element
	slice = RemoveElementFromSlice(e4, slice)
	assert.ElementsMatch(t, []*aoc2022.Elf{e2}, slice)

	// we remove an unexisting element
	slice = RemoveElementFromSlice(e5, slice)
	assert.ElementsMatch(t, []*aoc2022.Elf{e2}, slice)
}
