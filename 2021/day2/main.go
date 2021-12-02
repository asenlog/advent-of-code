package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int64
	Down int64
	UP int64
	Depth int64
}



func part1(data []string) {
	position := Position{}

	for _, line := range data {
		positionFields := strings.Fields(line)
		steps, _ := strconv.ParseInt(positionFields[1], 10, 64)

		switch positionFields[0] {
		case "forward":
			position.Horizontal += steps
		case "up":
			position.UP += steps
		case "down":
			position.Down += steps
		}
	}

	fmt.Println("Position and Depth: ",  position.Horizontal * (position.Down-position.UP))
}

func part2(data []string) {
	position := Position{}

	for _, line := range data {
		positionFields := strings.Fields(line)
		steps, _ := strconv.ParseInt(positionFields[1], 10, 64)

		switch positionFields[0] {
		case "forward":
			position.Horizontal += steps
			if (position.Down-position.UP) == 0 {
				break
			}
			position.Depth += steps * (position.Down-position.UP)
		case "up":
			position.UP += steps
		case "down":
			position.Down += steps
		}
	}

	fmt.Println("Position and Depth: ",  position.Horizontal * position.Depth)
}

func parseInput(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	return strings.Split(string(data), "\n"), nil
}

func main() {
	data, err := parseInput("day2_input.txt")
	if err != nil {
		log.Fatalf("failed to open the file")
	}
	part1(data)
	part2(data)
}