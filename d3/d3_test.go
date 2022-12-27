package d3

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitRucksackInTwoCompartiments_takesAStringSplitsItInHalfAndReturnsTheTwoResultingStrings(t *testing.T) {
	c1, c2 := splitRucksackInTwoCompartiments("")
	assert.Equal(t, "", c1)
	assert.Equal(t, "", c2)

	// I'll suppose that input has len(input) pair

	c1, c2 = splitRucksackInTwoCompartiments("ab")
	assert.Equal(t, "a", c1)
	assert.Equal(t, "b", c2)

	c1, c2 = splitRucksackInTwoCompartiments("abcd")
	assert.Equal(t, "ab", c1)
	assert.Equal(t, "cd", c2)
}

func Test_findCommonLetterInTwoContainers_shouldReturnTheLetterPresentInBothCompartiments(t *testing.T) {
	c1 := ""
	c2 := ""
	assert.Equal(t, "", findCommonLetterInTwoContainers(c1, c2))

	c1 = "ab"
	c2 = "cd"
	assert.Equal(t, "", findCommonLetterInTwoContainers(c1, c2))

	c1 = "abcd"
	c2 = "efgd"
	assert.Equal(t, "d", findCommonLetterInTwoContainers(c1, c2))
}

func Test_solvePartOne(t *testing.T) {

	_, err := solvePartOne("nonExistingFile")
	assert.EqualError(t, errors.New("open nonExistingFile: no such file or directory"), err.Error())

	sum, err := solvePartOne("test_input1")
	assert.Nil(t, err)
	assert.Equal(t, 25, sum)

	sum, err = solvePartOne("test_input2")
	assert.Nil(t, err)
	assert.Equal(t, 26, sum)

	sum, err = solvePartOne("input")
	assert.Nil(t, err)
	assert.Equal(t, 8088, sum)
}

func Test_findCommonLetterAmongThreeElves_shouldReturnTheLetterPresentInTheRucksackOfThreeElves(t *testing.T) {

	e1 := ""
	e2 := ""
	e3 := ""
	c := findCommonLetterAmongThreeElves(e1, e2, e3)
	assert.Equal(t, "", c)

	e1 = "a"
	e2 = "b"
	e3 = "c"
	c = findCommonLetterAmongThreeElves(e1, e2, e3)
	assert.Equal(t, "", c)

	e1 = "abcd"
	e2 = "efg"
	e3 = "hijkl"
	c = findCommonLetterAmongThreeElves(e1, e2, e3)
	assert.Equal(t, "", c)

	e1 = "abcd"
	e2 = "ebgb"
	e3 = "hivbjbl"
	c = findCommonLetterAmongThreeElves(e1, e2, e3)
	assert.Equal(t, "b", c)

	e1 = "abcd"
	e2 = "efga"
	e3 = "hiakl"
	c = findCommonLetterAmongThreeElves(e1, e2, e3)
	assert.Equal(t, "a", c)

}

func Test_solvePartTwo(t *testing.T) {
	_, err := solvePartTwo("nonExistingFile")
	assert.EqualError(t, errors.New("open nonExistingFile: no such file or directory"), err.Error())

	sum, err := solvePartTwo("test_input3")
	assert.Nil(t, err)
	assert.Equal(t, 70, sum)

	sum, err = solvePartTwo("input")
	assert.Nil(t, err)
	assert.Equal(t, 2522, sum)
}
