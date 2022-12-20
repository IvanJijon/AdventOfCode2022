package d1

import (
	"aoc2022/utils"
	"bufio"
	"os"
	"strconv"
)

func IsEmptyLine(l string) bool {
	return !(len(l) > 0)
}

func AssignTotalCaloriesToElvesFrom(inputFile string) ([]*Elf, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var elves []*Elf
	elves = AssignTotalCalories(elves, NewElf(), scanner)

	return elves, scanner.Err()
}

func AssignTotalCalories(elves []*Elf, elf *Elf, scanner *bufio.Scanner) []*Elf {
	if !scanner.Scan() {
		return elves
	}

	if !IsEmptyLine(scanner.Text()) {
		d, _ := strconv.Atoi(scanner.Text())
		elf.IncrementCalories(d)
		return AssignTotalCalories(elves, elf, scanner)
	}

	elves = append(elves, elf)
	elf = NewElf()

	return AssignTotalCalories(elves, elf, scanner)
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

func SolvePartOne(inputFile string) int {
	elves, _ := AssignTotalCaloriesToElvesFrom(inputFile)
	return FindElfCarryingTheMostCalories(elves).calories
}

func RemoveElfFromSlice(e *Elf, slice []*Elf) []*Elf {
	return utils.RemoveElementFromSlice(e, slice)
}

func SolvePartTwo(inputFile string) int {
	elves, _ := AssignTotalCaloriesToElvesFrom(inputFile)

	e1 := FindElfCarryingTheMostCalories(elves)

	elves = RemoveElfFromSlice(e1, elves)
	e2 := FindElfCarryingTheMostCalories(elves)

	elves = RemoveElfFromSlice(e2, elves)
	e3 := FindElfCarryingTheMostCalories(elves)

	return e1.calories + e2.calories + e3.calories
}
