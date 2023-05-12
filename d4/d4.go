package d4

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// How to solve Part One:
// For each line (pair of Elves):
// 	- Retrieve the section assigned to each Elf
// 	- Check if one sectionis contained by the other
//  - If that's the case, increment the sum
//  - Return the sum

type section struct {
	lowerBound int
	upperBound int
}

func newSection(s string) section {
	bounds := strings.Split(s, "-")
	lb, _ := strconv.Atoi(bounds[0])
	ub, _ := strconv.Atoi(bounds[1])
	return section{
		lowerBound: lb,
		upperBound: ub,
	}
}

func (s1 section) isContainedBy(s2 section) bool {
	return s1.lowerBound >= s2.lowerBound && s1.upperBound <= s2.upperBound
}

func (s1 section) containsOrIsContainedBy(s2 section) bool {
	return s1.isContainedBy(s2) || s2.isContainedBy(s1)
}

func (s1 section) overlapsWith(s2 section) bool {
	cond1 := s1.upperBound >= s2.lowerBound && s1.upperBound <= s2.upperBound
	cond2 := s1.lowerBound >= s2.lowerBound && s1.lowerBound <= s2.upperBound
	return cond1 || cond2 || s1.containsOrIsContainedBy(s2)
}

type sectionAssignmentPair struct {
	sections [2]section
}

func newSectionAssignmentPair(s string) (*sectionAssignmentPair, error) {
	if s == "" {
		return nil, errors.New("sections cannot be empty")
	}
	ss := strings.Split(s, ",")

	return &sectionAssignmentPair{[2]section{newSection(ss[0]), newSection(ss[1])}}, nil
}

func (sap *sectionAssignmentPair) assignSectionsTo(e1, e2 *elf) {
	e1.s = sap.sections[0]
	e2.s = sap.sections[1]
}

type elf struct {
	s section
}

func solvePartOne(inputFile string) (int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		sap, _ := newSectionAssignmentPair(scanner.Text())
		e1, e2 := &elf{}, &elf{}
		sap.assignSectionsTo(e1, e2)
		if e1.s.containsOrIsContainedBy(e2.s) {
			sum += 1
		}
	}
	return sum, nil
}

func solvePartTwo(inputFile string) (int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		sap, _ := newSectionAssignmentPair(scanner.Text())
		e1, e2 := &elf{}, &elf{}
		sap.assignSectionsTo(e1, e2)
		if e1.s.overlapsWith(e2.s) {
			sum += 1
		}
	}
	return sum, nil
}
