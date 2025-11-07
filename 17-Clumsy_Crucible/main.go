package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	col int
	row int
}
type Direction struct {
	col int
	row int
}

type State struct {
	loc    Coord
	facing Direction
	steps  int
}

type Item struct {
	value    State // The value of the item
	priority int   // The priority of the item
	index    int   // The index of the item in the heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// For a min-heap, use pq[i].priority < pq[j].priority
	// For a max-heap, use pq[i].priority > pq[j].priority
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, value State, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

var EAST = Direction{row: 0, col: 1}
var SOUTH = Direction{row: 1, col: 0}
var WEST = Direction{row: 0, col: -1}
var NORTH = Direction{row: -1, col: 0}

var Directions = []Direction{EAST, SOUTH, WEST, NORTH}

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

	var puzzle [][]int
	for scanner.Scan() {
		line := scanner.Text()
		valStrings := strings.Split(line, "")
		row := make([]int, 0)
		for _, v := range valStrings {
			n, _ := strconv.Atoi(v)
			row = append(row, n)
		}
		puzzle = append(puzzle, row)
	}

	var startTime = time.Now()
	p1 := part1(puzzle)
	fmt.Printf("Part 1: %d\n", p1)
	executionTime := float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 1 in %f seconds\n", executionTime)

}

func part1(puzzle [][]int) int {
	target := Coord{row: len(puzzle) - 1, col: len(puzzle[0]) - 1}
	return dijkstra(puzzle, Coord{0, 0}, target)
}

func dijkstra(puzzle [][]int, start, target Coord) int {

	pq := make(PriorityQueue, 0)
	visited := make(map[State]bool)

	eastStart := State{loc: Coord{row: 0, col: 1}, facing: EAST, steps: 1}
	southStart := State{loc: Coord{row: 1, col: 0}, facing: SOUTH, steps: 1}

	heap.Push(&pq, &Item{value: eastStart, priority: puzzle[0][1]})
	heap.Push(&pq, &Item{value: southStart, priority: puzzle[1][0]})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		state := item.value
		if visited[state] {
			continue
		}
		visited[state] = true
		if state.loc == target {
			return item.priority
		}

		neighbors := neighbors(state, len(puzzle), len(puzzle[0]))
		for _, neighbor := range neighbors {

			newHeat := item.priority + puzzle[neighbor.loc.row][neighbor.loc.col]
			heap.Push(&pq, &Item{value: neighbor, priority: newHeat})
		}
		visited[state] = true
	}

	return 0
}

func neighbors(state State, rows, cols int) []State {
	neighbors := make([]State, 0)
	for _, dir := range leftStraightRight(state.facing) {
		c := Coord{row: state.loc.row + dir.row, col: state.loc.col + dir.col}
		if c.col < 0 || c.col >= cols || c.row < 0 || c.row >= rows {
			continue
		}
		newState := State{loc: c, facing: dir, steps: 1}
		if dir == state.facing {
			if state.steps >= 3 {
				continue
			}
			newState.steps = state.steps + 1

		}
		neighbors = append(neighbors, newState)
	}
	return neighbors
}

func leftStraightRight(dir Direction) []Direction {
	var options []Direction
	switch dir {
	case EAST:
		options = []Direction{NORTH, EAST, SOUTH}
	case WEST:
		options = []Direction{SOUTH, WEST, NORTH}
	case NORTH:
		options = []Direction{WEST, NORTH, EAST}
	case SOUTH:
		options = []Direction{EAST, SOUTH, WEST}
	}
	return options
}
