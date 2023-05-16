package d5

import (
	"bufio"
	"os"
)

// How to solve Part One:
// - Create a stack structure (FILO)
// - Initialize each stack with the given crates
//
// - Create the logic that will be able to interpret the instructions of the rearrangement procedure
// - Create the logic that will execute the given instruction in order following the rearrangement procedure

func initializeStacks() []*stack {
	s1 := newStack("DBJV", 1)
	s2 := newStack("PVBWRDF", 2)
	s3 := newStack("RGFLDCWQ", 3)
	s4 := newStack("WJPMLNDB", 4)
	s5 := newStack("HNBPCSQ", 5)
	s6 := newStack("RDBSNG", 6)
	s7 := newStack("ZBPMQFSH", 7)
	s8 := newStack("WLF", 8)
	s9 := newStack("SVFMR", 9)

	var r []*stack
	r = append(r, s1, s2, s3, s4, s5, s6, s7, s8, s9)

	return r
}

func solvePartOne(inputFile string) (string, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	c := newCargo(initializeStacks())
	for scanner.Scan() {
		instruc, _ := parseInstruction(scanner.Text())
		src, _ := c.stackInPosition(instruc.source)
		tgt, _ := c.stackInPosition(instruc.target)
		src.moveTimes(instruc.times, tgt)
	}

	result := ""
	for _, s := range c.stacks {
		result = result + string(s.head.name)
	}
	return result, nil
}
