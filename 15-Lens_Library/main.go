package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile("([a-z]+)([=-])([0-9]?)")
var boxes [256][]lens

type lens struct {
	label       string
	focalLength int
}

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
	reader := bufio.NewReader(file)
	var bytes []byte = make([]byte, 0)
	var part1Total int
	for {
		b, err := reader.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}
		// process byte
		if b == ',' {
			part1Total += hashItUp(bytes)
			part2(bytes)
			bytes = make([]byte, 0)
		} else if b != '\n' && b != '\x00' {
			bytes = append(bytes, b)
		}

		if err != nil {
			// end of file
			part1Total += hashItUp(bytes)
			part2(bytes)
			break
		}
	}
	fmt.Printf("Part 1 Total: %d\n", part1Total)

	total := 0
	for i, box := range boxes {
		for j, lens := range box {
			t := ((i + 1) * (j + 1) * lens.focalLength)
			fmt.Printf("%s: (box %d) * (slot %d) * (length %d) = %d\n", lens.label, 1+i, j+1, lens.focalLength, t)
			total += t
		}
	}
	fmt.Printf("Part 2 Total: %d\n", total)
}

func hashItUp(str []byte) int {

	var hash int
	for _, b := range str {
		hash += int(b)
		hash = hash * 17
		hash = hash % 256
	}
	// fmt.Printf("%s  ==>  %d\n", string(str), hash)
	return hash
}

func part2(bytes []byte) {
	matches := re.FindStringSubmatch(string(bytes))

	label := matches[1]
	op := matches[2]
	flen := matches[3]

	hash := hashItUp([]byte(label))
	box := boxes[hash]

	if op == "=" {
		added := false
		focalLength, _ := strconv.Atoi(flen)
		newLens := lens{label: label, focalLength: focalLength}
		for i, l := range box {
			if l.label == label {
				box[i] = newLens
				added = true
				break
			}
		}
		if !added {
			box = append(box, newLens)
		}
		boxes[hash] = box
	} else {
		for i, l := range box {
			if l.label == label {
				if l.label == label {
					box = remove(box, i)
				}
			}
		}
		boxes[hash] = box
	}

}

func remove(slice []lens, s int) []lens {
	return append(slice[:s], slice[s+1:]...)
}
