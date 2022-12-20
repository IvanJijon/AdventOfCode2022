package d1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsEmptyLine_returnsTrueIfTheLineIsEmpty_FalseIfNot(t *testing.T) {
	emptyLine := ""
	b := IsEmptyLine(emptyLine)
	assert.Equal(t, true, b)

	lineWithContent := "this is not an empty line"
	b = IsEmptyLine(lineWithContent)
	assert.Equal(t, false, b)
}

func Test_AssignTotalCaloriesToElfs_shouldSetTheTotalCaloriesForEachElfAsDescribedInTheInputFile(t *testing.T) {
	elfs, err := AssignTotalCaloriesToElvesFrom("test_input")
	assert.Nil(t, err)
	assert.Equal(t, 123, elfs[0].calories)
	assert.Equal(t, 579, elfs[1].calories)
}

func Test_FindElfCarryingTheMostCalories_returnsTheElfCarryingTheMostCalories(t *testing.T) {
	var elves []*Elf
	e1 := NewElf()
	e2 := NewElf()
	e3 := NewElf()
	e1.IncrementCalories(1)
	e2.IncrementCalories(3)
	e3.IncrementCalories(2)
	elves = append(elves, e1)
	elves = append(elves, e2)
	elves = append(elves, e3)

	mostCaloricElf := FindElfCarryingTheMostCalories(elves)
	assert.Same(t, e2, mostCaloricElf)
}

func Test_SolvePartOne(t *testing.T) {
	solution := SolvePartOne("input")
	assert.Equal(t, 66616, solution)
}

func Test_RemoveElfFromSlice_ShouldRemoveASpecifiedElfFromSlice(t *testing.T) {
	// Each elf must have different calories for this to work
	e1 := NewElfWithCalories(1)
	e2 := NewElfWithCalories(2)
	e3 := NewElfWithCalories(3)
	e4 := NewElfWithCalories(4)

	var slice []*Elf
	slice = append(slice, e1, e2, e3)

	assert.Equal(t, []*Elf{e1, e2, e3}, slice)

	// We remove an elf in the middle
	slice = RemoveElfFromSlice(e2, slice)
	assert.Equal(t, []*Elf{e1, e3}, slice)

	// We remove an elf in the end
	slice = RemoveElfFromSlice(e1, slice)
	assert.Equal(t, []*Elf{e3}, slice)

	// We remove a non existing elf, the slice should remain the same
	slice = RemoveElfFromSlice(e4, slice)
	assert.Equal(t, []*Elf{e3}, slice)
}

func Test_SolvePartTwo(t *testing.T) {
	solution := SolvePartTwo("input")
	assert.Equal(t, 199172, solution)
}
