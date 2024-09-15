package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const E = byte(1) // 0001
const W = byte(2) // 0010
const N = byte(4) // 0100
const S = byte(8) // 1000

var dirMap = make(map[byte]string)

func main() {

	dirMap[E] = "E"
	dirMap[W] = "W"
	dirMap[N] = "N"
	dirMap[S] = "S"

	inputFile := "sample.txt"
	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		inputFile = os.Args[1]
	}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var puzzle [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		b := []byte(line)
		puzzle = append(puzzle, b)
	}

	fmt.Println("Starting part 1")

	var startTime = time.Now()
	p1 := part1(puzzle)
	fmt.Printf("Part 1: %d\n", p1)

	executionTime := float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 1 in %f seconds\n", executionTime)

	fmt.Println("\nStarting part 2")

	startTime = time.Now()
	p2 := part2(puzzle)
	fmt.Printf("Part 2: %d\n", p2)

	executionTime = float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 2 in %f seconds\n", executionTime)

}

func part2(puzzle [][]byte) int {

	maxValue := 0
	for y := 0; y < len(puzzle); y++ {
		energized := make(map[string]byte)
		advanceAndEvaluate(-1, y, E, puzzle, energized)
		if len(energized) > maxValue {
			maxValue = len(energized)
		}

		energized = make(map[string]byte)
		advanceAndEvaluate(len(puzzle[0]), y, W, puzzle, energized)
		if len(energized) > maxValue {
			maxValue = len(energized)
		}
	}

	for x := 0; x < len(puzzle[0]); x++ {
		energized := make(map[string]byte)
		advanceAndEvaluate(x, -1, S, puzzle, energized)
		if len(energized) > maxValue {
			maxValue = len(energized)
		}
		energized = make(map[string]byte)
		advanceAndEvaluate(x, len(puzzle), N, puzzle, energized)
		if len(energized) > maxValue {
			maxValue = len(energized)
		}
	}

	return maxValue
}

func part1(puzzle [][]byte) int {

	energized := make(map[string]byte)
	advanceAndEvaluate(-1, 0, E, puzzle, energized)
	return len(energized)
}

func advanceAndEvaluate(x, y int, direction byte, puzzle [][]byte, energized map[string]byte) {

	var row, col int

	switch direction {
	case E:
		row = y
		col = x + 1
	case W:
		row = y
		col = x - 1
	case N:
		row = y - 1
		col = x
	case S:
		row = y + 1
		col = x
	}

	if col < len(puzzle) && col >= 0 && row < len(puzzle[0]) && row >= 0 {

		cell := puzzle[row][col]
		cellString := nodeToString(col, row)
		cellEnergy, found := energized[cellString]

		if !found || (cellEnergy&direction != direction) {
			cellEnergy = cellEnergy | direction
			energized[cellString] = cellEnergy

			switch cell {
			case '-':
				if direction == E || direction == W {
					advanceAndEvaluate(col, row, direction, puzzle, energized)
				} else {
					advanceAndEvaluate(col, row, E, puzzle, energized)
					advanceAndEvaluate(col, row, W, puzzle, energized)
				}

			case '|':
				if direction == N || direction == S {
					advanceAndEvaluate(col, row, direction, puzzle, energized)
				} else {
					advanceAndEvaluate(col, row, N, puzzle, energized)
					advanceAndEvaluate(col, row, S, puzzle, energized)
				}

			case '/':
				switch direction {
				case E:
					advanceAndEvaluate(col, row, N, puzzle, energized)
				case W:
					advanceAndEvaluate(col, row, S, puzzle, energized)
				case N:
					advanceAndEvaluate(col, row, E, puzzle, energized)
				case S:
					advanceAndEvaluate(col, row, W, puzzle, energized)
				}

			case '\\':
				switch direction {
				case E:
					advanceAndEvaluate(col, row, S, puzzle, energized)
				case W:
					advanceAndEvaluate(col, row, N, puzzle, energized)
				case N:
					advanceAndEvaluate(col, row, W, puzzle, energized)
				case S:
					advanceAndEvaluate(col, row, E, puzzle, energized)
				}

			default:
				advanceAndEvaluate(col, row, direction, puzzle, energized)
			}
		}

	}

}

func nodeToString(x, y int) string {
	return fmt.Sprintf("%03dx%03d", x, y)
}
