package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type element struct {
	something int
}

func NewElementWithSomething(i int) *element {
	return &element{i}
}

func Test_RemoveElementFromSlice(t *testing.T) {

	// Each element must have different calories for this to work
	e1 := NewElementWithSomething(1)
	e2 := NewElementWithSomething(2)
	e3 := NewElementWithSomething(3)
	e4 := NewElementWithSomething(4)
	e5 := NewElementWithSomething(5)

	var slice []*element
	slice = append(slice, e1, e2, e3, e4)

	assert.Equal(t, []*element{e1, e2, e3, e4}, slice)    // order matters
	assert.NotEqual(t, []*element{e1, e2, e3, e5}, slice) // e5 in stead of e4
	assert.ElementsMatch(t, []*element{e1, e2, e3, e4}, slice)

	// we remove the first element
	slice = RemoveElementFromSlice(e1, slice)
	assert.Equal(t, []*element{e2, e3, e4}, slice)

	// we remove an element in the middle
	slice = RemoveElementFromSlice(e3, slice)
	assert.Equal(t, []*element{e2, e4}, slice)

	// we remove the last element
	slice = RemoveElementFromSlice(e4, slice)
	assert.ElementsMatch(t, []*element{e2}, slice)

	// we remove an unexisting element
	slice = RemoveElementFromSlice(e5, slice)
	assert.ElementsMatch(t, []*element{e2}, slice)
}
