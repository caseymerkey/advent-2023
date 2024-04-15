package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var startTime = time.Now()

	var puzzle [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		b := []byte(line)
		puzzle = append(puzzle, b)
	}
	fmt.Println(puzzle)
	puzzle = transpose(puzzle)
	fmt.Println(puzzle)

	for i, row := range puzzle {
		puzzle[i] = shiftRow(row)
	}
	fmt.Println(puzzle)

	total := 0

	for x := 0; x < len(puzzle[0]); x++ {
		for y := 0; y < len(puzzle); y++ {
			if puzzle[x][y] == 'O' {
				total += (y + 1)
			}
		}
	}

	executionTime := float32(time.Now().Sub(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Total = %d\nCompleted in %f sec\n", total, executionTime)
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
func transpose(slice [][]byte) [][]byte {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]byte, xl)
	for i := range result {
		result[i] = make([]byte, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[yl-j-1][i]
		}
	}
	return result
}
