package main

import (

	"log"

	"strings"

	"github.com/pavittr/aoc2021/utils"
)

func main() {
	utils.RunDay(solution1, solution2)
}

type Board struct {
	Rows []Row
	Id int
}

type Row struct {
	Numbers map[int]bool
}

func solution1(puzzleInput string) (string, error) {

	lines := strings.Split(strings.TrimSpace(puzzleInput), "\n")

	//calls := lines[0]

	for i:= 2; i< len(lines) - 4; i+=6 {
		log.Printf("Board %d is \n%s\n%s\n%s\n%s\n%s\n\n\n", i/5, lines[i], lines[i+1], lines[i+2], lines[i+3], lines[i+4])
		

	}
	

	return "", nil
}

func solution2(puzzleInput string) (string, error) {
	
	return "", nil
}
