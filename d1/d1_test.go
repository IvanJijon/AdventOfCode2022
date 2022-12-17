package aoc2022

import (
	"fmt"
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

func Test_AssignTotalCaloriesToElfs_shouldSetTheTotalCaloriesForEachElfAsDescribedInTheInputFile(t *testing.T) {
	elfs, err := AssignTotalCaloriesTo("test_input")
	assert.Nil(t, err)
	assert.Equal(t, 123, elfs[0].calories)
	assert.Equal(t, 579, elfs[1].calories)
}

func Test_FindElfCarryingTheMostCalories_returnsTheElfCarryingTheMostCalories(t *testing.T) {
	var elves []*Elf
	e1 := NewElf()
	e2 := NewElf()
	e3 := NewElf()
	e1.incrementCalories(1)
	e2.incrementCalories(3)
	e3.incrementCalories(2)
	elves = append(elves, e1)
	elves = append(elves, e2)
	elves = append(elves, e3)

	mostCaloricElf := FindElfCarryingTheMostCalories(elves)
	assert.Same(t, e2, mostCaloricElf)
}

func Test_Solve(t *testing.T) {
	c := SolvePartOne("input")
	fmt.Println(c)
}
