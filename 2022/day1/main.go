package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func part1(data []string) {
	counter := int64(0)
	var prevSum, sum int64
	elfWithMostCalories := int64(0)

	for _, line := range data {
		if line == "" {
			if sum > prevSum {
				elfWithMostCalories = counter
				prevSum = sum
			}

			counter += 1
			sum = 0
			continue
		}

		calories, _ := strconv.ParseInt(line, 10, 64)

		sum += calories
	}

	fmt.Printf("Elf caring the most calories: %d : %d", elfWithMostCalories, prevSum)
}

type Pair struct {
	Key   int64
	Value int64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func part2(data []string) {
	counter := int64(0)
	var prevSum, sum int64
	elvesWithCalories := map[int64]int64{}

	for _, line := range data {
		if line == "" {
			if sum > prevSum {
				elvesWithCalories[counter] = sum
			}

			counter += 1
			sum = 0
			continue
		}

		calories, _ := strconv.ParseInt(line, 10, 64)
		sum += calories
	}

	p := make(PairList, len(elvesWithCalories))

	i := 0
	for k, v := range elvesWithCalories {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	fmt.Printf("Total calories of top 3 elves: %+v",
		elvesWithCalories[p[len(p)-1].Key]+elvesWithCalories[p[len(p)-2].Key]+elvesWithCalories[p[len(p)-3].Key])
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
	fmt.Println()
	part2(data)
}
