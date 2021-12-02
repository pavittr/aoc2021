package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pavittr/aoc2021/utils"
)

func main() {
	utils.RunDay(solution1, solution2)
}

func solution1(puzzleInput string) (string, error) {

	counter := 0
	splitInput := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	for i := 0; i < len(splitInput)-1; i++ {
		lower, _ := strconv.Atoi(splitInput[i])
		upper, _ := strconv.Atoi(splitInput[i+1])
		if upper > lower {
			counter = counter + 1
		}
	}

	return fmt.Sprintf("%d", counter), nil

}

func solution2(puzzleInput string) (string, error) {

	counter := 0
	splitInput := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	for i := 0; i < len(splitInput)-3; i++ {
		lower, _ := strconv.Atoi(splitInput[i])
		upper, _ := strconv.Atoi(splitInput[i+3])
		if upper > lower {
			counter = counter + 1
		}
	}

	return fmt.Sprintf("%d", counter), nil

}
