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
	Elements []Element
	Id       int
}
type Element struct {
	X      int
	Y      int
	Val    int
	Called bool
}

func (b *Board) rows() [5][5]Element {
	rows := [5][5]Element{}

	for _, element := range b.Elements {
		rows[element.X][element.Y] = element
	}
	return rows
}

func (b *Board) columns() [5][5]Element {
	columns := [5][5]Element{}

	for _, element := range b.Elements {
		columns[element.Y][element.X] = element
	}
	return columns
}

func (b *Board) call(ball int) {
	for elemIndex, element := range b.Elements {
		if element.Val == ball {
			b.Elements[elemIndex].Called = true
		}
	}
}

func (b *Board) line() bool {
	for _, row := range b.rows() {
		line := true
		for _, cell := range row {
			if !cell.Called {
				line = false
				break
			}
		}
		if line {
			return true
		}
	}
	for _, column := range b.columns() {
		line := true
		for _, cell := range column {
			if !cell.Called {
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

func (b *Board) lineScore() (int, int) {
	marked, unmarked := 0, 0
	for _, cell := range b.Elements {
		if cell.Called {
			marked += cell.Val
		} else {
			unmarked += cell.Val
		}

	}
	return unmarked, marked
}

func NewBoard(lines []string, start int) (Board, error) {
	board := Board{
		Id:       start / 5,
		Elements: make([]Element, 0),
	}
	for i := 0; i < 5; i++ {
		for x, numberStr := range strings.Fields(lines[start+i]) {
			number, nErr := strconv.Atoi(numberStr)
			if nErr != nil {
				return Board{}, fmt.Errorf("failed to convert number from %s from offset %d: %+v", numberStr, start, nErr)
			}
			board.Elements = append(board.Elements, Element{
				X:      x,
				Y:      i,
				Called: false,
				Val:    number,
			})
		}
	}
	return board, nil
}

func solution1(puzzleInput string) (string, error) {
	lines := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	boards := make([]Board, 0)

	for i := 2; i < len(lines)-4; i += 6 {
		board, bErr := NewBoard(lines, i)
		if bErr != nil {
			return "", fmt.Errorf("failed to build board %d: %+v", i, bErr)
		}
		boards = append(boards, board)
	}

	for _, number := range strings.Split(lines[0], ",") {
		ball, bErr := strconv.Atoi(number)
		if bErr != nil {
			return "", fmt.Errorf("Failed to convert %s to a number:%+v", number, bErr)
		}

		for _, board := range boards {
			board.call(ball)
			if board.line() {
				score, _ := board.lineScore()
				return fmt.Sprintf("%d", ball*score), nil
			}
		}
	}
	return "", nil
}

func solution2(puzzleInput string) (string, error) {

	lines := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	boards := make([]Board, 0)

	for i := 2; i < len(lines)-4; i += 6 {
		board, bErr := NewBoard(lines, i)
		if bErr != nil {
			return "", fmt.Errorf("failed to build board %d: %+v", i, bErr)
		}
		boards = append(boards, board)
	}
	completeBoards := make([]int, 0)
	completeBoardMap := map[int]bool{}
	for _, number := range strings.Split(lines[0], ",") {
		ball, bErr := strconv.Atoi(number)
		if bErr != nil {
			return "", fmt.Errorf("Failed to convert %s to a number:%+v", number, bErr)
		}

		for _, board := range boards {

			board.call(ball)

			if board.line() {
				if _, found := completeBoardMap[board.Id]; !found {
					completeBoardMap[board.Id] = true
					completeBoards = append(completeBoards, board.Id)
				}
			}
		}

		if len(completeBoards) == len(boards) {
			winner := completeBoards[len(completeBoards)-1]
			for _, board := range boards {
				if board.Id == winner {
					score, _ := board.lineScore()
					return fmt.Sprintf("%d", ball*score), nil
				}
			}
		}

	}
	return "", nil
}