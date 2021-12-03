package utils

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func RunDay(solution1, solution2 func(string) (string, error))  error {
	var test1, test2, real1, real2 bool

	if len(os.Args) > 1 {
		for _, arg := range os.Args {
			switch arg{
			case "test":
				test1 = true
				test2 = true
			case "real":
				real1 = true
				real2 = true
			case "test1":
				test1 = true
			case "test2": 
			test2 = true
			case "real1":
				real1 = true
			case "real2":
				real2 = true

			}

		}

	} else {
		test1, test2, real1, real2 = true, true, true, true
	}
	currentDir, wdErr := os.Getwd()
	if wdErr != nil {
		return fmt.Errorf("failed to get workdir: %+v", wdErr)
	}
	base := filepath.Base(currentDir)
	dayNumber := strings.TrimPrefix(base, "day")
	log.Printf("Day %s\n", dayNumber)

	testData, testErr := ioutil.ReadFile("test.data")
	if testErr != nil {
		return fmt.Errorf("failed to read test data: %+v\n", testErr)
	}

	realData, realErr := ioutil.ReadFile("real.data")
	if realErr != nil {
		return fmt.Errorf("failed to read real data: %+v\n", testErr)
	}
if test1 {
	p1TestAnswer, p1TestSolErr := solution1(string(testData))
	if p1TestSolErr != nil {
		return fmt.Errorf("failed to run Part 1 test solution: %+v\n", p1TestSolErr)

	}
	log.Printf("Part 1: test answer is %s\n", p1TestAnswer)
}
if real1 {
	p1RealAnswer, p1RealSolErr := solution1(string(realData))
	if p1RealSolErr != nil {
		return fmt.Errorf("failed to run p1Real solution: %+v\n", p1RealSolErr)

	}

	log.Printf("Part1: real answer is %s\n", p1RealAnswer)
}
if test2 {
	p2TestAnswer, p2TestSolErr := solution2(string(testData))
	if p2TestSolErr != nil {
		return fmt.Errorf("failed to run Part 2 test solution: %+v\n", p2TestSolErr)

	}
	log.Printf("Part 2: test answer is %s\n", p2TestAnswer)
}
if real2 {
	p2RealAnswer, p2RealSolErr := solution2(string(realData))
	if p2RealSolErr != nil {
		return fmt.Errorf("failed to run p2Real solution: %+v\n", p2RealSolErr)

	}

	log.Printf("Part2: real answer is %s\n", p2RealAnswer)
}
	return nil



}
