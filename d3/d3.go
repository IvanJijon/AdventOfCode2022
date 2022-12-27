package d3

import (
	"bufio"
	"os"
	"strings"
)

// How to solve Part One:
// For each line:
// 	I need to split it in half
// 	Each half will represent a "compartiment"
// 	I need to find the letter present in both compartiments

// I need a letter -> priority conversion table
// a -> 1
// b -> 2
// ...

// Then I need to sum all the priority values corresponding
var stringToPriority = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func splitRucksackInTwoCompartiments(s string) (c1 string, c2 string) {
	if s == "" {
		return "", ""
	}

	c1 = s[:len(s)/2]
	c2 = s[len(s)/2:]

	return c1, c2
}

func findCommonLetterInTwoContainers(c1, c2 string) string {
	for _, v := range c1 {
		if strings.Contains(c2, string(v)) {
			return string(v)
		}
	}
	return ""
}

func solvePartOne(inputFile string) (int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		sum += stringToPriority[findCommonLetterInTwoContainers(splitRucksackInTwoCompartiments(scanner.Text()))]
	}
	return sum, nil
}

// How to solve Part Two:
// - Read the rucksack content for each group of 3 elves
// - Among the 3 elves of a group find the letter in common
//   they carry in their rucksacks
// - Using stringToPriority keep the sum of the priorities
//   found for each group identifier

// I would like to rewrite this function in a recursive way
func findCommonLetterAmongThreeElves(e1, e2, e3 string) string {
	for _, c := range e1 {
		if strings.Contains(e2, string(c)) && strings.Contains(e3, string(c)) {
			return string(c)
		}
	}
	return ""
}

func solvePartTwo(inputFile string) (int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		e1 := scanner.Text()
		scanner.Scan()
		e2 := scanner.Text()
		scanner.Scan()
		e3 := scanner.Text()

		sum += stringToPriority[findCommonLetterAmongThreeElves(e1, e2, e3)]

	}
	return sum, nil

}
