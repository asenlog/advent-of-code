package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day2_input")
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

func part1(data []string) {
	var twos, threes int

	for _, id := range data {
		counts := make(map[rune]int)
		for _, c := range id {
			counts[c]++
		}

		var hasTwo, hasThree bool
		for _, count := range counts {
			if count == 2 {
				hasTwo = true
			} else if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	checksum := twos * threes
	fmt.Printf("The checksum is %d\n", checksum)
}

func part2(data []string) {
	for i, id1 := range data {
		for j, id2 := range data {
			if i == j {
				continue
			}

			var diff int
			var common []string
			for k := range id1 {
				if id1[k] != id2[k] {
					diff++
				} else {
					common = append(common, string(id1[k]))
				}
			}

			if diff == 1 {
				fmt.Println(common)
				return
			}
		}
	}
}
