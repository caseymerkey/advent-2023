package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var rows []int
	var columns []int
	var colStrings []string
	part1Sum := 0
	part2Sum := 0
	patternNumber := 1

	var startTime = time.Now()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			line = strings.ReplaceAll(line, "#", "1")
			line = strings.ReplaceAll(line, ".", "0")

			n, _ := strconv.ParseInt(line, 2, 64)
			rows = append(rows, int(n))
			if len(colStrings) == 0 {
				for i := 0; i < len(line); i++ {
					colStrings = append(colStrings, "")
				}
			}
			for i := 0; i < len(line); i++ {
				ch := line[i : i+1]
				str := colStrings[i]
				colStrings[i] = str + ch
			}

		} else {
			for _, s := range colStrings {
				n, _ := strconv.ParseInt(s, 2, 64)
				columns = append(columns, int(n))
			}
			p1, p2 := evaluatePattern(rows, columns)
			fmt.Printf("Pattern %d, Part 1 Total: %d\n", patternNumber, p1)
			fmt.Printf("Pattern %d, Part 2 Total: %d\n", patternNumber, p2)
			part1Sum += p1
			part2Sum += p2
			colStrings = make([]string, 0)
			columns = make([]int, 0)
			rows = make([]int, 0)
			patternNumber++
		}
	}
	for _, s := range colStrings {
		n, _ := strconv.ParseInt(s, 2, 64)
		columns = append(columns, int(n))
	}
	p1, p2 := evaluatePattern(rows, columns)
	fmt.Printf("Pattern %d, Part 1 Total: %d\n", patternNumber, p1)
	fmt.Printf("Pattern %d, Part 2 Total: %d\n", patternNumber, p2)

	part1Sum += p1
	part2Sum += p2
	executionTime := float32(time.Now().Sub(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("\nCompleted in %f seconds\n", executionTime)
	fmt.Printf("\n *** Part 1 Total: %d\n", part1Sum)
	fmt.Printf("\n *** Part 2 Total: %d\n", part2Sum)
}

func evaluatePattern(rows []int, cols []int) (int, int) {

	p1 := 0

	fmt.Println("Rows Part 1")
	p1 += 100 * findReflectionPoints(rows, 0)
	fmt.Println("Columns Part 1")
	p1 += findReflectionPoints(cols, 0)

	p2 := 0
	fmt.Println("Rows Part 2")
	p2 += 100 * findReflectionPoints(rows, 1)
	fmt.Println("Columns Part 1")
	p2 += findReflectionPoints(cols, 1)

	return p1, p2

}

func findReflectionPoints(arr []int, allowedSmudges int) int {

	size := len(arr)
	total := 0

	for point := 1; point < size; point++ {
		broken := false
		offset := 0
		differencesFound := 0

		for !broken && ((point - offset) >= 1) && ((point + offset) < size) {
			left := arr[point-1-offset]
			right := arr[point+offset]

			// fmt.Printf("  Comparing %d(%d) to %d(%d)... \n", point-1-offset, left, point+offset, right)
			differencesFound += differences(left, right)
			if differencesFound > allowedSmudges {
				broken = true
			}
			offset++
		}
		if !broken && (differencesFound == allowedSmudges) {
			fmt.Printf("    Reflection found at %d\n", point)
			total += point
		}

	}

	return total
}

func differences(a int, b int) int {
	d := a ^ b
	return strings.Count(strconv.FormatInt(int64(d), 2), "1")
}
