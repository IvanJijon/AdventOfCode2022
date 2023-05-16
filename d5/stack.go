package d5

import "fmt"

type crate struct {
	name rune
	prev *crate
}

func newCrateWithNameOnly(name rune) *crate {
	return newCrateWithNameAndPrevious(name, nil)
}

func newCrateWithNameAndPrevious(name rune, previous *crate) *crate {
	return &crate{name: name, prev: previous}
}

func (c *crate) copy() *crate {
	return newCrateWithNameAndPrevious(c.name, c.prev)
}

type stack struct {
	position int
	head     *crate
}

func newStack(crates string, position int) *stack {
	s := &stack{position: position}
	for _, c := range crates {
		s.add(newCrateWithNameOnly(rune(c)))
	}
	return s
}

func (s *stack) displayStack() {
	fmt.Println("Stack:", s.position)
	c := s.head
	for c != nil {
		fmt.Println(c.name)
		c = c.prev
	}
	fmt.Println("-------------")
}

func (s *stack) add(new *crate) {
	new.prev = s.head
	s.head = new
}

func (source *stack) move(target *stack) {
	pivot := source.head.copy()
	target.add(source.head)
	source.head = pivot.prev
}

func (source *stack) moveTimes(times int, target *stack) {
	for i := 0; i < times; i++ {
		source.move(target)
	}
}

type cargo struct {
	stacks []*stack
}

func newCargo(ss []*stack) *cargo {
	return &cargo{ss}
}

func (c *cargo) stackInPosition(pos int) (*stack, bool) {
	for _, s := range c.stacks {
		if s.position == pos {
			return s, true
		}
	}

	return nil, false
}
