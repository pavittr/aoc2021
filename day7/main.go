package main

import (
	"fmt"
	"sort"

	"log"

	"strings"

	"strconv"

	"github.com/pavittr/aoc2021/utils"
)

func main() {
	if err := utils.RunDay(solution1, solution2); err != nil {
		log.Fatalf("Failed to run day: %+v\n", err)
	}
}

func solution1(puzzleInput string) (string, error) {
	input := strings.Split(strings.TrimSpace(puzzleInput), ",")
	crabs := make([]int, 0)
	for _, crabStr := range input {
		crab, cErr := strconv.Atoi(crabStr)
		if cErr != nil {
			return "", fmt.Errorf("failed to convert %s to int: %+v", crabStr, cErr)
		}
		crabs = append(crabs, crab)
	}

	sort.Ints(crabs)
	midpoint := int(len(crabs)/2) - 1
	
	baseLine := crabs[midpoint]
	distance := 0
	for _, crab := range crabs {
		if crab < baseLine {
			distance += (baseLine - crab)
		} else {
			distance += (crab - baseLine)
		}

	}

	return fmt.Sprintf("%d", distance), nil
}

func solution2(puzzleInput string) (string, error) {
	input := strings.Split(strings.TrimSpace(puzzleInput), ",")
	crabs := make([]int, 0)

	for _, crabStr := range input {
		crab, cErr := strconv.Atoi(crabStr)
		if cErr != nil {
			return "", fmt.Errorf("failed to convert %s to int: %+v", crabStr, cErr)
		}
		crabs = append(crabs, crab)
	}

	sort.Ints(crabs)
	smallest, biggest := crabs[0], crabs[len(crabs)-1]
	alignment := make([][]int, biggest+1)
	for _, crab := range crabs {
		alignment[crab] = append(alignment[crab], 1)

	}

	cost := 0

	for smallest != biggest {
		lowerCrabs := alignment[smallest]
		lowerCost := 0
		for _, lowerCrab := range lowerCrabs {
			lowerCost += lowerCrab 
		}
		higherCrabs := alignment[biggest]
		higherCost := 0
		for _, higherCrab := range higherCrabs {
			higherCost += higherCrab
		}

		if lowerCost <= higherCost {
			// move lowers up
			for _, lowerCrab := range lowerCrabs {
				cost += lowerCrab
				alignment[smallest+1] = append(alignment[smallest+1], lowerCrab+1)
			}
			smallest++
		} else {
			// move uppers down
			for _, higherCrab := range higherCrabs {
				cost += higherCrab
				alignment[biggest-1] = append(alignment[biggest-1], higherCrab+1)
			}
			biggest--
		}
	}


	return fmt.Sprintf("%d", cost), nil

}
