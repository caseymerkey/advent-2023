package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinate struct {
	x int
	y int
}

type Connector struct {
	north int
	south int
	east int
	west int
}

var maze [][]byte
var mazeWidth int
var mazeLength int

var mapKey = map[byte] Connector{
	'|': Connector{north: -1, south: 1},
	'-': Connector{west: -1, east: 1},
	'L': Connector{north: -1, east: 1},
	'J': Connector{north: -1, west: -1},
	'7': Connector{south: 1, west: -1},
	'F': Connector{south: 1, east: 1},
}



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
	evaluate()
}

func evaluate() {

	path := []Coordinate{}

	start := findStart()
	currentLocation := start
	path = append(path, start)

	exits := findExits(start)

	fmt.Println(currentLocation)

}

func findStart() Coordinate {

	for row := 0; row < mazeLength; row++ {
		for col := 0; col < mazeWidth; col++ {
			if maze[row][col] == 'S' {
				return Coordinate{x: col, y: row}
			}
		}
	}
	return Coordinate{x: -1, y: -1}
}

func findExits(point Coordinate) []Coordinate {

	exits := []Coordinate{}
	if point.y -1 >= 0 {
		ch := maze[point.y - 1][point.x]
		if (ch == '|' || ch == 'F' || ch == '7') {
			exits = append(exits, Coordinate{x:[point.x], y: [point.y]-1 })
		}
	}

}
