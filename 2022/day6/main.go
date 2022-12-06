package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("day6_input")
	if err != nil {
		log.Fatal("failed to open file")
	}

	part1(input)
	part2(input)
}

func findMarker(data []byte, distinctChars int) int {
	var lastXChars []byte
	for i, char := range data {
		lastXChars = append(lastXChars, char)

		if len(lastXChars) > distinctChars {
			lastXChars = lastXChars[1:]
		}

		if len(lastXChars) == distinctChars && areDifferent(lastXChars) {
			return i + 2
		}
	}

	return 0
}

func areDifferent(data []byte) bool {
	seen := make(map[byte]bool)
	for _, b := range data {
		if seen[b] {
			return false
		}
		seen[b] = true
	}

	return true
}

func part1(data []byte) {
	markerIndex := findMarker(data, 4)
	fmt.Println("First marker at index:", markerIndex)
}

func part2(data []byte) {
	markerIndex := findMarker(data, 14)
	fmt.Println("Second marker at index:", markerIndex)
}
