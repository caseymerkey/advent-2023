package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	// logFile, err := os.OpenFile("part1.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	inputFile := "input.txt"
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
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		numStrings := strings.Fields(line)
		slices.Reverse(numStrings)
		nextVal := evaluate(1, stringArrToIntArr(numStrings))
		fmt.Printf("%s => %d \n", line, nextVal)
		total += nextVal
	}
	fmt.Printf("Total: %d\n", total)
}

func stringArrToIntArr(numStrings []string) []int {
	var numbers = []int{}
	for _, i := range numStrings {
		n, _ := strconv.Atoi(i)
		numbers = append(numbers, n)
	}
	return numbers
}

func evaluate(level int, numbers []int) int {

	log.Println(level, numbers)
	diffNumbers := differences(numbers)
	if allZeroes(diffNumbers) {
		log.Println(numbers, " => ", numbers[0])
		return numbers[0]
	} else {
		nextInNextSequence := evaluate(level+1, diffNumbers) + numbers[len(numbers)-1]
		log.Println(numbers, " => ", nextInNextSequence)
		return nextInNextSequence
	}

}

func allZeroes(numbers []int) bool {

	return !slices.ContainsFunc(numbers, func(x int) bool { return x != 0 })

}

func differences(numbers []int) []int {
	var diffNumbers []int
	for i := 0; i < (len(numbers) - 1); i++ {
		diffNumbers = append(diffNumbers, (numbers[i+1] - numbers[i]))
	}
	return diffNumbers
}
