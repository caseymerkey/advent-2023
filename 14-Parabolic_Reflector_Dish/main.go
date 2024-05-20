package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

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

	puzzle = rotate(puzzle, 1)

	part1Copy := make([][]byte, len(puzzle))
	for i := range puzzle {
		part1Copy[i] = make([]byte, len(puzzle[i]))
		copy(part1Copy[i], puzzle[i])
	}

	fmt.Println("Part 1 Input --->")
	fmt.Println(toString(part1Copy))
	fmt.Println()
	part1(part1Copy)

	fmt.Println("Part 2 Input --->")
	fmt.Println(toString(rotate(puzzle, -1)))
	fmt.Println()
	part2(puzzle)
}

func part1(puzzle [][]byte) {

	var startTime = time.Now()
	for i, row := range puzzle {
		puzzle[i] = shiftRow(row)
	}
	fmt.Println(toString(puzzle))

	total := total(puzzle)

	executionTime := float32(time.Now().Sub(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Part 1 Total = %d\nCompleted in %f sec\n", total, executionTime)
}

func part2(puzzle [][]byte) {
	var startTime = time.Now()

	var str string
	var history map[string]int
	history = make(map[string]int)
	var indexedPuzzleStates []string

	goal := 1000000000
	for cycle := 0; cycle < goal; cycle++ {
		for n := 0; n < 4; n++ {
			for i, row := range puzzle {
				puzzle[i] = shiftRow(row)
			}
			puzzle = rotate(puzzle, 1)
		}

		str = toString(puzzle)
		loopStart, exists := history[str]
		if exists {
			// we've seen this before and will be repeating from here out
			fmt.Printf("Found a repeat. %d repeats %d\n\n", cycle, loopStart)
			loopLength := cycle - loopStart
			cyclesRemaining := goal - cycle - 1
			offset := cyclesRemaining % loopLength
			idx := loopStart + offset
			puzzleString := indexedPuzzleStates[idx]
			puzzle = make([][]byte, 0)
			for _, s := range strings.Split(puzzleString, "\n") {
				if len(s) > 0 {
					puzzle = append(puzzle, []byte(s))
				}
			}
			break
		} else {
			// Add it to the history
			history[str] = cycle
			indexedPuzzleStates = append(indexedPuzzleStates, str)
		}
	}

	fmt.Println(toString(rotate(puzzle, -1)))

	total := total(puzzle)
	executionTime := float32(time.Now().Sub(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Part 2 Total = %d\nCompleted in %f sec\n", total, executionTime)
}

func shiftRow(row []byte) []byte {
	lastOpenSpace := -1
	idx := len(row) - 1
	for idx >= 0 {

		b := row[idx]
		switch b {
		case 'O':
			if lastOpenSpace > idx {
				row[idx] = '.'
				row[lastOpenSpace] = 'O'
				lastOpenSpace--
			}
		case '.':
			if lastOpenSpace == -1 {
				lastOpenSpace = idx
			}
		case '#':
			lastOpenSpace = -1
		}

		idx--
	}
	return row
}

// rotate grid 90 degrees clockwise
// "North" becomes
func rotate(slice [][]byte, direction int) [][]byte {
	if direction == 0 {
		return slice
	}

	size := len(slice[0])

	result := make([][]byte, size)
	for i := range result {
		result[i] = make([]byte, size)
	}
	if direction > 0 {
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				result[r][c] = slice[size-c-1][r]
			}
		}
	} else {
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				result[r][c] = slice[c][size-r-1]
			}
		}
	}
	return result
}

func toString(puzzle [][]byte) string {

	var buffer bytes.Buffer
	for _, b := range puzzle {
		buffer.Write(b)
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

func total(puzzle [][]byte) int {
	total := 0

	for row := 0; row < len(puzzle[0]); row++ {
		for col := 0; col < len(puzzle); col++ {
			if puzzle[row][col] == 'O' {
				total += (col + 1)
			}
		}
	}
	return total
}
