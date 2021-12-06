package main

import (
	"fmt"

	"log"

	"strings"

	"strconv"

	"github.com/pavittr/aoc2021/utils"
)

type Fish int64

func (f Fish) proceed() []Fish {
	if f == 0 {
		return []Fish{6,8}
	}
	return []Fish{f -1}
}
func main() {
	if err := utils.RunDay(solution1, solution2); err != nil {
		log.Fatalf("Failed to run day: %+v\n", err)
	}
}

func solution1(puzzleInput string) (string, error) {
	txtLines := strings.Split(strings.TrimSpace(puzzleInput), ",")
	shoal := make([]Fish, 0)
	for _, element := range txtLines {
		r, fErr := strconv.Atoi(element)
		if fErr != nil {
			return "", fmt.Errorf("failed to turn %s into an int: %+v", element, fErr)
		}
		
		shoal = append(shoal, Fish(r))
	}
	nextShoal := make([]Fish,0)
	for i:= 0 ;i< 256 ; i++ {
		for _, fish := range shoal {
			nextShoal = append(nextShoal, fish.proceed()...)
		}
		shoal = nextShoal
		nextShoal = make([]Fish, 0)
	}
	log.Printf("shial %d\n", len(shoal))

	return fmt.Sprintf(""), nil
}

func solution2(puzzleInput string) (string, error) {

	return fmt.Sprintf(""), nil
}
