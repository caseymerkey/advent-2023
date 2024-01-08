package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`[\?#]+`)

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

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += evaluate(line)
	}
	fmt.Printf("Total: %d\n", total)
}

func stringToIntArr(listOfNumbers string) []int {
	a := strings.Split(listOfNumbers, ",")
	numbers := []int{}

	for _, n := range a {
		nn, _ := strconv.Atoi(n)
		numbers = append(numbers, nn)
	}
	return numbers
}

func evaluate(line string) int {

	s := strings.Split(line, " ")
	springs := s[0]
	groups := stringToIntArr(s[1])
	sort.Slice(groups, func(a, b int) bool {
		return groups[a] > groups[b]
	})

	matches := re.FindAllString(springs, -1)

	for i, region := range matches {
		fmt.Printf("Region #%d: %s\n", i, region)
	}

	return len(springs)
}
