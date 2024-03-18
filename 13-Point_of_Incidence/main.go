package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(n int) {
	*s = append(*s, n) // Simply append the new value to the end of the stack
}

var rows []int
var columns []int

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

	var colStrings []string

	sum := 0
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
			sum += evaluate()
			colStrings = make([]string, 0)
			columns = make([]int, 0)
			rows = make([]int, 0)
		}
	}
	for _, s := range colStrings {
		n, _ := strconv.ParseInt(s, 2, 64)
		columns = append(columns, int(n))
	}
	sum += evaluate()
	fmt.Printf("\n *** Total: %d\n", sum)
}

func evaluate() int {
	fmt.Printf("Rows: %v\n", rows)
	fmt.Printf("Cols: %v\n", columns)

	point := findReflectionPoint(columns)
	if point >= 0 {
		fmt.Printf("Found reflection at column %d\n", point)
		return point
	} else {
		point = findReflectionPoint(rows)
		fmt.Printf("Found reflection at row %d\n", point)
		return point * 100
	}

}

func findReflectionPoint(arr []int) int {
	size := len(arr)

	for point := 1; point < size; point++ {

		broken := false
		offset := 0
		for !broken && ((point - offset) >= 1) && ((point + offset) < size) {
			left := arr[point-1-offset]
			right := arr[point+offset]
			broken = (left != right)
			offset++
		}

		if !broken {
			return point
		}
	}

	return -1
}

func findReflectionPointOld(arr []int) int {
	for i, n := range arr {
		fmt.Printf("%d) %08b -> %d\n", i, n, n)
	}

	point := -1

	for testPosition := 0; testPosition < len(arr); testPosition++ {
		testSpan := 0
		broken := false

		for !broken {
			left := arr[testPosition-testSpan]
			right := ^arr[testPosition+1+testSpan]
			broken = (left^right > 0)
			testSpan++
			if (testPosition-testSpan) < 0 || (testPosition+1+testSpan) >= len(arr) {
				break
			}
		}
		if !broken {
			point = testPosition
			break
		}
	}

	return point
}
