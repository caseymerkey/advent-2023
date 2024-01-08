package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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

	result := 0
	directionsLength := len(directions)

	for _, n := range startingNodes {
		stepsTaken := 0
		currentNode := nodeMap[n]
		for !strings.HasSuffix(currentNode.name, "Z") {
			i := stepsTaken % directionsLength
			direction := directions[i]
			stepsTaken++
			if direction == 'L' {
				currentNode = nodeMap[currentNode.L]
			} else {
				currentNode = nodeMap[currentNode.R]
			}
		}
		if result == 0 {
			result = stepsTaken
		} else {
			result = lcm(result, stepsTaken)
		}
	}
	executionTime := float32(time.Now().Sub(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed in %f seconds\nSteps Taken: %d", executionTime, result)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
