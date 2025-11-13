package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Instruction struct {
	Direction string
	Distance  int
	Color     string
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

	re := regexp.MustCompile(`([LRUD]) ([\d]+) \(#([0-9a-f]+)\)`)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		d, _ := strconv.Atoi(matches[2])
		i := Instruction{Direction: matches[1], Distance: d, Color: matches[3]}
		instructions = append(instructions, i)
	}

	var startTime = time.Now()
	p1 := part1(instructions)
	fmt.Printf("Part 1: %d\n", p1)
	executionTime := float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 1 in %f seconds\n", executionTime)

	startTime = time.Now()
	p2 := part2(instructions)
	fmt.Printf("Part 2: %d\n", p2)
	executionTime = float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 2 in %f seconds\n", executionTime)
}

type Coord struct {
	row int
	col int
}

func (c Coord) Move(direction Coord, spaces int) Coord {
	return Coord{row: c.row + (direction.row * spaces), col: c.col + (direction.col * spaces)}
}

func (c Coord) String() string {
	return fmt.Sprintf("{x:%d, y:%d}", c.col, c.row)
}

func part1(instructions []Instruction) int {
	return shoelace(instructions)
}

func part2(instructions []Instruction) int {
	modifiedList := make([]Instruction, len(instructions))
	for k, inst := range instructions {
		d, _ := strconv.ParseInt(inst.Color[:5], 16, strconv.IntSize)
		inst.Distance = int(d)
		switch inst.Color[5] {
		case '0':
			inst.Direction = "R"
		case '1':
			inst.Direction = "D"
		case '2':
			inst.Direction = "L"
		case '3':
			inst.Direction = "U"
		}
		modifiedList[k] = inst
	}

	return shoelace(modifiedList)
}

func shoelace(instructions []Instruction) int {

	var directionMap = map[string]Coord{
		"R": {row: 0, col: 1},
		"L": {row: 0, col: -1},
		"U": {row: -1, col: 0},
		"D": {row: 1, col: 0},
	}
	area := 0
	start := Coord{0, 0}

	for _, i := range instructions {

		dir := directionMap[i.Direction]
		next := start.Move(dir, i.Distance)

		// combines Shoelace and Pick's theorem into the same step
		area += ((start.col*next.row - start.row*next.col) + i.Distance)
		start = next
	}

	return 1 + (int(math.Abs(float64(area))) / 2)
}
