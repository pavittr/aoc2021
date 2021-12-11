package main

import (
	"fmt"


	"log"

	"strings"



	"github.com/pavittr/aoc2021/utils"
)

func main() {
	if err := utils.RunDay(solution1, solution2); err != nil {
		log.Fatalf("Failed to run day: %+v\n", err)
	}
}

func solution1(puzzleInput string) (string, error) {
	inputLines := strings.Split(strings.TrimSpace(puzzleInput), "\n")

	// look for 1 4, 7, 8
	//          2, 4, 3, 7
	segmentsOf1478 := 0
	for _, inputLine := range inputLines {
		inputOutput := strings.Split(inputLine, "|")
		output := inputOutput[1]
		outputSegments := strings.Split(output, " ")
		for _, outputSegment := range outputSegments {
			segmentLength := len(outputSegment)
			if segmentLength == 2 ||
			   segmentLength == 4 ||
			   segmentLength == 3 ||
			   segmentLength == 7 {
				   segmentsOf1478++
			   }
		}
	}

	return fmt.Sprintf("%d", segmentsOf1478), nil
}

func solution2(puzzleInput string) (string, error) {
	inputLines := strings.Split(strings.TrimSpace(puzzleInput), "\n")

	totalCounter := 0
	for _, inputLine := range inputLines {
		// Assume a structure :
		//     aa
		//    b  c
		//     dd
		//    e  f
		//     gg
		// Use the input to calculate this mapping
		// Then use the mappign to map the output to the relevent string
		intMap := map[string]int{}
		segments := map[int][]string{}
		knownPositions := map[rune]rune{}
		inputOutput := strings.Split(inputLine, "|")
		input := inputOutput[0]
		uniqueDisplays := strings.Split(input, " ")
		for _, display := range uniqueDisplays {
			segments[len(display)] = append(segments[len(display)], display)
		}
		one := segments[2][0]
		four := segments[4][0]
		seven := segments[3][0]
		eight := segments[7][0]
		intMap[one] = 1
		intMap[four] = 4
		intMap[seven] = 7
		intMap[eight] = 8
		// need to find 0, 2,3, 5, 6, and 9
		// for six seg, if it contains all of 4 then its a 9
		nine := ""
		for _, sixSeg := range segments[6] {
			// break the sixSeg into runes and check if ever
			matches := true
			for _, fourSegRune := range four {
				if !strings.ContainsRune(sixSeg, fourSegRune) {
					matches = false
					break

				}
			}
			if matches {
				intMap[sixSeg] = 9
				nine = sixSeg
			}
		}

		// need to find 0, 2, 3, 5, and 6
				
		six := ""
		for _, sixSeg := range segments[6] {
			// break the sixSeg into runes and check if ever
			matches := true
			for _, oneSegRune := range one {
				if !strings.ContainsRune(sixSeg, oneSegRune) {
					matches = false
					break

				}
			}
			if matches {
				intMap[sixSeg] = 6
				six = sixSeg
			}
		}
		// The remaining sixSeg is 0
		zero := ""
		for _, sixSeg := range segments[6] {
			if sixSeg != six && sixSeg != nine {
				zero = sixSeg
				intMap[zero] = 0
			}
		}

		// need to find 2, 3, and 5
		// if it contains all of 1 then its 3
		three := ""
		for _, fiveSeg := range segments[5] {
			// break the fiveSeg into runes and check if ever
			matches := true
			for _, oneSegRune := range one {
				if !strings.ContainsRune(fiveSeg, oneSegRune) {
					matches = false
					break

				}
			}
			if matches {
				three = fiveSeg
				intMap[three] = 3
			}
		}

		// need to fidn 2 and 5
		// If you add the fiveSeg of five to one, then you should match 9. if not, its a 2





		output := inputOutput[1]
		outputSegments := strings.Split(output, " ")
		for _, outputSegment := range outputSegments {
			segmentLength := len(outputSegment)
			if segmentLength == 2 ||
			   segmentLength == 4 ||
			   segmentLength == 3 ||
			   segmentLength == 7 {
				   segmentsOf1478++
			   }
		}
	}

	return fmt.Sprintf("%d", 0), nil

}
