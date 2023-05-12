package d4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newSectionAssignmentPair_returnsAnErrorWhenTheProvidedPairOfSectionsIsEmpty(t *testing.T) {
	_, err := newSectionAssignmentPair("")
	assert.Error(t, err, "sections cannot be empty")
}

func Test_newSectionAssignmentPair_createsASectionAssignmentWithTheProvidedPairOfSections(t *testing.T) {
	sap, err := newSectionAssignmentPair("2-4,6-8")
	assert.Nil(t, err)
	assert.Equal(t, [2]section{newSection("2-4"), newSection("6-8")}, sap.sections)

	sap, err = newSectionAssignmentPair("5-7,7-9")
	assert.Nil(t, err)
	assert.Equal(t, [2]section{newSection("5-7"), newSection("7-9")}, sap.sections)
}

func Test_newSection_createsANewSectionFromTheProvidedString(t *testing.T) {
	expected := section{1, 2}
	actual := newSection("1-2")

	assert.Equal(t, expected, actual)
	assert.Equal(t, 1, actual.lowerBound)
	assert.Equal(t, 2, actual.upperBound)
}

func Test_assignSectionsTo_assignsTheCorrespondingSectionToEachElf(t *testing.T) {
	sap, _ := newSectionAssignmentPair("2-3,4-5")
	e1 := &elf{}
	e2 := &elf{}

	sap.assignSectionsTo(e1, e2)

	assert.Equal(t, newSection("2-3"), e1.s)
	assert.Equal(t, newSection("4-5"), e2.s)
}

func Test_isContainedIn_returnsTrueIfSection1IsContainedBySection2FalseIfNot(t *testing.T) {
	s1 := newSection("1-1") //1.........
	s2 := newSection("5-5") //....5.....
	assert.False(t, s1.isContainedBy(s2))

	s1 = newSection("5-6") //....56....
	s2 = newSection("7-8") //......78..
	assert.False(t, s1.isContainedBy(s2))

	s1 = newSection("4-8") //...45678..
	s2 = newSection("5-7") //....567...
	assert.False(t, s1.isContainedBy(s2))

	s1 = newSection("5-7") //....567...
	s2 = newSection("5-7") //....567...
	assert.True(t, s1.isContainedBy(s2))

	s1 = newSection("5-7") //....567...
	s2 = newSection("4-8") //...45678..
	assert.True(t, s1.isContainedBy(s2))
}

func Test_containsOrIsContainedBy_returnsTrueIfASectionIsContainedByAnotherFalseIfNot(t *testing.T) {
	s1 := newSection("1-1") //1.........
	s2 := newSection("5-5") //....5.....
	assert.False(t, s1.containsOrIsContainedBy(s2))

	s1 = newSection("5-6") //....56....
	s2 = newSection("7-8") //......78..
	assert.False(t, s1.containsOrIsContainedBy(s2))

	s1 = newSection("4-8") //...45678..
	s2 = newSection("5-7") //....567...
	assert.True(t, s1.containsOrIsContainedBy(s2))

	s1 = newSection("5-7") //....567...
	s2 = newSection("5-7") //....567...
	assert.True(t, s1.containsOrIsContainedBy(s2))

	s1 = newSection("5-7") //....567...
	s2 = newSection("4-8") //...45678..
	assert.True(t, s1.containsOrIsContainedBy(s2))
}

func Test_solvePartOne(t *testing.T) {
	r, _ := solvePartOne("test_input1")
	assert.Equal(t, 2, r)

	r, _ = solvePartOne("input")
	assert.Equal(t, 477, r)
}
