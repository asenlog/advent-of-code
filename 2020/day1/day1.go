package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	return strings.Split(string(data), "\n"), nil
}

func part1(data []string) {
	for _, s1 := range data {
		for _, s2 := range data {
			i1, _ := strconv.ParseInt(s1, 10, 64)
			i2, _ := strconv.ParseInt(s2, 10, 64)
			if i1+i2 == 2020 {
				fmt.Println(i1 * i2)
				return
			}
		}
	}
}

func part2(data []string) {
	for _, s1 := range data {
		for _, s2 := range data {
			for _, s3 := range data {
				i1, _ := strconv.ParseInt(s1, 10, 64)
				i2, _ := strconv.ParseInt(s2, 10, 64)
				i3, _ := strconv.ParseInt(s3, 10, 64)
				if i1+i2+i3 == 2020 {
					fmt.Println(i1 * i2 * i3)
					return
				}
			}
		}
	}
}

func main() {
	data, err := parseInput("day1_input.txt")
	if err != nil {
		log.Fatalf("failed to open the file")
	}
	part1(data)
	part2(data)
}
