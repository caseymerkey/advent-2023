package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//cardRE := `Card +([\d]+): ([\d ]+)\| ([\d ]+)`
	cardRE := regexp.MustCompile(`Card +([\d]+): ([\d ]+)\| ([\d ]+)`)
	splitRE := regexp.MustCompile(` +`)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		matches := cardRE.FindStringSubmatch(line)
		cardNumber := matches[1]
		playNumbers := splitRE.Split(strings.Trim(matches[2], " "), -1)
		winningNumbers := splitRE.Split(strings.Trim(matches[3], " "), -1)

		//fmt.Printf("Card %s : %s %s\n", cardNumber, playNumbers, winningNumbers)

		value := 0
		for _, number := range playNumbers {
			if slices.Contains(winningNumbers, number) {
				if value > 0 {
					value = value * 2
				} else {
					value = 1
				}
			}
		}

		fmt.Printf("Card %s value: %d\n", cardNumber, value)
		total += value
	}

	fmt.Printf("Total Value: %d", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
