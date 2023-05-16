package d5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInstruction_createsAnInstructionStructureFromARawInstructionProvidedInInput(t *testing.T) {
	_, ok := parseInstruction("")
	assert.False(t, ok)

	inst, ok := parseInstruction("move 2 from 4 to 7")
	assert.True(t, ok)
	assert.Equal(t, 2, inst.times)
	assert.Equal(t, 4, inst.source)
	assert.Equal(t, 7, inst.target)

	inst, ok = parseInstruction("move 8 from 3 to 1")
	assert.True(t, ok)
	assert.Equal(t, 8, inst.times)
	assert.Equal(t, 3, inst.source)
	assert.Equal(t, 1, inst.target)
}

func Test_newCrateWithNameOnly_createsACrateWithNameButNilPrevious(t *testing.T) {
	expectedName := rune('A')
	c := newCrateWithNameOnly(expectedName)
	assert.Equal(t, expectedName, c.name)
	assert.Nil(t, c.prev)
}

func Test_newCrateWithNameAndPrevious_createsACrateWithNameAndPrevious(t *testing.T) {
	expectedName := rune('B')
	c1 := newCrateWithNameOnly(rune('A'))
	c2 := newCrateWithNameAndPrevious('B', c1)
	assert.Equal(t, expectedName, c2.name)
	assert.Equal(t, c1, c2.prev)
}

func Test_copy_createsANewCrateObjectWithTheSameValuesAsTheOriginalCrate(t *testing.T) {
	originalCrate := newCrateWithNameOnly('A')

	copyCrate := originalCrate.copy()

	assert.NotSame(t, originalCrate, copyCrate)
	assert.Equal(t, originalCrate.name, copyCrate.name)
	assert.Same(t, originalCrate.prev, copyCrate.prev)

	thePreviousCrate := newCrateWithNameOnly('P')
	originalCrate = newCrateWithNameAndPrevious('A', thePreviousCrate)

	copyCrate = originalCrate.copy()
	assert.NotSame(t, originalCrate, copyCrate)
	assert.Same(t, originalCrate.prev, copyCrate.prev)

	assert.Equal(t, originalCrate.name, copyCrate.name)
	assert.Equal(t, originalCrate.prev, copyCrate.prev)
}

func Test_newStack_createsANewCStackFromCratesProvidedAsStringAndAPosition(t *testing.T) {
	s1 := newStack("XYZ", 1)

	x := newCrateWithNameOnly('X')
	y := newCrateWithNameAndPrevious('Y', x)
	z := newCrateWithNameAndPrevious('Z', y)

	// Pointer variable equality is determined based on the equality of the referenced **values** (as opposed to the **memory addresses**)
	assert.Equal(t, z, s1.head)

	s2 := newStack("DEF", 2)

	d := newCrateWithNameOnly('D')
	// e will be missing
	f := newCrateWithNameAndPrevious('F', d)

	assert.NotEqual(t, f, s2.head)

	s3 := newStack("OQ", 3) // No P

	o := newCrateWithNameOnly('O')
	p := newCrateWithNameAndPrevious('P', o) // p should not be present in s3
	q := newCrateWithNameAndPrevious('Q', p)

	assert.NotEqual(t, q, s3.head)
}

func Test_add_addsACrateToAStack(t *testing.T) {
	s1 := newStack("ABC", 1)
	d := newCrateWithNameOnly('D')

	s1.add(d)

	assert.Equal(t, d, s1.head)
	assert.Same(t, d, s1.head)
}

func Test_move_movesACrateFromAStackToAnother(t *testing.T) {
	s1 := newStack("ABC", 1)
	c := s1.head
	b := c.prev

	s2 := newStack("D", 2)
	d := s2.head

	s1.move(s2)

	assert.Same(t, c, s2.head)
	assert.NotEqual(t, b, c.prev)
	assert.Same(t, d, c.prev)
}

func Test_moveTimes_movesCratesFromAStackToAnotherNTimesAlwaysMovingTheCrateOnTopOfTheSourceStack(t *testing.T) {
	s1 := newStack("ABC", 1)
	c := s1.head
	b := c.prev
	a := b.prev

	s2 := newStack("D", 2)
	d := s2.head

	s1.moveTimes(2, s2) //A, DCB

	assert.Same(t, a, s1.head)
	assert.Same(t, b, s2.head)
	assert.NotEqual(t, b, c.prev)
	assert.NotEqual(t, a, b.prev)
	assert.Same(t, d, c.prev)
	assert.Same(t, c, b.prev)
}

func Test_stackInPosition_returnsTheStackInTheRequiredPositionWithOkIdiom(t *testing.T) {
	s1 := newStack("ABC", 1)
	s2 := newStack("D", 2)
	s3 := newStack("EF", 3)

	c := &cargo{[]*stack{s1, s2, s3}}

	found, ok := c.stackInPosition(3)
	if !ok {
		assert.Fail(t, "could not find the required stack")
	}
	assert.Equal(t, s3, found)
}

func Test_solvePartOne(t *testing.T) {
	res, _ := solvePartOne("input")

	assert.Equal(t, "BSDMQFLSP", res)
}

// TODO : crear una funcion que "atrape" los errores
// cuando se devuelven 2 o más elementos
// siendo uno de ellos un error
