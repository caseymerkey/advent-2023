package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Coordinate struct {
	col int
	row int
}

var maze [][]byte
var mazeWidth int
var mazeLength int

func main() {

	inputFile := "sample2.txt"
	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		inputFile = os.Args[1]
	}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if row == 0 {
			mazeWidth = len(line)
		}
		b := []byte(line)
		maze = append(maze, b)
		row++
	}
	mazeLength = row
	path := FindPath()
	part1(path)
	part2(path)
}

func FindPath() []Coordinate {
	path := []Coordinate{}

	start := findStart()
	path = append(path, start)

	nextCell := findExitFromStart(start)
	path = append(path, nextCell)
	fmt.Printf("%s %+v --> ", string(maze[start.row][start.col]), start)
	fmt.Printf("%s %+v\n", string(maze[nextCell.row][nextCell.col]), nextCell)

	var exits []Coordinate
	for !reflect.DeepEqual(nextCell, start) {
		fmt.Printf("%s %+v --> ", string(maze[nextCell.row][nextCell.col]), nextCell)
		exits = findExits(nextCell)

		previous := path[len(path)-2]

		beenThere := reflect.DeepEqual(previous, exits[0])

		if !beenThere {
			nextCell = exits[0]
		} else {
			nextCell = exits[1]
		}

		path = append(path, nextCell)
		fmt.Printf("%s %+v\n", string(maze[nextCell.row][nextCell.col]), nextCell)

	}
	return path
}

func part1(path []Coordinate) {

	fmt.Printf("Full path is %d. Midpoint is %d\n", len(path), len(path)/2)
}

func part2(path []Coordinate) {
	area := AreaShoelace(path)
	fmt.Printf("Area: %d\n", area)
	fmt.Printf("Interior Points: %d\n", InteriorPoints(area, len(path)))
}

func findStart() Coordinate {

	for row := 0; row < mazeLength; row++ {
		for col := 0; col < mazeWidth; col++ {
			if maze[row][col] == 'S' {
				return Coordinate{col: col, row: row}
			}
		}
	}
	return Coordinate{col: -1, row: -1}
}

func findExitFromStart(point Coordinate) Coordinate {
	if point.row > 0 {
		// North
		ch := maze[point.row-1][point.col]
		if ch == '|' || ch == 'F' || ch == '7' || ch == 'S' {
			return Coordinate{col: point.col, row: point.row - 1}
		}
	}
	if point.row < mazeLength {
		// South
		ch := maze[point.row+1][point.col]
		if ch == '|' || ch == 'L' || ch == 'J' || ch == 'S' {
			return Coordinate{col: point.col, row: point.row + 1}
		}
	}
	if point.col > 0 {
		// West
		ch := maze[point.row][point.col-1]
		if ch == '-' || ch == 'L' || ch == 'F' || ch == 'S' {
			return Coordinate{col: point.col - 1, row: point.row}
		}
	}
	if point.col < mazeWidth {
		// West
		ch := maze[point.row][point.col+1]
		if ch == '-' || ch == 'J' || ch == '7' || ch == 'S' {
			return Coordinate{col: point.col + 1, row: point.row}
		}
	}
	return point
}

func findExits(point Coordinate) []Coordinate {

	exits := []Coordinate{}
	ch := maze[point.row][point.col]

	switch ch {
	case '|':
		if point.row > 0 {
			exits = append(exits, Coordinate{row: point.row - 1, col: point.col})
		}
		if point.row < mazeLength {
			exits = append(exits, Coordinate{row: point.row + 1, col: point.col})
		}
	case '-':
		if point.col > 0 {
			exits = append(exits, Coordinate{row: point.row, col: point.col - 1})
		}
		if point.col < mazeWidth {
			exits = append(exits, Coordinate{row: point.row, col: point.col + 1})
		}
	case 'L':
		if point.row > 0 {
			exits = append(exits, Coordinate{row: point.row - 1, col: point.col})
		}
		if point.col < mazeWidth {
			exits = append(exits, Coordinate{row: point.row, col: point.col + 1})
		}
	case 'J':
		if point.row > 0 {
			exits = append(exits, Coordinate{row: point.row - 1, col: point.col})
		}
		if point.col > 0 {
			exits = append(exits, Coordinate{row: point.row, col: point.col - 1})
		}
	case '7':
		if point.row < mazeLength {
			exits = append(exits, Coordinate{row: point.row + 1, col: point.col})
		}
		if point.col > 0 {
			exits = append(exits, Coordinate{row: point.row, col: point.col - 1})
		}
	case 'F':
		if point.row < mazeLength {
			exits = append(exits, Coordinate{row: point.row + 1, col: point.col})
		}
		if point.col < mazeWidth {
			exits = append(exits, Coordinate{row: point.row, col: point.col + 1})
		}
	}

	return exits
}

func AreaShoelace(path []Coordinate) int {

	length := len(path)
	path = append(path, path[0])
	sum := 0
	for i := 0; i < length; i++ {
		sum += path[i].col * path[i+1].row
		sum -= path[i+1].col * path[i].row
	}

	return Abs(sum / 2)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func InteriorPoints(area int, boundaryPoints int) int {

	return -1 * ((boundaryPoints / 2) - 1 - area)

}
