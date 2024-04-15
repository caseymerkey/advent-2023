package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

var emptyRows []int
var emptyCols []int
var galaxies [][2]int

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

	var skymap [][]byte
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		b := []byte(line)
		skymap = append(skymap, b)
		fmt.Printf("%d: %s\n", row, line)
		row++
	}

	galaxies, emptyCols, emptyRows = findEmptyColumnsAndRows(skymap)
	fmt.Printf("Empty rows: %v\n", emptyRows)
	fmt.Printf("Empty cols: %v\n", emptyCols)

	part1()
	part2()
}

func part1() {
	d := 0
	for a := range galaxies {
		for b := a + 1; b < len(galaxies); b++ {
			if a != b {
				d += distance(galaxies[a][0], galaxies[a][1], galaxies[b][0], galaxies[b][1], 2)
			}
		}
	}
	fmt.Printf("Distance: %d\n", d)
}

func part2() {
	factor := 1000000
	d := 0
	for a := range galaxies {
		for b := a + 1; b < len(galaxies); b++ {
			if a != b {
				d += distance(galaxies[a][0], galaxies[a][1], galaxies[b][0], galaxies[b][1], factor)
			}
		}
	}
	fmt.Printf("Distance with factor %d: %d\n", factor, d)
}

func findEmptyColumnsAndRows(skymap [][]byte) ([][2]int, []int, []int) {
	var g [][2]int
	var emptyRows []int
	var emptyCols []int
	for col := range skymap[0] {
		empty := true
		for row := range skymap {
			if col == 0 && !slices.Contains(skymap[row], '#') {
				emptyRows = append(emptyRows, row)
			}
			if skymap[row][col] == '#' {
				empty = false
				p := [2]int{row, col}
				g = append(g, p)
			}
		}
		if empty {
			emptyCols = append(emptyCols, col)
		}
	}
	return g, emptyCols, emptyRows
}

func distance(row1, col1, row2, col2, factor int) int {

	var rowRange []int
	var colRange []int
	if row1 < row2 {
		rowRange = makeRange(row1, row2)
	} else {
		rowRange = makeRange(row2, row1)
	}
	if col1 < col2 {
		colRange = makeRange(col1, col2)
	} else {
		colRange = makeRange(col2, col1)
	}

	d := 0

	if len(rowRange) > 1 {
		for i, row := range rowRange {
			if i > 0 {
				if slices.Contains(emptyRows, row) {
					d += factor
				} else {
					d++
				}
			}
		}
	}

	if len(colRange) > 1 {
		for i, col := range colRange {
			if i > 0 {
				if slices.Contains(emptyCols, col) {
					d += factor
				} else {
					d++
				}
			}
		}
	}

	return int(math.Abs(float64(d)))
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
