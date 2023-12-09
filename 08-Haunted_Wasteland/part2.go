package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

type Node struct {
	name, L, R string
}

var directions string
var nodeMap = map[string]Node{}
var startingNodes []string
var startTime = time.Now()

// This doesn't work yet.

func main() {

	inputFile := "input.txt"
	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		inputFile = os.Args[1]
	}
	nodeRE := regexp.MustCompile(`([A-Z1-9]{3}) = \(([A-Z1-9]{3}), ([A-Z1-9]{3})\)`)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if directions == "" {
			directions = line
		} else {
			fields := nodeRE.FindStringSubmatch(line)
			if len(fields) > 0 {
				nodeMap[fields[1]] = Node{name: fields[1], L: fields[2], R: fields[3]}
				if fields[1][2] == 'A' {
					startingNodes = append(startingNodes, fields[1])
					fmt.Println(fields[1])
				}
			}
		}
	}

	evaluate()
}

func evaluate() {

	stepsTaken := 0
	directionsLength := len(directions)

	for !areWeThereYet() {
		i := stepsTaken % directionsLength
		direction := directions[i]
		stepsTaken++
		for pathNumber, currentNode := range startingNodes {

			if direction == 'L' {
				startingNodes[pathNumber] = nodeMap[currentNode].L
			} else {
				startingNodes[pathNumber] = nodeMap[currentNode].R
			}

		}
	}
	executionTime := float32(time.Now().Sub(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed in %f seconds\nSteps Taken: %d", executionTime, stepsTaken)
}

func areWeThereYet() bool {
	for _, v := range startingNodes {
		if v[2] != 'Z' {
			return false
		}
	}
	return true
}
