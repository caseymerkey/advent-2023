package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var calc_memo map[string]int

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

	p1Total := 0
	p2Total := 0
	lineNumber := 0
	var startTime = time.Now()
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		springs, groups := parseLine(line)
		lineTotal := calc(springs, groups)
		// fmt.Printf("## Line %d: %d combinations\n", lineNumber, lineTotal)
		p1Total += lineTotal

		springs, groups = unfold(line)
		lineTotal = calc(springs, groups)
		// fmt.Printf("## Line %d: %d combinations\n", lineNumber, lineTotal)
		p2Total += lineTotal
	}
	executionTime := float32(time.Now().Sub(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("\nCompleted in %f seconds\n", executionTime)
	fmt.Printf("Part 1 Total: %d\n", p1Total)
	fmt.Printf("Part 2 Total: %d\n", p2Total)
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

func parseLine(line string) (string, []int) {
	s := strings.Split(line, " ")
	springs := s[0]
	groups := stringToIntArr(s[1])

	return springs, groups
}

func unfold(line string) (string, []int) {

	s := strings.Split(line, " ")
	springs := s[0]
	groupsString := s[1]
	for i := 0; i < 4; i++ {
		springs += ("?" + s[0])
		groupsString += ("," + s[1])
	}
	groups := stringToIntArr(groupsString)
	return springs, groups

}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

// I followed along with Reddit user u/StaticMoose and their walkthrough here:
// https://www.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/
// which finally turned the light bulb on for how to approach this problem.
// Converted from their Python code to Go, and wrote come custom implementations of features not availbable
// in Go, notably memoization and a replacement for tuples
func calc(record string, groups []int) int {

	if calc_memo == nil {
		calc_memo = make(map[string]int)
	}

	key := record + "|" + arrayToString(groups, ",")
	// fmt.Printf("Key: %s \n", key)
	v, ok := calc_memo[key]
	if ok {
		return v
	}

	var out int

	// Did we run out of groups? We might still be valid
	if len(groups) == 0 {
		// Make sure there aren't any more damaged springs. If so, we're valid
		if !strings.Contains(record, "#") {
			// This will return true even if record is empty, which is valid
			return 1
		} else {
			// More damaged springs that aren't in the groups
			return 0
		}
	}

	// There are more groups, but no more record
	if len(record) == 0 {
		// We can't fit. Exit
		return 0
	}

	// Look at the next element in each record and group
	next_character := record[0]
	next_group := groups[0]

	// Logic that treats the first character as pound-sign "#"
	pound := func() int {
		// If the first is a pound, then the first n characters must be
		// able to be treated as a pound, where n is the first group number
		var this_group string
		if len(record) < next_group {
			this_group = record
		} else {
			this_group = record[:next_group]
		}
		this_group = strings.ReplaceAll(this_group, "?", "#")

		// If the next group can't fit all the damaged springs, then abort
		if this_group != strings.Repeat("#", next_group) {
			return 0
		}

		// If the rest of the record is just the last group, then we're
		// done and there's only one possibility
		if len(record) == next_group {
			// Make sure this is the last group
			if len(groups) == 1 {
				// We are valid
				return 1
			} else {
				// There are more groups, and this won't work
				return 0
			}
		}

		// Make sure the character that follows this group can be a separator
		if record[next_group] == '.' || record[next_group] == '?' {
			// It can be separator, so skip it and reduce to the next group

			return calc(record[(next_group+1):], groups[1:])
		}

		//Can't be handled, there are no possibilites
		return 0
	}

	// Logic that treats the first character as dot "."
	dot := func() int {
		// We just skip over the dot looking for the next pound
		return calc(record[1:], groups)
	}

	switch next_character {
	case '#':
		// Test pound logic
		out = pound()
	case '.':
		// Test dot logic
		out = dot()
	case '?':
		// This character could be either character, so we'll explore both
		// possibilities
		out = dot() + pound()
	}

	// Help with debugging
	// fmt.Println(record, groups, "->", out)
	calc_memo[key] = out
	return out
}

/*
 * This is all stuff that "worked" for part 1, but was not viable for part 2. Keeping it
 * in the file for posterity.
 */

type EvaluationState struct {
	testString       string
	i                int
	state            int
	groupIndex       int
	currentGroupSize int
}

const (
	in_group       = 1
	between_groups = 0
)

func evalEverything(springs string, groups []int) int {

	questions := strings.Count(springs, "?")
	permutations := int64(math.Pow(2, float64(questions)))

	fmt.Printf("Evaluating %d permutations\n", permutations)

	var m map[int64]EvaluationState
	m = make(map[int64]EvaluationState)

	for p := int64(0); p < permutations; p++ {
		springsArr := []byte(springs)
		bin := PadLeft(strconv.FormatInt(p, 2), questions)
		b := []byte(bin)
		idx := 0
		for _, ch := range b {
			var target byte
			if ch == '0' {
				target = '.'
			} else {
				target = '#'
			}
			idx = replaceFirstFrom(springsArr, '?', target, idx)
		}
		test := string(springsArr)
		evalState := EvaluationState{test, 0, between_groups, -1, 0}

		if testStringAtPosition(&evalState, &groups) {
			m[p] = evalState
		}
	}

	for i := 1; i < len(springs); i++ {

		for p, evalState := range m {
			evalState.i = i
			if testStringAtPosition(&evalState, &groups) {
				m[p] = evalState
			} else {
				delete(m, p)
			}
		}

	}

	// for _, v := range m {
	// 	fmt.Printf("%s passed\n", v.testString)
	// }
	return len(m)
}

func testStringAtPosition(evalState *EvaluationState, groups *[]int) bool {

	ch := evalState.testString[evalState.i]
	if ch == '#' {
		switch evalState.state {
		case between_groups:
			evalState.groupIndex++
			if evalState.groupIndex >= len((*groups)) {
				// We've found more groups of hashes than we should have
				return false
			}
			evalState.state = in_group
			evalState.currentGroupSize = (*groups)[evalState.groupIndex] - 1
		case in_group:
			if evalState.currentGroupSize == 0 {
				// more hashes in this group than there should be
				return false
			}
			evalState.currentGroupSize--
		}
	} else {
		// ch == '.'
		switch evalState.state {
		case in_group:
			if evalState.currentGroupSize > 0 {
				// This group of hashes was not large enough
				return false
			}
			evalState.state = between_groups

		case between_groups:
			// Still between groups. This is fine.
		}
	}

	if evalState.i >= len(evalState.testString)-1 {
		if evalState.groupIndex < (len((*groups))-1) || evalState.currentGroupSize != 0 {
			return false
		}
	}

	return true
}

func evaluate(springs string, groups []int) int {

	rowTotal := 0

	questions := strings.Count(springs, "?")
	permutations := int64(math.Pow(2, float64(questions)))

	fmt.Printf("Evaluating %d permutations\n", permutations)

	for i := int64(0); i < permutations; i++ {
		springsArr := []byte(springs)
		bin := PadLeft(strconv.FormatInt(i, 2), questions)
		b := []byte(bin)
		idx := 0
		for _, ch := range b {
			var target byte
			if ch == '0' {
				target = '.'
			} else {
				target = '#'
			}
			idx = replaceFirstFrom(springsArr, '?', target, idx)
		}
		test := string(springsArr)

		itFits := customEval(groups, test)
		if itFits {
			rowTotal++
		}

	}

	return rowTotal

}

func replaceFirstFrom(str []byte, old, new byte, idx int) int {

	if idx >= len(str) {
		return -1
	}

	for i := idx; i < len(str); i++ {
		if str[i] == old {
			str[i] = new
			return i
		}
	}
	return -1

}

func PadLeft(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}

func customEval(groups []int, testString string) bool {

	state := between_groups
	groupIdx := -1
	currentGroupSize := 0

	for i := 0; i < len(testString); i++ {
		ch := testString[i]
		if ch == '#' {
			switch state {
			case between_groups:
				groupIdx++
				if groupIdx >= len(groups) {
					// We've found more groups of hashes than we should have
					return false
				}
				state = in_group
				currentGroupSize = groups[groupIdx] - 1
			case in_group:
				if currentGroupSize == 0 {
					// more hashes in this group than there should be
					return false
				}
				currentGroupSize--
			}
		} else {
			// ch == '.'
			switch state {
			case in_group:
				if currentGroupSize > 0 {
					// This group of hashes was not large enough
					return false
				}
				state = between_groups

			case between_groups:
				// Still between groups. This is fine.
			}
		}
	}
	if currentGroupSize != 0 || groupIdx != (len(groups)-1) {
		return false
	}
	return true

}
