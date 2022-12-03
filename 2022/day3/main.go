package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day3_input")
	if err != nil {
		log.Fatal("failed to open file")
	}

	data := strings.Split(string(input), "\n")
	if err != nil {
		log.Fatalf("failed split the input lines")
	}

	part1(data)
	part2(data)
}

func findCommonLetter(p1, p2 string) string {
	for _, letter := range strings.Split(p2, "") {
		if strings.Contains(p1, letter) {
			// fmt.Printf("for line %s%s found common letter: %s \n", p1, p2, letter)
			return letter
		}
	}

	return ""
}

func findCommonLetterMulti(p1, p2, p3 string) string {
	for _, letter := range strings.Split(p1, "") {
		if strings.Contains(p2, letter) && strings.Contains(p3, letter) {
			return letter
		}
	}

	return ""
}

func findLetterPriority(letter string) int64 {
	var priority int64
	priority = 1
	for l := 'a'; l <= 'z'; l++ {
		// fmt.Println(string(l))
		if letter == string(l) {
			return priority
		}

		priority++
	}

	for l := 'A'; l <= 'Z'; l++ {
		// fmt.Println(l)
		if letter == string(l) {
			return priority
		}

		priority++
	}

	return 0
}

func part1(data []string) {
	var sum int64
	for _, line := range data {
		p1 := line[:len(line)/2]
		p2 := line[len(line)/2:]
		// fmt.Printf("%s  :  %s \n", p1, p2)

		if letter := findCommonLetter(p1, p2); letter != "" {
			if p := findLetterPriority(letter); p != 0 {
				sum += p
			}
		}
	}

	fmt.Println("Sum of priorities is: ", sum)
}

func part2(data []string) {
	var sum int64
	groupsCounter := 0
	var lines []string

	for _, line := range data {
		lines = append(lines, line)
		if groupsCounter < 2 {
			groupsCounter++
			continue
		}

		// fmt.Printf("len is: %d and value %v\n", len(lines), lines)

		if letter := findCommonLetterMulti(lines[0], lines[1], lines[2]); letter != "" {
			if p := findLetterPriority(letter); p != 0 {
				sum += p
			}
		}

		lines = nil
		groupsCounter = 0
	}

	fmt.Println("Sum of priorities of groups is: ", sum)

}
