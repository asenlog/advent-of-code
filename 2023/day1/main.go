package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func part1(input []string) {
	totalSum := 0

	for _, line := range input {
		fd, ld := 0, 0

		// normal order
		for _, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				fd = digit
				//fmt.Println(fd)
				break
			}
		}

		// reverse order
		for i := len(line) - 1; i >= 0; i-- {
			if digit, err := strconv.Atoi(string(line[i])); err == nil {
				ld = digit
				//fmt.Println(ld)
				break
			}
		}

		concatValue := fd*10 + ld
		totalSum += concatValue
	}

	fmt.Println("Part 1:", totalSum)
}

func findFirstDigit(s string) int {
	stringDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Iinit index
	lowestIndex := len(s) + 1
	foundDigit := 0

	// Check string digits
	for i, digit := range stringDigits {
		index := strings.Index(s, digit)
		if index != -1 && index < lowestIndex {
			// string digit with a lower index
			lowestIndex = index
			foundDigit = i + 1
		}
	}

	// Check number digits
	for i, char := range s {
		if unicode.IsDigit(char) && i < lowestIndex {
			// numer digit with a lower index
			lowestIndex = i
			foundDigit, _ = strconv.Atoi(string(char))
		}
	}

	if lowestIndex == len(s)+1 {
		log.Println("No digit found")
		return 0
	}

	return foundDigit
}

func findLastDigit(s string) int {
	stringDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// init negative index
	highestIndex := -1
	foundDigit := 0

	// Check string digits
	for i, digit := range stringDigits {
		index := strings.LastIndex(s, digit)
		if index > highestIndex {
			// string digit with a higher index
			highestIndex = index
			foundDigit = i + 1
		}
	}

	// Check number digits
	for i, char := range s {
		if unicode.IsDigit(char) && i > highestIndex {
			// number digit with a higher index
			highestIndex = i
			foundDigit, _ = strconv.Atoi(string(char))
		}
	}

	if highestIndex == -1 {
		log.Println("No digit found")
		return 0
	}

	return foundDigit
}

func part2(input []string) {
	totalSum := 0

	for _, line := range input {
		fd, ld := 0, 0

		fd = findFirstDigit(line)
		ld = findLastDigit(line)

		concatValue := fd*10 + ld
		totalSum += concatValue
	}

	fmt.Println("Part 2:", totalSum)
}

func parseInput(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	return strings.Split(string(data), "\n"), nil
}

func main() {
	data, err := parseInput("day1_input")
	if err != nil {
		log.Fatalf("failed to open the file")
	}

	part1(data)
	part2(data)
}
