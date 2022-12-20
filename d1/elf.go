package d1

type Elf struct {
	calories int
}

func NewElf() *Elf {
	return &Elf{0}
}

func NewElfWithCalories(c int) *Elf {
	e := NewElf()
	e.IncrementCalories(c)
	return e
}

func (e *Elf) IncrementCalories(c int) {
	e.calories += c
}
