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

func Test_findCommonLetter_shouldReturnTheLetterPresentInBothCompartiments(t *testing.T) {
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
