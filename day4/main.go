package main

import (
	"fmt"

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

type Board struct {
	Rows []Row
	Id int
}

type Row map[int]bool

func (b *Board) call(ball int) {
	for _, row := range b.Rows {
		if _, found := row[ball]; found {
			row[ball] = true
		}
	}
}

func (b *Board) line() bool {
	for _, row := range b.Rows {
		line:= true
		for _, called := range row {
			if ! called {
				line = false
				break
			}
		}
		if line {
			return true
		}
	}
	return false
}

func (b *Board) lineScore() (int,int) {
	marked, unmarked := 0, 0
	for _, row := range b.Rows {
		log.Printf("checkung %d %+v\n")
		for ball, called := range row {
			if ! called {
				marked += ball
			} else {
				unmarked += ball
			}

		}
	}
	return unmarked , marked
}


func NewBoard(lines []string, start int) (Board, error) {
	log.Printf("Creating board %d\n", start)
	board := Board{
		Id: start /5,
		Rows: make([]Row,0),
	}
	for i:= 0;i < 5 ; i++ {
		row := Row{}
		for _, numberStr := range strings.Fields(lines[start +i]) {
			number, nErr := strconv.Atoi(numberStr)
			if nErr != nil {
				return Board{}, fmt.Errorf("failed to convert number from %s from offset %d: %+v", numberStr, start, nErr)
			}
			row[number] = false
		}
		board.Rows = append(board.Rows, row)
	}
	return board, nil
}

func solution1(puzzleInput string) (string, error) {

	lines := strings.Split(strings.TrimSpace(puzzleInput), "\n")

	boards := make([]Board, 0)
	//calls := lines[0]

	for i:= 2; i< len(lines) - 4; i+=6 {
		log.Printf("Board %d is \n%s\n%s\n%s\n%s\n%s\n\n\n", i/5, lines[i], lines[i+1], lines[i+2], lines[i+3], lines[i+4])

		board, bErr := NewBoard(lines, i)
		if bErr != nil {
			return "", fmt.Errorf("failed to build board %d: %+v", i, bErr)
		}
		boards = append(boards, board)
	}

	// now feed the numbers in and see if any boards are done
	for _, number := range strings.Split(lines[0], ",") {
		ball, bErr := strconv.Atoi(number)
		if bErr != nil {
			return "", fmt.Errorf("Failed to convert %s to a number:%+v", number, bErr)
		}
		log.Printf("Calling %d\n", ball)
		for _, board := range boards {
			board.call(ball)
			if board.line() {
				core, _ := board.lineScore()
				log.Printf ("Winner %d: %d\n", board.Id, core )
		}
	}
}
	

	return "", nil
}

func solution2(puzzleInput string) (string, error) {
	
	return "", nil
}
