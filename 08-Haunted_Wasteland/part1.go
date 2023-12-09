package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Node struct {
	name, L, R string
}

var directions string
var nodeMap = map[string]Node{}

func main() {

	nodeRE := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if directions == "" {
			directions = line
		} else {
			fields := nodeRE.FindStringSubmatch(line)
			if len(fields) > 0 {
				nodeMap[fields[1]] = Node{name: fields[1], L: fields[2], R: fields[3]}
			}
		}
	}

	evaluate()
}

func evaluate() {
	stepsTaken := 0
	directionsLength := len(directions)
	currentNode := nodeMap["AAA"]
	for currentNode.name != "ZZZ" {
		i := stepsTaken % directionsLength
		direction := directions[i]
		stepsTaken++
		if direction == 'L' {
			currentNode = nodeMap[currentNode.L]
		} else {
			currentNode = nodeMap[currentNode.R]
		}
	}

	fmt.Println(stepsTaken)
}
