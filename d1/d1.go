package aoc2022

import (
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
		elf.incrementCalories(d)
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
