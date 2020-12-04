package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(data *os.File) {
}

func part2(data *os.File) {
	sc := bufio.NewScanner(data)
	var (
		rightSteps, trees [5]int
		line              int
	)
	for sc.Scan() {
		for slope := 0; slope < 5; slope++ {
			if slope == 4 && line%2 != 0 {
				continue
			}
			if sc.Text()[rightSteps[slope]%len(sc.Text())] == '#' {
				trees[slope]++
			}
			rightSteps[slope] += (2*slope + 1) % 8
		}
		line++
	}
	fmt.Println(trees[0] * trees[1] * trees[2] * trees[3] * trees[4])
}

func main() {
	data, err := os.Open("day3_input.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer data.Close()

	part2(data)
}
