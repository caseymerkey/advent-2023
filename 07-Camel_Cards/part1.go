package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    string
	bid      int
	value    int
	hexcards string
}

func newHand(cards string, bid int) Hand {
	h := Hand{cards: cards, bid: bid}

	cards = strings.ReplaceAll(cards, "A", "E")
	cards = strings.ReplaceAll(cards, "K", "D")
	cards = strings.ReplaceAll(cards, "Q", "C")
	cards = strings.ReplaceAll(cards, "J", "B")
	cards = strings.ReplaceAll(cards, "T", "A")
	h.hexcards = cards

	b := []byte(h.cards)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	/*
		0: high card
		1: one pair
		2: two pairs
		3: three of a kind
		4: full house
		5: four of a kind
		6: five of a kind
	*/
	value := 0
	var previousCard byte
	possibleFullHouse := false
	for _, v := range b {
		if previousCard == v {
			// see  what we've got
			switch value {
			case 0:
				// was high card, now a pair
				value = 1
			case 1:
				// was a pair, now either three of a kind, two pairs
				if possibleFullHouse {
					value = 2
				} else {
					value = 3
				}
			case 2:
				// was two pairs, now full house
				value = 4
			case 3:
				// was three of a kind, now either four of a kind or a full house
				if possibleFullHouse {
					value = 4
				} else {
					value = 5
				}
			case 4:
				// was four of a kind, now five of a kind
				value = 5
			}
		} else {
			if value > 0 {
				possibleFullHouse = true
			}
		}
		previousCard = v

	}
	h.value = value

	return h
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var allHands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		bid, _ := strconv.Atoi(fields[1])
		h := newHand(fields[0], bid)
		allHands = append(allHands, h)
	}

	sort.Slice(allHands, func(i, j int) bool {

		iValue := allHands[i].value
		jValue := allHands[j].value

		if iValue == jValue {
			return allHands[i].hexcards < allHands[j].hexcards
		} else {
			return (iValue < jValue)
		}

	})

	total := 0
	for i, h := range allHands {
		total += h.bid * (i + 1)
		fmt.Printf("%4d - cards: %s, bid: %4d, value: %d, hex: %s\n", i, h.cards, h.bid, h.value, h.hexcards)
	}
	fmt.Printf("Total: %d\n", total)
}
