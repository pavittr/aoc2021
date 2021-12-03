package main

import (
	"fmt"

	"github.com/pavittr/aoc2021/utils"
	"strconv"
	"strings"
)

func main() {
	utils.RunDay(solution1, solution2)
}

func solution1(puzzleInput string) (string, error) {

	lines := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	gamma := make([]rune, len(lines[0]))
	epsilon := make([]rune, len(lines[0]))
	ones := make([]int, len(lines[0]))
	zeroes := make([]int, len(lines[0]))
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				ones[i] = ones[i] + 1
			} else {
				zeroes[i] = zeroes[i] + 1
			}
		}
	}
	for i := 0; i < len(ones); i++ {
		if ones[i] > zeroes[i] {
			gamma[i] = '1'
			epsilon[i] = '0'
		} else {
			gamma[i] = '0'
			epsilon[i] = '1'
		}
	}

	gammaInt, gErr := strconv.ParseInt(string(gamma), 2, 64)
	if gErr != nil {
		return "", fmt.Errorf("failed to parse gammd (%s) as binary", string(gamma))
	}
	epsilonInt, eErr := strconv.ParseInt(string(epsilon), 2, 64)
	if eErr != nil {
		return "", fmt.Errorf("failed to parse epsilon (%s) as binary", string(epsilon))
	}

	return fmt.Sprintf("%d", gammaInt*epsilonInt), nil
}

func solution2(puzzleInput string) (string, error) {
	lines := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	oxyInt, _ := findInt(0, '0', lines)
	scrubInt, _ := findInt(0, '1', lines)
	return fmt.Sprintf("%d", oxyInt*scrubInt), nil
}

func findInt(element int, priority rune, lines []string) (int64, error) {
	tmpLines := make([]string, 0)
	var zeroes, ones int
	for _, line := range lines {
		if line[element] == '1' {
			ones++
		} else {
			zeroes++
		}

	}

	for _, line := range lines {
		if (zeroes > ones && line[element] == byte(priority)) ||
			(zeroes <= ones && line[element] != byte(priority)) {
			tmpLines = append(tmpLines, line)
		}
	}

	if len(tmpLines) == 1 {
		retInt, rErr := strconv.ParseInt(tmpLines[0], 2, 64)
		if rErr != nil {
			return 0, fmt.Errorf("failed to parse last element (%s) as binary on element %d: %+v", tmpLines[0], element, rErr)
		}
		return retInt, nil
	}
	return findInt(element+1, priority, tmpLines)
}
