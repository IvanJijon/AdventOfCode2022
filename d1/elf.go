package aoc2022

type Elf struct {
	calories int
}

func NewElf() *Elf {
	return &Elf{0}
}

func (e *Elf) incrementCalories(c int) {
	e.calories += c
}