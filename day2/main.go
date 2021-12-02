package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")
	testData, testErr := ioutil.ReadFile("test.data")
	if testErr != nil {
		log.Fatalf("Failed to read test data: %+v\n", testErr)
	}

	realData, realErr := ioutil.ReadFile("real.data")
	if realErr != nil {
		log.Fatalf("Failed to read real data: %+v\n", realErr)
	}

	p1TestAnswer, p1TestSolErr := solution1(string(testData))
	if p1TestSolErr != nil {
		log.Fatalf("Failed to run Part 1 test solution: %+v\n", p1TestSolErr)

	}
	log.Printf("Part 1: test answer is %s\n", p1TestAnswer)

	p1RealAnswer, p1RealSolErr := solution1(string(realData))
	if p1RealSolErr != nil {
		log.Fatalf("Failed to run p1Real solution: %+v\n", p1RealSolErr)

	}

	log.Printf("Part1: real Answer is %s\n", p1RealAnswer)

	p2TestAnswer, p2TestSolErr := solution2(string(testData))
	if p2TestSolErr != nil {
		log.Fatalf("Failed to run Part 2 test solution: %+v\n", p2TestSolErr)

	}
	log.Printf("Part 2: test answer is %s\n", p2TestAnswer)

	p2RealAnswer, p2RealSolErr := solution2(string(realData))
	if p2RealSolErr != nil {
		log.Fatalf("Failed to run p2Real solution: %+v\n", p2RealSolErr)

	}

	log.Printf("Part2: real Answer is %s\n", p2RealAnswer)

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
