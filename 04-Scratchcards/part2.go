package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
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

	cardCopiesMap := make(map[int]int)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := cardRE.FindStringSubmatch(line)
		cardNumber, _ := strconv.Atoi(matches[1])
		playNumbers := splitRE.Split(strings.Trim(matches[2], " "), -1)
		winningNumbers := splitRE.Split(strings.Trim(matches[3], " "), -1)

		// First things first. How many copies of this card do I have
		// 1 plus the number of copies previously earned
		copyCountCurrentCard := 1
		if copies, exists := cardCopiesMap[cardNumber]; exists {
			copyCountCurrentCard += copies
		}

		fmt.Printf("%d copies of Card %d\n", copyCountCurrentCard, cardNumber)
		total += copyCountCurrentCard

		// Now, how many copies of other cards am I gaining
		matchCount := 0
		for _, number := range playNumbers {
			if slices.Contains(winningNumbers, number) {
				matchCount++
			}
		}
		fmt.Printf("  Card %d has %d matches\n", cardNumber, matchCount)

		for i := 1; i <= matchCount; i++ {
			cardCopyNumber := cardNumber + i
			cardCopyCount := cardCopiesMap[cardCopyNumber]
			fmt.Printf("  Adding %d copies of card %d, total copies = %d\n", copyCountCurrentCard, cardCopyNumber, cardCopyCount+copyCountCurrentCard)
			cardCopiesMap[cardCopyNumber] = cardCopyCount + copyCountCurrentCard
		}

	}

	fmt.Printf("Total number of cards: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
