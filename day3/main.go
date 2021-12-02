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

	xPos := 0
	zPos := 0
	splitInput := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	for i := 0; i < len(splitInput); i++ {
		input := strings.Split(splitInput[i], " ")
		value, _ := strconv.Atoi(input[1])
		switch input[0] {
		case "forward": xPos = xPos + value
		case "up" : zPos = zPos - value
		case "down" : zPos = zPos + value
		}
		
	}

	return fmt.Sprintf("%d", zPos * xPos), nil

}

func solution2(puzzleInput string) (string, error) {

	xPos := 0
	zPos := 0
	aim := 0
	splitInput := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	for i := 0; i < len(splitInput); i++ {
		input := strings.Split(splitInput[i], " ")
		value, _ := strconv.Atoi(input[1])
		switch input[0] {
		case "forward":
			xPos = xPos + value
			zPos = zPos + (aim * value)
		case "up" : aim = aim - value
		case "down" : aim = aim + value
		}
	}

	return fmt.Sprintf("%d", zPos * xPos), nil

}

