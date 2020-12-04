package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(data []byte) {
	validPass := 0
	for _, line := range strings.Split(string(data), "\n") {
		items := strings.Split(line, " ")
		minMax := strings.Split(items[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		letter := strings.Split(items[1], ":")
		lc := strings.Count(items[2], letter[0])

		if lc <= max && lc >= min {
			validPass++
		}
	}

	fmt.Println(validPass)
}

func part2(data []byte) {
	validPass := 0
	for _, line := range strings.Split(string(data), "\n") {
		items := strings.Split(line, " ")
		positions := strings.Split(items[0], "-")
		pos1, _ := strconv.Atoi(positions[0])
		pos2, _ := strconv.Atoi(positions[1])
		letter := strings.Split(items[1], ":")
		i := items[2]

		if ((letter[0] == string(i[pos1-1])) && !(letter[0] == string(i[pos2-1]))) ||
			(!(letter[0] == string(i[pos1-1])) && (letter[0] == string(i[pos2-1]))) {
			validPass++
		}
	}

	fmt.Println(validPass)
}

func main() {
	data, err := ioutil.ReadFile("day2_input.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}
	part1(data)
	part2(data)
}
