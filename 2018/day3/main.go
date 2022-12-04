package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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

func part1(data []string) {
	claimRegex := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	var claims []string
	for _, line := range data {
		match := claimRegex.FindStringSubmatch(line)
		if match != nil {
			id := match[1]
			x := match[2]
			y := match[3]
			w := match[4]
			h := match[5]
			claims = append(claims, fmt.Sprintf("#%s @ %s,%s: %sx%s", id, x, y, w, h))
		}
	}

	var fabric [10000][10000]int

	for _, claim := range claims {
		var id, x, y, w, h int
		fmt.Sscanf(claim, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				fabric[i][j]++
			}
		}
	}

	fmt.Println(len(fabric))
	var overlap int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] >= 2 {
				overlap++
			}
		}
	}
	fmt.Println(overlap)
}

func part2(data []string) {
	claimRegex := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	var claims []string
	for _, line := range data {
		match := claimRegex.FindStringSubmatch(line)
		if match != nil {
			id := match[1]
			x := match[2]
			y := match[3]
			w := match[4]
			h := match[5]
			claims = append(claims, fmt.Sprintf("#%s @ %s,%s: %sx%s", id, x, y, w, h))
		}
	}

	var fabric [10000][10000]int

	for _, claim := range claims {
		var id, x, y, w, h int
		fmt.Sscanf(claim, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				fabric[i][j]++
			}
		}
	}

	fmt.Println(len(fabric))
	for _, claim := range claims {
		var id, x, y, w, h int
		fmt.Sscanf(claim, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		nonOverlapping := true
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				if fabric[i][j] != 1 {
					nonOverlapping = false
					break
				}
			}
			if !nonOverlapping {
				break
			}
		}
		if nonOverlapping {
			fmt.Printf("%d\n", id)
		}
	}
}
