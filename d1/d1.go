package aoc2022

import (
	"bufio"
	"os"
	"strconv"
)

type Elf struct {
	calories int
}

func NewElf() *Elf {
	return &Elf{0}
}

func (e *Elf) incrementCalories(c int) {
	e.calories += c
}

func IsEmptyLine(l string) bool {
	return !(len(l) > 0)
}

// Sería ideal hacerlo recursivo
func AssignTotalCaloriesToElfs(inputFile string) ([]*Elf, error) {
	elfs := InitElfsWithZeroCaloriesFrom(inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	currentElf := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if !IsEmptyLine(scanner.Text()) {
			d, _ := strconv.Atoi(scanner.Text())
			elfs[currentElf].incrementCalories(d)
			continue
		}
		currentElf++
	}

	return elfs, scanner.Err()
}

func InitElfsWithZeroCaloriesFrom(inputFile string) []*Elf {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var elves []*Elf
	for scanner.Scan() {
		if IsEmptyLine(scanner.Text()) {
			elves = append(elves, NewElf())
		}
	}

	return elves
}

func FindElfCarryingTheMostCalories(elves []*Elf) *Elf {
	max := elves[0]
	for _, e := range elves {
		if e.calories > max.calories {
			max = e
		}
	}
	return max
}

func Solve(inputFile string) int {
	elves, _ := AssignTotalCaloriesToElfs(inputFile)
	return FindElfCarryingTheMostCalories(elves).calories
}
