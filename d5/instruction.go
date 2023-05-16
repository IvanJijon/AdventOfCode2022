package d5

import (
	"strconv"
	"strings"
)

type instruction struct {
	times  int
	source int
	target int
}

func parseInstruction(raw string) (*instruction, bool) {
	if raw == "" {
		return nil, false
	}

	pieces := strings.Split(raw, " ")

	times, _ := strconv.Atoi(pieces[1])
	source, _ := strconv.Atoi(pieces[3])
	target, _ := strconv.Atoi(pieces[5])

	return &instruction{times, source, target}, true
}
