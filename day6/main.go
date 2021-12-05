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

type Coord struct {
	X int
	Y int
}

type Line []Coord

func coordFrom(coordStr string) (Coord, error) {
	xy := strings.Split(coordStr, ",")
	x, xErr := strconv.Atoi(xy[0])
	if xErr != nil {
		return Coord{}, fmt.Errorf("failed to parse x from %+v of line %+v: %+v\n", coordStr, xy, xErr)
	}
	y, yErr := strconv.Atoi(xy[1])
	if yErr != nil {
		return Coord{}, fmt.Errorf("failed to parse y from %+v of line %+v: %+v\n", coordStr, xy, yErr)
	}
	return Coord{x, y}, nil
}
func fromLine(line string,  includeDiagonals bool) (Line, error) {
	coords := strings.Split(line, " -> ")
	src, srcErr := coordFrom(coords[0])
	if srcErr != nil {
		return Line{}, fmt.Errorf("faile to parse src from line %s: %+v", line, srcErr)
	}
	dst, dstErr := coordFrom(coords[1])
	if dstErr != nil {
		return Line{}, fmt.Errorf("faile to parse dst from line %s: %+v", line, dstErr)
	}

	retLine := make(Line, 0)
	if src.X == dst.X {
		if src.Y > dst.Y {
			tmp := dst
			dst = src
			src = tmp
		}
		for y := src.Y; y <= dst.Y; y++ {
			retLine = append(retLine, Coord{src.X, y})
		}
	} else if src.Y == dst.Y {
		if src.X > dst.X {
			tmp := dst
			dst = src
			src = tmp
		}

		for x := src.X; x <= dst.X; x++ {
			retLine = append(retLine, Coord{x, src.Y})
		}
	} else if includeDiagonals {
		if src.X < dst.X && src.Y < dst.Y {
			for x, y := src.X, src.Y; x <= dst.X; x++ {
				retLine = append(retLine, Coord{x, y})
				y++
			}
		} else if src.X > dst.X && src.Y < dst.Y {
			for x, y := src.X, src.Y; x >= dst.X; x-- {
				retLine = append(retLine, Coord{x, y})
				y++
			}
	} else if src.X < dst.X && src.Y > dst.Y {
		for x, y := src.X, src.Y; x <= dst.X; x++ {
				retLine = append(retLine, Coord{x, y})
				y--
			}
	} else if src.X > dst.X && src.Y > dst.Y {
		for x, y := src.X, src.Y; x >= dst.X; x-- {
				retLine = append(retLine, Coord{x, y})
				y--
			}
		}

	}

	return retLine, nil

}
func solution1(puzzleInput string) (string, error) {
	txtLines := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	overlap := map[Coord]int{}
	for _, txtLine := range txtLines {
		line, lineErr := fromLine(txtLine, false)
		if lineErr != nil {
			return "", fmt.Errorf("failed to pro ess line %s: %+v", txtLine, lineErr)
		}
		for _, coord := range line {
			overlap[coord]++
		}
	}
	counter := 0
	for _, count := range overlap {
		if count > 1 {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter), nil
}

func solution2(puzzleInput string) (string, error) {
	txtLines := strings.Split(strings.TrimSpace(puzzleInput), "\n")
	overlap := map[Coord]int{}
	for _, txtLine := range txtLines {
		line, lineErr := fromLine(txtLine, true)
		if lineErr != nil {
			return "", fmt.Errorf("failed to pro ess line %s: %+v", txtLine, lineErr)
		}
		for _, coord := range line {
			overlap[coord]++
		}
	}
	counter := 0
	for _, count := range overlap {
		if count > 1 {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter), nil
}
