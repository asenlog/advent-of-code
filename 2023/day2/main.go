package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

func part1(input []string) {
	totalSum := 0
	gameScores := make(map[int]bool)

	for _, line := range input {
		// Assign Game Numbers to 0
		game := strings.Split(line, ":")
		gameNumber := regexp.MustCompile(`\d+`).FindString(game[0])
		gn, _ := strconv.Atoi(gameNumber)
		gameScores[gn] = true

		sets := strings.Split(game[1], ";")
		for _, set := range sets {
			colorsSets := strings.Split(set, ",")
			mapColors := make(map[string]int)
			for _, colorsSet := range colorsSets {
				//fmt.Println(colorsSet)
				parts := strings.Fields(colorsSet)
				color := parts[1]
				value := 0
				if v, err := strconv.Atoi(parts[0]); err == nil {
					value = v
				}
				mapColors[color] = value
			}

			if mapColors["red"] > red || mapColors["green"] > green || mapColors["blue"] > blue {
				gameScores[gn] = false
				break
			}
		}
	}

	for k, v := range gameScores {
		//fmt.Println("Game:", k, "Score:", v)
		if v {
			totalSum += k
		}
	}

	fmt.Println("Part 1:", totalSum)
}

func part2(input []string) {
	totalSum := 0

	for _, line := range input {
		// Assign Game Numbers to 0
		game := strings.Split(line, ":")
		sets := strings.Split(game[1], ";")
		mapColors := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
		for _, set := range sets {
			colorsSets := strings.Split(set, ",")
			for _, colorsSet := range colorsSets {
				//fmt.Println(colorsSet)
				parts := strings.Fields(colorsSet)
				color := parts[1]
				value := 0
				if v, err := strconv.Atoi(parts[0]); err == nil {
					value = v
				}
				if value > mapColors[color] {
					mapColors[color] = value
				}
			}

		}
		totalSum += mapColors["red"] * mapColors["green"] * mapColors["blue"]
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
	data, err := parseInput("day2_input")
	if err != nil {
		log.Fatalf("failed to open the file")
	}

	part1(data)
	part2(data)
}
