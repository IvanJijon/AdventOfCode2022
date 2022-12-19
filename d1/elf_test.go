package aoc2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewElf_aNewlyCreatedElfShouldHaveZeroCalories(t *testing.T) {
	e := NewElf()
	assert.Equal(t, 0, e.calories, "A newly created Elf should have zero calories")
}

func Test_NewElfWithCalories_anElfCreatedInThisWayShouldHaveTheNumberOfCaloriesPassedInArgument(t *testing.T) {
	e1 := NewElfWithCalories(42)
	assert.Equal(t, 42, e1.calories)
	assert.NotEqual(t, 40, e1.calories)

	e2 := NewElfWithCalories(1337)
	assert.Equal(t, 1337, e2.calories)
	assert.NotEqual(t, 77, e2.calories)

}
func Test_incrementCalories_shouldIncrementTheCaloriesAnElfIsCarrying(t *testing.T) {
	e := NewElf()
	assert.Equal(t, 0, e.calories)
	e.IncrementCalories(1)
	assert.Equal(t, 1, e.calories)
}

func Test_SomethingInterestingAboutObjectsInGolang(t *testing.T) {
	assert.Equal(t, NewElf(), NewElf())
	assert.Equal(t, NewElfWithCalories(0), NewElf())
	assert.NotEqual(t, NewElfWithCalories(0), NewElfWithCalories(3))
}
