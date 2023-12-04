package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	card    string
	winning []int
	numbers []int
}

type Cards []Card

func sortNums(numStr string) []int {
	var ints []int

	ns := strings.Fields(numStr)
	for i := range ns {
		num, err := strconv.Atoi(ns[i])
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, num)
	}

	sort.Ints(ints)

	return ints
}

func part1(data []string) {
	totalSum := 0
	var cards Cards

	for _, v := range data {
		card := strings.Split(v, ": ")
		split := strings.Split(card[1], " | ")
		numbers := sortNums(split[1])
		winning := sortNums(split[0])
		cardInfo := Card{
			card:    card[0],
			winning: winning,
			numbers: numbers,
		}
		cards = append(cards, cardInfo)
	}

	for _, val := range cards {
		points := 0
		for _, num := range val.numbers {
			if slices.Contains(val.winning, num) && points == 0 {
				points = 1
			} else if slices.Contains(val.winning, num) {
				points = points * 2
			}
		}
		totalSum += points
	}

	fmt.Println("Part 1:", totalSum)
}

func part2(data []string) {
	totalSum := 0
	linesLen := len(data)
	played := make([]int, linesLen)

	for i := 0; i < linesLen; i++ {
		played = append(played, 0)
	}

	index := 0

	for _, v := range data {
		played[index] += 1

		game := strings.Split(v, ": ")[1]

		winning := strings.Split(game, " | ")[0]

		ours := strings.Split(game, " | ")[1]

		winnum := numSlice(winning)
		ournum := numSlice(ours)

		intersection := findIntersection(winnum, ournum)

		for w := range intersection {
			played[index+w+1] += played[index]
		}

		index++
	}

	for i := range played {
		totalSum += played[i]
	}

	fmt.Println("Part 2:", totalSum)
}

func findIntersection(slice1, slice2 []int) []int {
	sort.Ints(slice1)
	sort.Ints(slice2)

	var intersection []int

	for _, num := range slice1 {
		if existsInSlice(slice2, num) && !existsInSlice(intersection, num) {
			intersection = append(intersection, num)
		}
	}

	return intersection
}

// Helper function to check if an element exists in a slice
func existsInSlice(slice []int, element int) bool {
	for _, num := range slice {
		if num == element {
			return true
		}
	}
	return false
}

func numSlice(nss string) []int {
	numstrsl := strings.Fields(nss)

	var numintsl []int

	for _, n := range numstrsl {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numintsl = append(numintsl, num)
	}

	return numintsl
}

func parseInput(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	return strings.Split(string(data), "\n"), nil
}

func main() {
	data, err := parseInput("day4_input")
	if err != nil {
		log.Fatalf("failed to open the file")
	}

	part1(data)
	part2(data)
}
